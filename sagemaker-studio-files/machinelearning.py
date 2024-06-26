### !pip install --upgrade pandas xgboost
!pip install --force-reinstall statsmodels==0.14.0 pydeseq2==0.4.3

# Import necessary libraries
import boto3
import xgboost as xgb
import pandas as pd
from sklearn.metrics import mean_absolute_error
from sklearn.model_selection import train_test_split
from sagemaker import get_execution_role
from sagemaker.inputs import TrainingInput
from sagemaker.xgboost.estimator import XGBoost
import matplotlib.pyplot as plt
from datetime import timedelta


#############################################################

# Setup and configuration
role = get_execution_role()
bucket = 'stock-trading-bucket-crimson'
symbol = "XOM"
data_key = f'updated-csv/{symbol}2.csv'
data_location = f's3://{bucket}/{data_key}'

# Load data from S3
s3_client = boto3.client('s3')
obj = s3_client.get_object(Bucket=bucket, Key=data_key)
data = pd.read_csv(obj['Body'])

# Convert 'Datetime' to datetime and set it as index
data['Datetime'] = pd.to_datetime(data['Datetime'])
data.set_index('Datetime', inplace=True)

# start_date = '2012-01-12'  # Replace YYYY-MM-DD with the actual start date

# # Filter the data to include only records from the start date onwards
# data = data[start_date:]


# Print the column names to verify them
print(data.columns)
# ############################################################################
# Prepare the data
# Columns to drop
columns_to_drop = ['High', 'Low', 'Open', 'Period', 'Symbol', 'MovingAverage50Days', 'MovingAverage200Days', 'MarketCap']

# Dropping specified columns
X = data.drop(columns_to_drop + ['Close'], axis=1)
y = data['Close']



interval = 5  # split to data ratio 1 out of every 5 

# Create boolean mask
mask = [True if i % interval == 0 else False for i in range(len(X))]

# Split the data using the mask
X_train = X[~pd.Series(mask, index=X.index)]
X_val = X[pd.Series(mask, index=X.index)]
y_train = y[~pd.Series(mask, index=y.index)]
y_val = y[pd.Series(mask, index=y.index)]

# Print the sizes of each dataset to verify the split
print(X.columns)
print(f"Training data size: {X_train.shape[0]}")
print(f"Validation data size: {X_val.shape[0]}")


#####################################################################


# Convert data into DMatrix objects
dtrain = xgb.DMatrix(X_train, label=y_train)
dval = xgb.DMatrix(X_val, label=y_val)


# XGBoost parameters
params = {
    'objective': 'reg:squarederror',
    'max_depth': 3,
    'eta': 0.1,
    'eval_metric': 'rmse'
}
num_round = 1000
evals = [(dtrain, 'train'), (dval, 'eval')]

# Train the model
model = xgb.train(params, dtrain, num_boost_round=num_round, evals=evals)

# Predicting with the validation dataset
dval = xgb.DMatrix(X_val)
val_predictions = model.predict(dval)
###################################################################
def stock_rate_one_year_position_calculator(finalPosition, previousPosition):
    if previousPosition == 0:
        return 1.0
    return finalPosition / previousPosition


def stock_position_no_acceleration_calculator(finalPosition, previousPosition ):
    # x = xi + vt (t being 1)
    # v = delta(x)/ t being 1
    velocity = finalPosition  - previousPosition
    predictedPosition = finalPosition + velocity

    return predictedPosition


def stock_one_year_position_calculator(currentPosition, currentVelocityRate, oldPosition, oldVelocityRate):
    # x + v * t + 1/2 * a (t)^2
    #  xf = xi + v + 1/2 a - plugging in 1 for the time interval. 
    # a will be the change of growth rate
    # since this growth rate is year over year time will be 1
    #  xf = xi + v + 1/2 a - a simple formula to estimate the new position of the stock
    # velocity rate   = currentPosition /oldPosition 
    
    # Calculate velocity 
    velocityCurrent  = currentPosition - oldPosition
    twoYearOldPosition =  oldPosition / oldVelocityRate
    velocityOld = oldPosition - twoYearOldPosition
    acceleration = velocityCurrent - velocityOld
    # Took the absolute value to cushion some of the results from the negative
    if velocityCurrent < 0:
        return currentPosition
    predictedPosition = currentPosition + velocityCurrent + (.5 * abs(acceleration))
    

    return predictedPosition


def future_data_frame(currentDate, X, itt):
    one_year_out = currentDate + timedelta(days=365)
    future_data = pd.DataFrame(index=[one_year_out], columns=X.columns)
    ittPrevious = itt -252
    
    future_data.loc[one_year_out,'Volume'] = stock_position_no_acceleration_calculator(X['Volume'].iloc[itt],X['Volume'].iloc[ittPrevious])
    future_data.loc[one_year_out,'OperatingCashFlowPerShare'] = stock_one_year_position_calculator(X['OperatingCashFlowPerShare'].iloc[itt],X['YearOverYearRateOperatingCashFlowPerShare'].iloc[itt],X['OperatingCashFlowPerShare'].iloc[ittPrevious],X['YearOverYearRateOperatingCashFlowPerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateOperatingCashFlowPerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'OperatingCashFlowPerShare'], X['OperatingCashFlowPerShare'].iloc[itt])
    future_data.loc[one_year_out,'FreeCashFlowPerShare'] = stock_one_year_position_calculator(X['FreeCashFlowPerShare'].iloc[itt],X['YearOverYearRateFreeCashFlowPerShare'].iloc[itt],X['FreeCashFlowPerShare'].iloc[ittPrevious],X['YearOverYearRateFreeCashFlowPerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateFreeCashFlowPerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'FreeCashFlowPerShare'], X['FreeCashFlowPerShare'].iloc[itt])
    future_data.loc[one_year_out,'CashPerShare'] = stock_one_year_position_calculator(X['CashPerShare'].iloc[itt],X['YearOverYearRateCashPerShare'].iloc[itt],X['CashPerShare'].iloc[ittPrevious],X['YearOverYearRateCashPerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateCashPerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'CashPerShare'], X['CashPerShare'].iloc[itt])
    future_data.loc[one_year_out,'PriceToSalesRatio'] = stock_one_year_position_calculator(X['PriceToSalesRatio'].iloc[itt],X['YearOverYearRatePriceToSalesRatio'].iloc[itt],X['PriceToSalesRatio'].iloc[ittPrevious],X['YearOverYearRatePriceToSalesRatio'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRatePriceToSalesRatio'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'PriceToSalesRatio'], X['PriceToSalesRatio'].iloc[itt])
    future_data.loc[one_year_out,'PayoutRatio'] = stock_one_year_position_calculator(X['PayoutRatio'].iloc[itt],X['YearOverYearRatePayoutRatio'].iloc[itt],X['PayoutRatio'].iloc[ittPrevious],X['YearOverYearRatePayoutRatio'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRatePayoutRatio'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'PayoutRatio'], X['PayoutRatio'].iloc[itt])
    future_data.loc[one_year_out,'RevenuePerShare'] = stock_one_year_position_calculator(X['RevenuePerShare'].iloc[itt],X['YearOverYearRateRevenuePerShare'].iloc[itt],X['RevenuePerShare'].iloc[ittPrevious],X['YearOverYearRateRevenuePerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateRevenuePerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'RevenuePerShare'], X['RevenuePerShare'].iloc[itt])
    future_data.loc[one_year_out,'BookValuePerShare'] = stock_one_year_position_calculator(X['BookValuePerShare'].iloc[itt],X['YearOverYearRateBookValuePerShare'].iloc[itt],X['BookValuePerShare'].iloc[ittPrevious],X['YearOverYearRateBookValuePerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateBookValuePerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'BookValuePerShare'], X['BookValuePerShare'].iloc[itt])
    future_data.loc[one_year_out,'PeRatio'] = stock_one_year_position_calculator(X['PeRatio'].iloc[itt],X['YearOverYearRatePeRatio'].iloc[itt],X['PeRatio'].iloc[ittPrevious],X['YearOverYearRatePeRatio'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRatePeRatio'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'PeRatio'], X['PeRatio'].iloc[itt])
    future_data.loc[one_year_out,'PfcfRatio'] = stock_one_year_position_calculator(X['PfcfRatio'].iloc[itt],X['YearOverYearRatePfcfRatio'].iloc[itt],X['PfcfRatio'].iloc[ittPrevious],X['YearOverYearRatePfcfRatio'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRatePfcfRatio'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'PfcfRatio'], X['PfcfRatio'].iloc[itt])
    future_data.loc[one_year_out,'EvToOperatingCashFlow'] = stock_one_year_position_calculator(X['EvToOperatingCashFlow'].iloc[itt],X['YearOverYearRateEvToOperatingCashFlow'].iloc[itt],X['EvToOperatingCashFlow'].iloc[ittPrevious],X['YearOverYearRateEvToOperatingCashFlow'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateEvToOperatingCashFlow'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'EvToOperatingCashFlow'], X['EvToOperatingCashFlow'].iloc[itt])
    future_data.loc[one_year_out,'NetDebtToEBITDA'] = stock_one_year_position_calculator(X['NetDebtToEBITDA'].iloc[itt],X['YearOverYearRateNetDebtToEBITDA'].iloc[itt],X['NetDebtToEBITDA'].iloc[ittPrevious],X['YearOverYearRateNetDebtToEBITDA'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateNetDebtToEBITDA'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'NetDebtToEBITDA'], X['NetDebtToEBITDA'].iloc[itt])
    future_data.loc[one_year_out,'StockBasedCompensationToRevenue'] = stock_one_year_position_calculator(X['StockBasedCompensationToRevenue'].iloc[itt],X['YearOverYearRateStockBasedCompensationToRevenue'].iloc[itt],X['StockBasedCompensationToRevenue'].iloc[ittPrevious],X['YearOverYearRateStockBasedCompensationToRevenue'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateStockBasedCompensationToRevenue'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'StockBasedCompensationToRevenue'], X['StockBasedCompensationToRevenue'].iloc[itt])
    future_data.loc[one_year_out,'GrahamNumber'] = stock_one_year_position_calculator(X['GrahamNumber'].iloc[itt],X['YearOverYearRateGrahamNumber'].iloc[itt],X['GrahamNumber'].iloc[ittPrevious],X['YearOverYearRateGrahamNumber'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateGrahamNumber'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'GrahamNumber'], X['GrahamNumber'].iloc[itt])
    future_data.loc[one_year_out,'Roic'] = stock_one_year_position_calculator(X['Roic'].iloc[itt],X['YearOverYearRateRoic'].iloc[itt],X['Roic'].iloc[ittPrevious],X['YearOverYearRateRoic'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateRoic'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'Roic'], X['Roic'].iloc[itt])
    future_data.loc[one_year_out,'Roe'] = stock_one_year_position_calculator(X['Roe'].iloc[itt],X['YearOverYearRateRoe'].iloc[itt],X['Roe'].iloc[ittPrevious],X['YearOverYearRateRoe'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateRoe'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'Roe'], X['Roe'].iloc[itt])
    future_data.loc[one_year_out,'CapexPerShare'] = stock_one_year_position_calculator(X['CapexPerShare'].iloc[itt],X['YearOverYearRateCapexPerShare'].iloc[itt],X['CapexPerShare'].iloc[ittPrevious],X['YearOverYearRateCapexPerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateCapexPerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'CapexPerShare'], X['CapexPerShare'].iloc[itt])
    return future_data



def prepare_and_predict(model, X):
    # Convert all columns to numeric, handling non-numeric gracefully
    for column in X.columns:
        X[column] = pd.to_numeric(X[column], errors='coerce')
    # Handle NaN values
    X.fillna(X.mean(), inplace=True)

    # Check data types
    print(X.dtypes)

    # Create DMatrix
    future_data_dmatrix = xgb.DMatrix(X)

    # Make predictions
    predictions = model.predict(future_data_dmatrix)
    return predictions


def predict_one_year_out(model, X,future_predicted_data):
    # Initialize an empty DataFrame to hold all predictions
    all_predicted_data = pd.DataFrame()
    
    # Create a DataFrame for the future data
    predicted_data = pd.DataFrame(columns=X.columns)

    start_index = len(X) - 1  # Last index of the DataFrame

    # Loop over the dataset from 'start_index' back to the start of the DataFrame every 90 days
    for itt in range(start_index, -1, -5):  # Decrement by 5
        # Ensure the index does not go out of bounds
        if itt < len(X):
            currentDate = X.index[itt]
            # Call the future data frame calculation
            predicted_data = future_data_frame(currentDate, X, itt)
            
            # if itt% 100: 
            #     print(predicted_data)
            
            # Append the calculated future data for prediction
            all_predicted_data = pd.concat([all_predicted_data, predicted_data], axis=0)

        
    # Create DMatrix for XGBoost prediction

    predictions = prepare_and_predict(model, all_predicted_data)
    return (pd.Series(predictions, index=all_predicted_data.index),all_predicted_data)

# Usage of the function
future_predicted_data = pd.DataFrame()
predictionTuple = predict_one_year_out(model, X, future_predicted_data)
esitmated_one_year_predictions = predictionTuple[0]
future_predicted_data = pd.concat([future_predicted_data, predictionTuple[1]], axis=0)
###########################################################################
# Validate function prints properly 
print(esitmated_one_year_predictions)
print(val_predictions) 
# print(future_predicted_data)

###########################################################################

# Ensure dates are sorted 
X_val_sorted = X_val.sort_index()
y_sorted = y_val.sort_index(ascending=True)
z_sorted = esitmated_one_year_predictions.sort_index()

# print(X_val_sorted)
print(X_val.index.equals(y_val.index))  
print(y_sorted)
print('z_sorted')
print(z_sorted)
# Print the first few rows of the actual data to confirm it's correct
print(y_val.head())
# start_date = '2025-02-25'
# print(z_sorted[start_date])

# ###################################################################

# Plotting both actual and predicted values
plt.figure(figsize=(14, 7))
plt.plot(X_val_sorted.index, y_sorted, label='Actual', marker='o', color='blue')
plt.plot(X_val.index, val_predictions, label='Predicted Close Price', marker='x', linestyle='--', color='red')
plt.plot(z_sorted.index, esitmated_one_year_predictions, label='1 year out - Precidted Price wth generated data', marker='x', linestyle='--', color='purple')
plt.title('Comparison of Actual and Predicted Values')
plt.xlabel('Date')
plt.ylabel('Value')
plt.legend()
plt.grid(True)
plt.show()

########################################################################

plt.figure(figsize=(14, 7))
plt.plot(y_val.index, y_val, label='Actual Close Price', marker='o', color='blue')
plt.title('Actual Close Prices')
plt.xlabel('Date')
plt.ylabel('Close Price')
plt.grid(True)
plt.show()

plt.figure(figsize=(14, 7))
plt.plot(X_val.index, val_predictions, label='Predicted Close Price', marker='x', linestyle='--', color='red')
plt.title('Predicted Close Prices')
plt.xlabel('Date')
plt.ylabel('Close Price')
plt.grid(True)
plt.show()


plt.figure(figsize=(14, 7))
plt.plot(z_sorted.index, esitmated_one_year_predictions, label='1 year out Close Price', marker='x', linestyle='--', color='purple')
plt.title('Predicted Close Prices')
plt.xlabel('Date')
plt.ylabel('Close Price')
plt.grid(True)
plt.show()
#################################################################

xgb.plot_importance(model, importance_type='gain')
plt.show()

importances = model.get_score(importance_type='gain')

# 'importances' is a dictionary where the key is the feature name and the value is the score

sorted_importances = sorted(importances.items(), key=lambda x: x[1], reverse=True)
feature_names_sorted = [item[0] for item in sorted_importances]
top_five_feature_names_sorted = feature_names_sorted[:5]
print(top_five_feature_names_sorted)

#################################################################

# store feature imporance across database, maybe ranking not actual value 

def median_future_periods(data, months, start_date):
    print(start_date)
    median_values = []
    for m in months:
        date = start_date + timedelta(days=30*m) 
        # Define the start date and the window size

        window_size = 3  # days on each side

        # Calculate the range of dates
        start_range = pd.to_datetime(date) - pd.Timedelta(days=window_size)
        end_range = pd.to_datetime(date) + pd.Timedelta(days=window_size)
        
        # Filter the DataFrame for this range
        date_range_data = data.loc[start_range:end_range]
        # Calculate the median of the values in this date range
        median_value = date_range_data.median()
        updatedDate = str(date)
        updatedDate =  updatedDate[:10]
        print(updatedDate)
        median_values.append((updatedDate,median_value))
    return median_values
        

# Predict future periods
months_ahead = [3, 6, 9, 12]
start_date = X.index[0]
four_predictions = median_future_periods(esitmated_one_year_predictions, months_ahead,start_date)
print(four_predictions)

#################################################################

import math

def median_backTest_periods(data, months):
    median_values = []
    for date in months:
        window_size = 3  # days on each side

        # Calculate the range of dates
        start_range = pd.to_datetime(date) - pd.Timedelta(days=window_size)
        end_range = pd.to_datetime(date) + pd.Timedelta(days=window_size)
        
        # Filter the DataFrame for this range
        date_range_data = data.loc[start_range:end_range]
        # print(date_range_data)
        # Calculate the median of the values in this date range
        median_value = date_range_data.median()
       
        if math.isnan(median_value):
            continue
        median_values.append((date,median_value))
       
    return median_values
        

# Predict future periods
backTestDates = ["2012-02-14", "2012-05-14", "2012-08-14", "2012-11-14","2013-02-14", "2013-05-14", "2013-08-14", "2013-11-14", 
                "2014-02-14", "2014-05-14", "2014-08-14", "2014-11-14","2015-02-13", "2015-05-14", "2015-08-14", "2015-11-13",
                "2016-02-12", "2016-05-13", "2016-08-15", "2016-11-14","2017-02-14", "2017-05-15", "2017-08-14", "2017-11-14",
                "2018-02-14", "2018-05-14", "2018-08-14", "2018-11-14","2019-02-14", "2019-05-14", "2019-08-14", "2019-11-14",
                "2020-02-14", "2020-05-14", "2020-08-14", "2020-11-13","2021-02-12", "2021-05-14", "2021-08-13", "2021-11-15",
                "2022-02-14", "2022-05-13", "2022-08-15", "2022-11-14","2023-02-14", "2023-05-12", "2023-08-14", "2023-11-14",
                "2024-02-14", "2024-05-13"
               ]
start_date = X.index[0]
backTestPredictions = median_backTest_periods(esitmated_one_year_predictions, backTestDates)
print(backTestPredictions)

#################################################################


def Calculate_Best_Error(actual, predicted, median, mean,belowCounter):
    mean_predicted = predicted + (mean * belowCounter)
    median_predicted = predicted + (median * belowCounter)

    # actual and predicted are arrays or series of the actual and predicted values
    mae_raw_predicted = mean_absolute_error(actual, predicted)
    mae_median_predicted = mean_absolute_error(actual, median_predicted)
    mae_mean_predicted = mean_absolute_error(actual, mean_predicted)
    
    print("Mean Absolute Error Raw:", mae_raw_predicted)
    print("Mean Absolute Error Median:", mae_median_predicted)
    print("Mean Absolute Error Mean:", mae_mean_predicted)
    
    return(mae_raw_predicted, mae_median_predicted, mae_mean_predicted )

def median_distance_Actual_to_Predicted(actual, predicted):
    end_date = actual.index[len(actual) -1]
    start_date = predicted.index[0]
    # start_date = '2010-01-12'
    
    # count how often prediction is below the line 
    belowCounter = 0
    ActualToPredictedDistance = []
    
    # Filter the DataFrame for this range
    actual_pruned = actual.loc[start_date:end_date]
    predicted_pruned = predicted.loc[start_date:end_date]
    
    for i in range (len(actual_pruned)):
        if actual_pruned.iloc[i] >  predicted_pruned.iloc[i]:
            belowCounter += 1
        tempDistance = actual_pruned.iloc[i] - predicted_pruned.iloc[i]
        ActualToPredictedDistance.append(tempDistance)
        # print("ran", i, actual_pruned.iloc[i], predicted_pruned.iloc[i])
    # Calculate the median of the values in this date range
    ActualToPredictedDistance.sort()
    medianPoint = round(len(actual) / 2)
    averageDistance = sum(ActualToPredictedDistance) / len(ActualToPredictedDistance)
    belowCounter = belowCounter / len(actual_pruned)
    # when belowCounter is close to 1 or 0 its often bewlow or above often. True uncertainty is not at 1 but at .5
    meanCounter = belowCounter
    if belowCounter > .5:
        # if the counter is to close to 1 or 0 the score there will be a higher mean and median. This is to help counter act that
        meanCounter = 1 - meanCounter
    
    error_calculations = Calculate_Best_Error(actual, predicted, ActualToPredictedDistance[medianPoint], averageDistance, meanCounter)
    # error_calculations2 = Calculate_Best_Error(actual, predicted, ActualToPredictedDistance[medianPoint], averageDistance, belowCounter)
    # error_calculations3 = Calculate_Best_Error(actual, predicted, ActualToPredictedDistance[medianPoint] / 2, averageDistance / 2, meanCounter)
    return (ActualToPredictedDistance[medianPoint], averageDistance, belowCounter, error_calculations)

# Predict future periods
medianDistance = median_distance_Actual_to_Predicted(y_sorted, z_sorted)
print(medianDistance)

#################################################################

predictionList = backTestPredictions + four_predictions
print(predictionList)
#################################################################
def median_distance_Actual_to_Predicted_3_years(actual, predicted, dates):
    # end_date = actual.index[len(actual) -1]
    # start_date = predicted.index[0]
    # start_date = '2010-01-12'
    three_year_out_prediction_Flag = False
    first_three_year_out_prediction_date = ''
    
    new_predictionList = []
    for date in dates:
        end_date = pd.to_datetime(date[0])
        end_date = end_date - timedelta(days=365)
        lastIndex = actual[:end_date]
        # print("last Index:", lastIndex)
        end_date = lastIndex.index[len(lastIndex) -1]
        # print(repr(end_date))
        position_in_y_val = actual.index.get_loc(end_date)
        start_date = actual.index[position_in_y_val - 151]
        
        # print("start_date:", start_date)
        # print("end_date:", end_date)
        # count how often prediction is below the line 
        belowCounter = 0
        ActualToPredictedDistance = []

        # Filter the DataFrame for this range
        actual_pruned = actual.loc[start_date:end_date]
        predicted_pruned = predicted.loc[start_date:end_date]
        # print("actual_pruned length:",len(actual_pruned))
        # print("actual_pruned:", actual_pruned)
        # print("predicted_pruned length:",len(predicted_pruned))
        # print("predicted_pruned:", predicted_pruned)
        if len(predicted_pruned) == 0:
            # print("NOT enough DATA yet to Back Test Data")
            continue
        if not three_year_out_prediction_Flag:
            three_year_out_prediction_Flag = True
            first_three_year_out_prediction_date = end_date
        actual_pruned, predicted_pruned = truncate_to_shorter_list(actual_pruned,predicted_pruned)
        for i in range (len(actual_pruned)):
            if actual_pruned.iloc[i] > predicted_pruned.iloc[i]:
                belowCounter += 1
            tempDistance = actual_pruned.iloc[i] - predicted_pruned.iloc[i]
            ActualToPredictedDistance.append(tempDistance)
            # print("ran", i, actual_pruned.iloc[i], predicted_pruned.iloc[i])
        # Calculate the median of the values in this date range
        ActualToPredictedDistance.sort()
        medianPoint = round(len(actual_pruned) / 2)
        averageDistance = sum(ActualToPredictedDistance) / len(ActualToPredictedDistance)
        belowCounter = belowCounter / len(actual_pruned)
        # when belowCounter is close to 1 or 0 its often bewlow or above often. True uncertainty is not at 1 but at .5
        meanCounter = belowCounter
        if belowCounter > .5:
            # if the counter is to close to 1 or 0 the score there will be a higher mean and median. This is to help counter act that
            meanCounter = 1 - meanCounter

        error_calculations = Calculate_Best_Error(actual_pruned, predicted_pruned, ActualToPredictedDistance[medianPoint], averageDistance, meanCounter)
        # error_calculations2 = Calculate_Best_Error(actual, predicted, ActualToPredictedDistance[medianPoint], averageDistance, belowCounter)
        # error_calculations3 = Calculate_Best_Error(actual, predicted, ActualToPredictedDistance[medianPoint] / 2, averageDistance / 2, meanCounter)
        prediction = (ActualToPredictedDistance[medianPoint], averageDistance, belowCounter, error_calculations, first_three_year_out_prediction_date)
        new_predictionList.append(prediction)
    return new_predictionList


# Predict future periods
updatedMedianList = median_distance_Actual_to_Predicted_3_years(y_sorted, z_sorted, predictionList)
print(updatedMedianList)
print(len(updatedMedianList))

#################################################################
# slicing off the difference between the predictionList and 3 year sliding if not enough data 

frontSlice = len(predictionList) - len(updatedMedianList)
if frontSlice != 0:
    predictionList = predictionList[frontSlice:]
print(predictionList)
print(len(predictionList))
print(len(updatedeMdianList))

#################################################################

import boto3
from botocore.exceptions import ClientError
from decimal import Decimal

def float_to_decimal(value):
    if value is None:
        return None
    return Decimal(str(value))

def prediction_batch_insert(symbol, predictionList,medianDistance, top_five_feature_names_sorted):
    table_name ='stock-predictions'
    # Create a DynamoDB client
    dynamodb = boto3.resource('dynamodb', region_name='us-west-1')
    table = dynamodb.Table(table_name)

    with table.batch_writer() as batch:
        for i, prediction in enumerate(predictionList):
            try: 
                batch.put_item(Item={
                    "symbol": symbol,
                    "date": str(prediction[0]),
                    "price": float_to_decimal(prediction[1]),
                    "features": top_five_feature_names_sorted,
                    "bias": float_to_decimal(medianDistance[i][2]),
                    "medianDistance": float_to_decimal(medianDistance[i][0]),
                    "meanDistance": float_to_decimal(medianDistance[i][1]),
                    "defaultError": float_to_decimal(medianDistance[i][3][0]),
                    "medianError": float_to_decimal(medianDistance[i][3][1]),
                    "meanError": float_to_decimal(medianDistance[i][3][2]),
                })
            except ClientError as e:
                print("Failed to insert item:", e)

prediction_batch_insert(symbol, predictionList, updatedMedianList, top_five_feature_names_sorted)
print("uploaded:", symbol)

#################################################################