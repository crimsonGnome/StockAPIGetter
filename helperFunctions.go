package main

import (
	"encoding/csv"
	"fmt"
	"os"
	env "stockpulling/main/env"
	"strconv"
	"time"
)

func top25() {
	for _, stockSymbol := range env.StockList {
		generateSingleStockCSV(stockSymbol)
	}
}

func apiToStockDataStruct(historicStockFinancialsArray *[]HistoricStockFinancials, dailyStockPriceArray *BodyDailyStockData, stockSymbol string) *[]StockData {
	var stockDataArray []StockData

	// Iterator used to track which quarter the date respond to
	quarterlyReportCounter := 0
	// Date in which quarterly report was added
	dateStringQuarterly := (*historicStockFinancialsArray)[quarterlyReportCounter].Date
	// format date

	datePreviousQuarterly, err := time.Parse("2006-01-02", dateStringQuarterly)
	if err != nil {
		fmt.Println(err)
	}

	// forLoop over dailyStock Prices
	for _, dailyStockPrice := range (*dailyStockPriceArray).Values {

		// convert string data into Floats
		close, err := strconv.ParseFloat(dailyStockPrice.Close, 64)
		if err != nil {
			fmt.Println(err)
		}

		high, err := strconv.ParseFloat(dailyStockPrice.High, 64)
		if err != nil {
			fmt.Println(err)
		}

		low, err := strconv.ParseFloat(dailyStockPrice.Low, 64)
		if err != nil {
			fmt.Println(err)
		}

		open, err := strconv.ParseFloat(dailyStockPrice.Open, 64)
		if err != nil {
			fmt.Println(err)
		}

		volume, err := strconv.ParseFloat(dailyStockPrice.Volume, 64)
		if err != nil {
			fmt.Println(err)
		}

		// - add Daily stock values to meta data
		currentStock := StockData{
			Symbol:   stockSymbol,
			Datetime: dailyStockPrice.Datetime,
			Close:    close,
			High:     high,
			Low:      low,
			Open:     open,
			Volume:   volume,
		}
		// check if dailyStock date is >= Historic Financial Date
		dateStringDaily := currentStock.Datetime

		// Convert current date TIme string into a Date time
		dateTimeDaily, err := time.Parse("2006-01-02", dateStringDaily)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Sprintln(" dateTimeDaily VS  datePreviousQuarterly")
		fmt.Sprintln(dateTimeDaily)

		// Compare if the quartely report is active
		// printing Comparisons

		fmt.Sprintln(datePreviousQuarterly)

		for dateTimeDaily.Unix() <= datePreviousQuarterly.Unix() {
			// Check to see if historical array is in range
			if len((*historicStockFinancialsArray)) <= (quarterlyReportCounter + 2) {
				break
			}
			// Set the previous Quarter Financial data to current Quarter
			quarterlyReportCounter = quarterlyReportCounter + 1
			dateStringQuarterly = (*historicStockFinancialsArray)[quarterlyReportCounter].Date
			datePreviousQuarterly, err = time.Parse("2006-01-02", dateStringQuarterly)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Copy data metrics into currentStock
		currentStock.Period = (*historicStockFinancialsArray)[quarterlyReportCounter].Period
		currentStock.OperatingCashFlowPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].OperatingCashFlowPerShare
		currentStock.FreeCashFlowPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].FreeCashFlowPerShare
		currentStock.CashPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].CashPerShare
		currentStock.PriceToSalesRatio = (*historicStockFinancialsArray)[quarterlyReportCounter].PriceToSalesRatio
		currentStock.DividendYield = (*historicStockFinancialsArray)[quarterlyReportCounter].DividendYield
		currentStock.PayoutRatio = (*historicStockFinancialsArray)[quarterlyReportCounter].PayoutRatio
		currentStock.RevenuePerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].RevenuePerShare
		currentStock.NetIncomePerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].NetIncomePerShare
		currentStock.BookValuePerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].BookValuePerShare
		currentStock.ShareholdersEquityPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].ShareholdersEquityPerShare
		currentStock.InterestDebtPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].InterestDebtPerShare
		currentStock.MarketCap = (*historicStockFinancialsArray)[quarterlyReportCounter].MarketCap
		currentStock.EnterpriseValue = (*historicStockFinancialsArray)[quarterlyReportCounter].EnterpriseValue
		currentStock.PeRatio = (*historicStockFinancialsArray)[quarterlyReportCounter].PeRatio
		currentStock.Pocfratio = (*historicStockFinancialsArray)[quarterlyReportCounter].Pocfratio
		currentStock.PfcfRatio = (*historicStockFinancialsArray)[quarterlyReportCounter].PfcfRatio
		currentStock.Pbratio = (*historicStockFinancialsArray)[quarterlyReportCounter].Pbratio
		currentStock.PtbRatio = (*historicStockFinancialsArray)[quarterlyReportCounter].PtbRatio
		currentStock.EvToSales = (*historicStockFinancialsArray)[quarterlyReportCounter].EvToSales
		currentStock.EnterpriseValueOverEBITDA = (*historicStockFinancialsArray)[quarterlyReportCounter].EnterpriseValueOverEBITDA
		currentStock.EvToOperatingCashFlow = (*historicStockFinancialsArray)[quarterlyReportCounter].EvToOperatingCashFlow
		currentStock.EarningsYield = (*historicStockFinancialsArray)[quarterlyReportCounter].EarningsYield
		currentStock.FreeCashFlowYield = (*historicStockFinancialsArray)[quarterlyReportCounter].FreeCashFlowYield
		currentStock.DebtToEquity = (*historicStockFinancialsArray)[quarterlyReportCounter].DebtToEquity
		currentStock.DebtToAssets = (*historicStockFinancialsArray)[quarterlyReportCounter].DebtToAssets
		currentStock.NetDebtToEBITDA = (*historicStockFinancialsArray)[quarterlyReportCounter].NetDebtToEBITDA
		currentStock.CurrentRatio = (*historicStockFinancialsArray)[quarterlyReportCounter].CurrentRatio
		currentStock.InterestCoverage = (*historicStockFinancialsArray)[quarterlyReportCounter].InterestCoverage
		currentStock.IncomeQuality = (*historicStockFinancialsArray)[quarterlyReportCounter].IncomeQuality
		currentStock.SalesGeneralAndAdministrativeToRevenue = (*historicStockFinancialsArray)[quarterlyReportCounter].SalesGeneralAndAdministrativeToRevenue
		currentStock.ResearchAndDevelopmentToRevenue = (*historicStockFinancialsArray)[quarterlyReportCounter].ResearchAndDevelopmentToRevenue
		currentStock.IntangiblesToTotalAssets = (*historicStockFinancialsArray)[quarterlyReportCounter].IntangiblesToTotalAssets
		currentStock.CapexToOperatingCashFlow = (*historicStockFinancialsArray)[quarterlyReportCounter].CapexToOperatingCashFlow
		currentStock.CapexToRevenue = (*historicStockFinancialsArray)[quarterlyReportCounter].CapexToRevenue
		currentStock.CapexToDepreciation = (*historicStockFinancialsArray)[quarterlyReportCounter].CapexToDepreciation
		currentStock.StockBasedCompensationToRevenue = (*historicStockFinancialsArray)[quarterlyReportCounter].StockBasedCompensationToRevenue
		currentStock.GrahamNumber = (*historicStockFinancialsArray)[quarterlyReportCounter].GrahamNumber
		currentStock.Roic = (*historicStockFinancialsArray)[quarterlyReportCounter].Roic
		currentStock.ReturnOnTangibleAssets = (*historicStockFinancialsArray)[quarterlyReportCounter].ReturnOnTangibleAssets
		currentStock.GrahamNetNet = (*historicStockFinancialsArray)[quarterlyReportCounter].GrahamNetNet
		currentStock.WorkingCapital = (*historicStockFinancialsArray)[quarterlyReportCounter].WorkingCapital
		currentStock.TangibleAssetValue = (*historicStockFinancialsArray)[quarterlyReportCounter].TangibleAssetValue
		currentStock.NetCurrentAssetValue = (*historicStockFinancialsArray)[quarterlyReportCounter].NetCurrentAssetValue
		currentStock.InvestedCapital = (*historicStockFinancialsArray)[quarterlyReportCounter].InvestedCapital
		currentStock.AverageReceivables = (*historicStockFinancialsArray)[quarterlyReportCounter].AverageReceivables
		currentStock.AveragePayables = (*historicStockFinancialsArray)[quarterlyReportCounter].AveragePayables
		currentStock.AverageInventory = (*historicStockFinancialsArray)[quarterlyReportCounter].AverageInventory
		currentStock.DaysSalesOutstanding = (*historicStockFinancialsArray)[quarterlyReportCounter].DaysSalesOutstanding
		currentStock.DaysPayablesOutstanding = (*historicStockFinancialsArray)[quarterlyReportCounter].DaysPayablesOutstanding
		currentStock.DaysOfInventoryOnHand = (*historicStockFinancialsArray)[quarterlyReportCounter].DaysOfInventoryOnHand
		currentStock.ReceivablesTurnover = (*historicStockFinancialsArray)[quarterlyReportCounter].ReceivablesTurnover
		currentStock.PayablesTurnover = (*historicStockFinancialsArray)[quarterlyReportCounter].PayablesTurnover
		currentStock.InventoryTurnover = (*historicStockFinancialsArray)[quarterlyReportCounter].InventoryTurnover
		currentStock.Roe = (*historicStockFinancialsArray)[quarterlyReportCounter].Roe
		currentStock.CapexPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].CapexPerShare

		stockDataArray = append(stockDataArray, currentStock)
	}

	return &stockDataArray

}

func CSVconverter2(stockSymbol string, stockDataArray *[]StockDataImproved2) error {
	destinationFileName := fmt.Sprintf("%s2.csv", stockSymbol)

	outputFile, err := os.Create(destinationFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{
		"Symbol",
		"Datetime",
		"Close",
		"High",
		"Low",
		"Open",
		"Volume",
		"Period",
		"OperatingCashFlowPerShare",
		"YearOverYearRateOperatingCashFlowPerShare",
		"FreeCashFlowPerShare",
		"YearOverYearRateFreeCashFlowPerShare",
		"CashPerShare",
		"YearOverYearRateCashPerShare",
		"PriceToSalesRatio",
		"YearOverYearRatePriceToSalesRatio",
		"PayoutRatio",
		"YearOverYearRatePayoutRatio",
		"RevenuePerShare",
		"YearOverYearRateRevenuePerShare",
		"BookValuePerShare",
		"YearOverYearRateBookValuePerShare",
		"MarketCap",
		"PeRatio",
		"YearOverYearRatePeRatio",
		"PfcfRatio",
		"YearOverYearRatePfcfRatio",
		"EvToOperatingCashFlow",
		"YearOverYearRateEvToOperatingCashFlow",
		"NetDebtToEBITDA",
		"YearOverYearRateNetDebtToEBITDA",
		"StockBasedCompensationToRevenue",
		"YearOverYearRateStockBasedCompensationToRevenue",
		"GrahamNumber",
		"YearOverYearRateGrahamNumber",
		"Roic",
		"YearOverYearRateRoic",
		"Roe",
		"YearOverYearRateRoe",
		"CapexPerShare",
		"YearOverYearRateCapexPerShare",
		"MovingAverage50Days",
		"MovingAverage200Days",
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range *stockDataArray {
		var csvRow []string
		csvRow = append(csvRow,
			r.Symbol,
			r.Datetime,
			fmt.Sprintf("%f", r.Close),
			fmt.Sprintf("%f", r.High),
			fmt.Sprintf("%f", r.Low),
			fmt.Sprintf("%f", r.Open),
			fmt.Sprintf("%f", r.Volume),
			r.Period,
			fmt.Sprintf("%f", r.OperatingCashFlowPerShare),
			fmt.Sprintf("%f", r.YearOverYearRateOperatingCashFlowPerShare),
			fmt.Sprintf("%f", r.FreeCashFlowPerShare),
			fmt.Sprintf("%f", r.YearOverYearRateFreeCashFlowPerShare),
			fmt.Sprintf("%f", r.CashPerShare),
			fmt.Sprintf("%f", r.YearOverYearRateCashPerShare),
			fmt.Sprintf("%f", r.PriceToSalesRatio),
			fmt.Sprintf("%f", r.YearOverYearRatePriceToSalesRatio),
			fmt.Sprintf("%f", r.PayoutRatio),
			fmt.Sprintf("%f", r.YearOverYearRatePayoutRatio),
			fmt.Sprintf("%f", r.RevenuePerShare),
			fmt.Sprintf("%f", r.YearOverYearRateRevenuePerShare),
			fmt.Sprintf("%f", r.BookValuePerShare),
			fmt.Sprintf("%f", r.YearOverYearRateBookValuePerShare),
			fmt.Sprintf("%f", r.MarketCap),
			fmt.Sprintf("%f", r.PeRatio),
			fmt.Sprintf("%f", r.YearOverYearRatePeRatio),
			fmt.Sprintf("%f", r.PfcfRatio),
			fmt.Sprintf("%f", r.YearOverYearRatePfcfRatio),
			fmt.Sprintf("%f", r.EvToOperatingCashFlow),
			fmt.Sprintf("%f", r.YearOverYearRateEvToOperatingCashFlow),
			fmt.Sprintf("%f", r.NetDebtToEBITDA),
			fmt.Sprintf("%f", r.YearOverYearRateNetDebtToEBITDA),
			fmt.Sprintf("%f", r.StockBasedCompensationToRevenue),
			fmt.Sprintf("%f", r.YearOverYearRateStockBasedCompensationToRevenue),
			fmt.Sprintf("%f", r.GrahamNumber),
			fmt.Sprintf("%f", r.YearOverYearRateGrahamNumber),
			fmt.Sprintf("%f", r.Roic),
			fmt.Sprintf("%f", r.YearOverYearRateRoic),
			fmt.Sprintf("%f", r.Roe),
			fmt.Sprintf("%f", r.YearOverYearRateRoe),
			fmt.Sprintf("%f", r.CapexPerShare),
			fmt.Sprintf("%f", r.YearOverYearRateCapexPerShare),
			fmt.Sprintf("%f", r.MovingAverage50Days),
			fmt.Sprintf("%f", r.MovingAverage200Days),
		)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func CSVconverter(stockSymbol string, stockDataArray *[]StockDataImproved) error {
	destinationFileName := fmt.Sprintf("%s.csv", stockSymbol)

	outputFile, err := os.Create(destinationFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{
		"Symbol",
		"Datetime",
		"Close",
		"High",
		"Low",
		"Open",
		"Volume",
		"Period",
		"OperatingCashFlowPerShare",
		"YearOverYearRateOperatingCashFlowPerShare",
		"FreeCashFlowPerShare",
		"YearOverYearRateFreeCashFlowPerShare",
		"CashPerShare",
		"PriceToSalesRatio",
		"PayoutRatio",
		"RevenuePerShare",
		"YearOverYearRateRevenuePerShare",
		"BookValuePerShare",
		"YearOverYearRateBookValuePerShare",
		"MarketCap",
		"PeRatio",
		"PfcfRatio",
		"EvToOperatingCashFlow",
		"YearOverYearRateEvToOperatingCashFlow",
		"NetDebtToEBITDA",
		"YearOverYearRateNetDebtToEBITDA",
		"StockBasedCompensationToRevenue",
		"GrahamNumber",
		"YearOverYearRateGrahamNumber",
		"Roic",
		"YearOverYearRateRoic",
		"Roe",
		"CapexPerShare",
		"MovingAverage50Days",
		"MovingAverage200Days",
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range *stockDataArray {
		var csvRow []string
		csvRow = append(csvRow,
			r.Symbol,
			r.Datetime,
			fmt.Sprintf("%f", r.Close),
			fmt.Sprintf("%f", r.High),
			fmt.Sprintf("%f", r.Low),
			fmt.Sprintf("%f", r.Open),
			fmt.Sprintf("%f", r.Volume),
			r.Period,
			fmt.Sprintf("%f", r.OperatingCashFlowPerShare),
			fmt.Sprintf("%f", r.YearOverYearRateOperatingCashFlowPerShare),
			fmt.Sprintf("%f", r.FreeCashFlowPerShare),
			fmt.Sprintf("%f", r.YearOverYearRateFreeCashFlowPerShare),
			fmt.Sprintf("%f", r.CashPerShare),
			fmt.Sprintf("%f", r.PriceToSalesRatio),
			fmt.Sprintf("%f", r.PayoutRatio),
			fmt.Sprintf("%f", r.RevenuePerShare),
			fmt.Sprintf("%f", r.YearOverYearRateRevenuePerShare),
			fmt.Sprintf("%f", r.BookValuePerShare),
			fmt.Sprintf("%f", r.YearOverYearRateBookValuePerShare),
			fmt.Sprintf("%f", r.MarketCap),
			fmt.Sprintf("%f", r.PeRatio),
			fmt.Sprintf("%f", r.PfcfRatio),
			fmt.Sprintf("%f", r.EvToOperatingCashFlow),
			fmt.Sprintf("%f", r.YearOverYearRateEvToOperatingCashFlow),
			fmt.Sprintf("%f", r.NetDebtToEBITDA),
			fmt.Sprintf("%f", r.YearOverYearRateNetDebtToEBITDA),
			fmt.Sprintf("%f", r.StockBasedCompensationToRevenue),
			fmt.Sprintf("%f", r.GrahamNumber),
			fmt.Sprintf("%f", r.YearOverYearRateGrahamNumber),
			fmt.Sprintf("%f", r.Roic),
			fmt.Sprintf("%f", r.YearOverYearRateRoic),
			fmt.Sprintf("%f", r.Roe),
			fmt.Sprintf("%f", r.CapexPerShare),
			fmt.Sprintf("%f", r.MovingAverage50Days),
			fmt.Sprintf("%f", r.MovingAverage200Days),
		)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}
func YearOverYearArrayPositionCalculator(stockDataArray *[]StockData, currentIndex int) int {
	// 274 Days - calculate 1  year away
	//find next matching Quater
	YearOnYearIterator := currentIndex + 250

	if YearOnYearIterator >= len(*stockDataArray) {
		// -1 will be a break return, if negative 1 is returned end for loop array
		return -1
	}
	for i := YearOnYearIterator; i < len(*stockDataArray); i++ {
		// match Period to last previous years Period
		if (*stockDataArray)[i].Period == (*stockDataArray)[currentIndex].Period {
			return i
		}
	}
	// no year over Year period left in data
	return -1
}

func GrowthRateRatio(final float64, initial float64) float64 {
	// cant divide by 0;
	if initial == 0 {
		return 1
	}
	answer := final / initial

	return answer
}

func generateImprovedStockArray2(stockDataArray *[]StockData) *[]StockDataImproved2 {
	// var StockDataImprovedArray []StockDataImproved
	var StockDataImprovedArray2 []StockDataImproved2
	var movingAverage50Sum float64 = 0
	var movingAverage200Sum float64 = 0
	var tempMovingAverage50 float64 = 0
	var tempMovingAverage200 float64 = 0

	// Calculate 50 and 200 day moving averages
	for i := 0; i < 200; i++ {
		movingAverage200Sum = movingAverage200Sum + (*stockDataArray)[i].Close
		// Only get the first 50 to start the array
		if i < 50 {
			movingAverage50Sum = movingAverage50Sum + (*stockDataArray)[i].Close
		}
	}

	// forLoop over dailyStock Prices
	for i, current := range *stockDataArray {
		// Calculate the Moving averages
		tempMovingAverage50 = movingAverage50Sum / 50
		tempMovingAverage200 = movingAverage200Sum / 200
		YearOnYearIterator := YearOverYearArrayPositionCalculator(stockDataArray, i)
		if YearOnYearIterator == -1 {
			break
		}

		// Calculate Year over year increase Rate
		YearOverYearOperatingCashFlowPerShare := GrowthRateRatio(current.OperatingCashFlowPerShare, (*stockDataArray)[YearOnYearIterator].OperatingCashFlowPerShare)
		YearOverYearFreeCashFlowPerShare := GrowthRateRatio(current.FreeCashFlowPerShare, (*stockDataArray)[YearOnYearIterator].FreeCashFlowPerShare)
		YearOverYearRevenuePerShare := GrowthRateRatio(current.RevenuePerShare, (*stockDataArray)[YearOnYearIterator].RevenuePerShare)
		YearOverYearBookValuePerShare := GrowthRateRatio(current.BookValuePerShare, (*stockDataArray)[YearOnYearIterator].BookValuePerShare)
		YearOverYearEvToOperatingCashFlow := GrowthRateRatio(current.EvToOperatingCashFlow, (*stockDataArray)[YearOnYearIterator].EvToOperatingCashFlow)
		YearOverYearNetDebtToEBITDA := GrowthRateRatio(current.NetDebtToEBITDA, (*stockDataArray)[YearOnYearIterator].NetDebtToEBITDA)
		YearOverYearGrahamNumber := GrowthRateRatio(current.GrahamNumber, (*stockDataArray)[YearOnYearIterator].GrahamNumber)
		YearOverYearRoic := GrowthRateRatio(current.Roic, (*stockDataArray)[YearOnYearIterator].Roic)
		YearOverYearCashPerShare := GrowthRateRatio(current.CashPerShare, (*stockDataArray)[YearOnYearIterator].CashPerShare)
		YearOverYearPriceToSalesRatio := GrowthRateRatio(current.PriceToSalesRatio, (*stockDataArray)[YearOnYearIterator].PriceToSalesRatio)
		YearOverYearPayoutRatio := GrowthRateRatio(current.PayoutRatio, (*stockDataArray)[YearOnYearIterator].PayoutRatio)
		YearOverYearPfcfRatio := GrowthRateRatio(current.PfcfRatio, (*stockDataArray)[YearOnYearIterator].PfcfRatio)
		YearOverYearPeRatio := GrowthRateRatio(current.PeRatio, (*stockDataArray)[YearOnYearIterator].PeRatio)
		YearOverYearStockBasedCompensationToRevenue := GrowthRateRatio(current.StockBasedCompensationToRevenue, (*stockDataArray)[YearOnYearIterator].StockBasedCompensationToRevenue)
		YearOverYearRoe := GrowthRateRatio(current.Roe, (*stockDataArray)[YearOnYearIterator].Roe)
		YearOverYearCapexPerShare := GrowthRateRatio(current.CapexPerShare, (*stockDataArray)[YearOnYearIterator].CapexPerShare)

		currentStock2 := StockDataImproved2{
			Symbol:                    current.Symbol,
			Datetime:                  current.Datetime,
			Close:                     current.Close,
			High:                      current.High,
			Low:                       current.Low,
			Open:                      current.Open,
			Volume:                    current.Volume,
			Period:                    current.Period,
			OperatingCashFlowPerShare: current.OperatingCashFlowPerShare,
			YearOverYearRateOperatingCashFlowPerShare:       YearOverYearOperatingCashFlowPerShare,
			FreeCashFlowPerShare:                            current.FreeCashFlowPerShare,
			YearOverYearRateFreeCashFlowPerShare:            YearOverYearFreeCashFlowPerShare,
			CashPerShare:                                    current.CashPerShare,
			YearOverYearRateCashPerShare:                    YearOverYearCashPerShare,
			PriceToSalesRatio:                               current.PriceToSalesRatio,
			YearOverYearRatePriceToSalesRatio:               YearOverYearPriceToSalesRatio,
			PayoutRatio:                                     current.PayoutRatio,
			YearOverYearRatePayoutRatio:                     YearOverYearPayoutRatio,
			RevenuePerShare:                                 current.RevenuePerShare,
			YearOverYearRateRevenuePerShare:                 YearOverYearRevenuePerShare,
			BookValuePerShare:                               current.BookValuePerShare,
			YearOverYearRateBookValuePerShare:               YearOverYearBookValuePerShare,
			MarketCap:                                       current.MarketCap,
			PeRatio:                                         current.PeRatio,
			YearOverYearRatePeRatio:                         YearOverYearPeRatio,
			PfcfRatio:                                       current.PfcfRatio,
			YearOverYearRatePfcfRatio:                       YearOverYearPfcfRatio,
			EvToOperatingCashFlow:                           current.EvToOperatingCashFlow,
			YearOverYearRateEvToOperatingCashFlow:           YearOverYearEvToOperatingCashFlow,
			NetDebtToEBITDA:                                 current.NetDebtToEBITDA,
			YearOverYearRateNetDebtToEBITDA:                 YearOverYearNetDebtToEBITDA,
			StockBasedCompensationToRevenue:                 current.StockBasedCompensationToRevenue,
			YearOverYearRateStockBasedCompensationToRevenue: YearOverYearStockBasedCompensationToRevenue,
			GrahamNumber:                                    current.GrahamNumber,
			YearOverYearRateGrahamNumber:                    YearOverYearGrahamNumber,
			Roic:                                            current.Roic,
			YearOverYearRateRoic:                            YearOverYearRoic,
			Roe:                                             current.Roe,
			YearOverYearRateRoe:                             YearOverYearRoe,
			CapexPerShare:                                   current.CapexPerShare,
			YearOverYearRateCapexPerShare:                   YearOverYearCapexPerShare,
			MovingAverage50Days:                             tempMovingAverage50,
			MovingAverage200Days:                            tempMovingAverage200,
		}
		// Append Array
		StockDataImprovedArray2 = append(StockDataImprovedArray2, currentStock2)

		// Calculate new average
		movingAverage50Sum = movingAverage50Sum - current.Close + (*stockDataArray)[i+50].Close
		movingAverage200Sum = movingAverage200Sum - current.Close + (*stockDataArray)[i+200].Close

	}

	return &StockDataImprovedArray2

}

func generateImprovedStockArray(stockDataArray *[]StockData) *[]StockDataImproved {
	var StockDataImprovedArray []StockDataImproved
	var movingAverage50Sum float64 = 0
	var movingAverage200Sum float64 = 0
	var tempMovingAverage50 float64 = 0
	var tempMovingAverage200 float64 = 0

	// Calculate 50 and 200 day moving averages
	for i := 0; i < 200; i++ {
		movingAverage200Sum = movingAverage200Sum + (*stockDataArray)[i].Close
		// Only get the first 50 to start the array
		if i < 50 {
			movingAverage50Sum = movingAverage50Sum + (*stockDataArray)[i].Close
		}
	}

	// forLoop over dailyStock Prices
	for i, current := range *stockDataArray {
		// Calculate the Moving averages
		tempMovingAverage50 = movingAverage50Sum / 50
		tempMovingAverage200 = movingAverage200Sum / 200
		YearOnYearIterator := YearOverYearArrayPositionCalculator(stockDataArray, i)
		if YearOnYearIterator == -1 {
			break
		}

		// Calculate Year over year increase Rate
		YearOverYearOperatingCashFlowPerShare := current.OperatingCashFlowPerShare / (*stockDataArray)[YearOnYearIterator].OperatingCashFlowPerShare
		YearOverYearFreeCashFlowPerShare := current.FreeCashFlowPerShare / (*stockDataArray)[YearOnYearIterator].FreeCashFlowPerShare
		YearOverYearRevenuePerShare := current.RevenuePerShare / (*stockDataArray)[YearOnYearIterator].RevenuePerShare
		YearOverYearBookValuePerShare := current.BookValuePerShare / (*stockDataArray)[YearOnYearIterator].BookValuePerShare
		YearOverYearEvToOperatingCashFlow := current.EvToOperatingCashFlow / (*stockDataArray)[YearOnYearIterator].EvToOperatingCashFlow
		YearOverYearNetDebtToEBITDA := current.NetDebtToEBITDA / (*stockDataArray)[YearOnYearIterator].NetDebtToEBITDA
		YearOverYearGrahamNumber := current.GrahamNumber / (*stockDataArray)[YearOnYearIterator].GrahamNumber
		YearOverYearRoic := current.Roic / (*stockDataArray)[YearOnYearIterator].Roic

		currentStock := StockDataImproved{
			Symbol:                    current.Symbol,
			Datetime:                  current.Datetime,
			Close:                     current.Close,
			High:                      current.High,
			Low:                       current.Low,
			Open:                      current.Open,
			Volume:                    current.Volume,
			Period:                    current.Period,
			OperatingCashFlowPerShare: current.OperatingCashFlowPerShare,
			YearOverYearRateOperatingCashFlowPerShare: YearOverYearOperatingCashFlowPerShare,
			FreeCashFlowPerShare:                      current.FreeCashFlowPerShare,
			YearOverYearRateFreeCashFlowPerShare:      YearOverYearFreeCashFlowPerShare,
			CashPerShare:                              current.CashPerShare,
			PriceToSalesRatio:                         current.PriceToSalesRatio,
			PayoutRatio:                               current.PayoutRatio,
			RevenuePerShare:                           current.RevenuePerShare,
			YearOverYearRateRevenuePerShare:           YearOverYearRevenuePerShare,
			BookValuePerShare:                         current.BookValuePerShare,
			YearOverYearRateBookValuePerShare:         YearOverYearBookValuePerShare,
			MarketCap:                                 current.MarketCap,
			PeRatio:                                   current.PeRatio,
			PfcfRatio:                                 current.PfcfRatio,
			EvToOperatingCashFlow:                     current.EvToOperatingCashFlow,
			YearOverYearRateEvToOperatingCashFlow:     YearOverYearEvToOperatingCashFlow,
			NetDebtToEBITDA:                           current.NetDebtToEBITDA,
			YearOverYearRateNetDebtToEBITDA:           YearOverYearNetDebtToEBITDA,
			StockBasedCompensationToRevenue:           current.StockBasedCompensationToRevenue,
			GrahamNumber:                              current.GrahamNumber,
			YearOverYearRateGrahamNumber:              YearOverYearGrahamNumber,
			Roic:                                      current.Roic,
			YearOverYearRateRoic:                      YearOverYearRoic,
			Roe:                                       current.Roe,
			CapexPerShare:                             current.CapexPerShare,
			MovingAverage50Days:                       tempMovingAverage50,
			MovingAverage200Days:                      tempMovingAverage200,
		}

		// Append Array
		StockDataImprovedArray = append(StockDataImprovedArray, currentStock)

		// Calculate new average
		movingAverage50Sum = movingAverage50Sum - current.Close + (*stockDataArray)[i+50].Close
		movingAverage200Sum = movingAverage200Sum - current.Close + (*stockDataArray)[i+200].Close

	}

	return &StockDataImprovedArray

}
