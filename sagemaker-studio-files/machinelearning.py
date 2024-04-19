!!pip install --upgrade pandas xgboost
!pip install --force-reinstall statsmodels==0.14.0 pydeseq2==0.4.3

# Import necessary libraries
import boto3
import xgboost as xgb
import pandas as pd
from sklearn.model_selection import train_test_split
from sagemaker import get_execution_role
from sagemaker.inputs import TrainingInput
from sagemaker.xgboost.estimator import XGBoost
import matplotlib.pyplot as plt
from datetime import timedelta

# Setup and configuration
role = get_execution_role()
bucket = 'stock-trading-bucket-crimson'
data_key = 'raw-data/AAPL.csv'
data_location = f's3://{bucket}/{data_key}'
# ############################################################################

# Load data from S3
s3_client = boto3.client('s3')
obj = s3_client.get_object(Bucket=bucket, Key=data_key)
data = pd.read_csv(obj['Body'])

# Convert 'Datetime' to datetime and set it as index
data['Datetime'] = pd.to_datetime(data['Datetime'])
data.set_index('Datetime', inplace=True)

# # Print the column names to verify them
# print(data.columns)


# Prepare the data
# Columns to drop
columns_to_drop = ['High', 'Low', 'Open', 'Period', 'Symbol', 'MovingAverage50Days', 'MovingAverage200Days', 'MarketCap']

# Dropping specified columns
X = data.drop(columns_to_drop + ['Close'], axis=1)
y = data['Close']


# Assuming 'X' and 'y' are already defined and Datetime-indexed
interval = 5  # Change this based on your desired validation set size

# Create boolean mask
mask = [True if i % interval == 0 else False for i in range(len(X))]

# Split the data using the mask
X_train = X[~pd.Series(mask, index=X.index)]
X_val = X[pd.Series(mask, index=X.index)]
y_train = y[~pd.Series(mask, index=y.index)]
y_val = y[pd.Series(mask, index=y.index)]

# Print the sizes of each dataset to verify the split
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


def stock_rate_one_year_position_calculator(finalPosition, previousPosition):
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
    velocityCurrent  = currentPosition / currentVelocityRate
    velocityOld =  oldPosition / oldVelocityRate
    acceleration = velocityCurrent - velocityOld

    predictedPosition = currentPosition + currentVelocityRate + (.5 * acceleration)

    return predictedPosition


def future_data_frame(currentDate, X, itt):
    one_year_out = currentDate + timedelta(days=365)
    future_data = pd.DataFrame(index=[one_year_out], columns=X.columns)
    ittPrevious = itt -365
    
    future_data.loc[one_year_out,'Volume'] = stock_position_no_acceleration_calculator(X['Volume'].iloc[itt],X['Volume'].iloc[ittPrevious])
    future_data.loc[one_year_out,'OperatingCashFlowPerShare'] = stock_one_year_position_calculator(X['OperatingCashFlowPerShare'].iloc[itt],X['YearOverYearRateOperatingCashFlowPerShare'].iloc[itt],X['OperatingCashFlowPerShare'].iloc[ittPrevious],X['YearOverYearRateOperatingCashFlowPerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateOperatingCashFlowPerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'OperatingCashFlowPerShare'], X['OperatingCashFlowPerShare'].iloc[itt])
    future_data.loc[one_year_out,'FreeCashFlowPerShare'] = stock_one_year_position_calculator(X['FreeCashFlowPerShare'].iloc[itt],X['YearOverYearRateFreeCashFlowPerShare'].iloc[itt],X['FreeCashFlowPerShare'].iloc[ittPrevious],X['YearOverYearRateFreeCashFlowPerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateFreeCashFlowPerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'FreeCashFlowPerShare'], X['FreeCashFlowPerShare'].iloc[itt])
    future_data.loc[one_year_out,'CashPerShare'] = stock_position_no_acceleration_calculator(X['CashPerShare'].iloc[itt],X['CashPerShare'].iloc[ittPrevious])
    future_data.loc[one_year_out,'PriceToSalesRatio'] = stock_position_no_acceleration_calculator(X['PriceToSalesRatio'].iloc[itt],X['PriceToSalesRatio'].iloc[ittPrevious])
    future_data.loc[one_year_out,'PayoutRatio'] = stock_position_no_acceleration_calculator(X['PayoutRatio'].iloc[itt],X['PayoutRatio'].iloc[ittPrevious])
    future_data.loc[one_year_out,'RevenuePerShare'] = stock_one_year_position_calculator(X['RevenuePerShare'].iloc[itt],X['YearOverYearRateRevenuePerShare'].iloc[itt],X['RevenuePerShare'].iloc[ittPrevious],X['YearOverYearRateRevenuePerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateRevenuePerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'RevenuePerShare'], X['RevenuePerShare'].iloc[itt])
    future_data.loc[one_year_out,'BookValuePerShare'] = stock_one_year_position_calculator(X['BookValuePerShare'].iloc[itt],X['YearOverYearRateBookValuePerShare'].iloc[itt],X['BookValuePerShare'].iloc[ittPrevious],X['YearOverYearRateBookValuePerShare'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateBookValuePerShare'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'BookValuePerShare'], X['BookValuePerShare'].iloc[itt])
    future_data.loc[one_year_out,'PeRatio'] = stock_position_no_acceleration_calculator(X['PeRatio'].iloc[itt],X['PeRatio'].iloc[ittPrevious])
    future_data.loc[one_year_out,'PfcfRatio'] = stock_position_no_acceleration_calculator(X['PfcfRatio'].iloc[itt],X['PfcfRatio'].iloc[ittPrevious])
    future_data.loc[one_year_out,'EvToOperatingCashFlow'] = stock_one_year_position_calculator(X['EvToOperatingCashFlow'].iloc[itt],X['YearOverYearRateEvToOperatingCashFlow'].iloc[itt],X['EvToOperatingCashFlow'].iloc[ittPrevious],X['YearOverYearRateEvToOperatingCashFlow'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateEvToOperatingCashFlow'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'EvToOperatingCashFlow'], X['EvToOperatingCashFlow'].iloc[itt])
    future_data.loc[one_year_out,'NetDebtToEBITDA'] = stock_one_year_position_calculator(X['NetDebtToEBITDA'].iloc[itt],X['YearOverYearRateNetDebtToEBITDA'].iloc[itt],X['NetDebtToEBITDA'].iloc[ittPrevious],X['YearOverYearRateNetDebtToEBITDA'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateNetDebtToEBITDA'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'NetDebtToEBITDA'], X['NetDebtToEBITDA'].iloc[itt])
    future_data.loc[one_year_out,'StockBasedCompensationToRevenue'] = stock_position_no_acceleration_calculator(X['StockBasedCompensationToRevenue'].iloc[itt],X['StockBasedCompensationToRevenue'].iloc[ittPrevious])
    future_data.loc[one_year_out,'GrahamNumber'] = stock_one_year_position_calculator(X['GrahamNumber'].iloc[itt],X['YearOverYearRateGrahamNumber'].iloc[itt],X['GrahamNumber'].iloc[ittPrevious],X['YearOverYearRateGrahamNumber'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateGrahamNumber'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'GrahamNumber'], X['GrahamNumber'].iloc[itt])
    future_data.loc[one_year_out,'Roic'] = stock_one_year_position_calculator(X['Roic'].iloc[itt],X['YearOverYearRateRoic'].iloc[itt],X['Roic'].iloc[ittPrevious],X['YearOverYearRateRoic'].iloc[ittPrevious] )
    future_data.loc[one_year_out,'YearOverYearRateRoic'] = stock_rate_one_year_position_calculator(future_data.loc[one_year_out,'Roic'], X['Roic'].iloc[itt])
    future_data.loc[one_year_out,'Roe'] = stock_position_no_acceleration_calculator(X['Roe'].iloc[itt],X['Roe'].iloc[ittPrevious])
    future_data.loc[one_year_out,'CapexPerShare'] = stock_position_no_acceleration_calculator(X['CapexPerShare'].iloc[itt],X['CapexPerShare'].iloc[ittPrevious])
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


def predict_one_year_out(model, X):
    # Initialize an empty DataFrame to hold all predictions
    all_predicted_data = pd.DataFrame()
    
    # Create a DataFrame for the future data
    predicted_data = pd.DataFrame(columns=X.columns)

    start_index = len(X) - 1  # Last index of the DataFrame

    # Loop over the dataset from 'start_index' back to the start of the DataFrame every 90 days
    for itt in range(start_index, -1, -5):  # Decrement by 90
        # Ensure the index does not go out of bounds
        if itt < len(X):
            currentDate = X.index[itt]
            # Call the future data frame calculation
            predicted_data = future_data_frame(currentDate, X, itt)

            # Append the calculated future data for prediction
            all_predicted_data = pd.concat([all_predicted_data, predicted_data], axis=0)

    
    # Create DMatrix for XGBoost prediction

    predictions = prepare_and_predict(model, all_predicted_data)
    return pd.Series(predictions, index=all_predicted_data.index)


# Usage of the function
esitmated_one_year_predictions = predict_one_year_out(model, X )
###########################################################################

# Ensure dates are sorted (if your data isn't time series, you might skip indexing by date)
X_val_sorted = X_val.sort_index()
y_sorted = y_val.reindex(X_val_sorted.index) # getting only Y values 
z_sorted = esitmated_one_year_predictions.sort_index()

# print(X_val_sorted)
print(X_val.index.equals(y_val.index))  
print(y_sorted)
print('z_sorted')
print(z_sorted)
# Print the first few rows of the actual data to confirm it's correct
print(y_val.head())



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


