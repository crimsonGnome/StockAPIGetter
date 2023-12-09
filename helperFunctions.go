package main

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"time"
	"os"
)

func top25(constants []string) {
	for _, stockSymbol := range StockList {
		generateCSV(stockSymbol)
	}
}

fucntion apiToStockDataStruct(historicStockFinancialsArray *[]HistoricStockFinancials, dailyStockPriceArray *BodyDailyStockData) *[]StockData {
	var stockDataArray []StockData

	// Iterator used to track which quarter the date respond to 
	quartleyReportCounter := 0
	// Date in which quarterly report was added
	dateStringQuartley := historicStockFinancialsArray[quartleyReportCounter].Date
	// format date
	dateQuartley, err := time.Parse("2006-01-02", dateStringQuartley)
	if err != nil {
		fmt.Println(err)
	}

	// forLoop over dailyStock Prices
	for _, dailyStockPrice := range dailyStockPriceArray.Values {

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
		dateTimeDaily, err := time.Parse("2006-01-02 15:04:05", dateStringDaily)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Sprintln(dateTimeDaily)

		// Compare if the quartely report is active
		for dateTimeDaily.Unix() < dateQuartley.Unix() {
			quartleyReportCounter = quartleyReportCounter + 1
			dateStringQuartley = historicStockFinancialsArray[quartleyReportCounter].Date
			dateQuartley, err = time.Parse("2006-01-02", dateStringQuartley)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Copy data metrics into currentStock 
		currentStock.Period = historicStockFinancialsArray[quartleyReportCounter].Period
		currentStock.OperatingCashFlowPerShare = historicStockFinancialsArray[quartleyReportCounter].OperatingCashFlowPerShare
		currentStock.FreeCashFlowPerShare = historicStockFinancialsArray[quartleyReportCounter].FreeCashFlowPerShare
		currentStock.CashPerShare = historicStockFinancialsArray[quartleyReportCounter].CashPerShare
		currentStock.DividendYield = historicStockFinancialsArray[quartleyReportCounter].DividendYield
		currentStock.PayoutRatio = historicStockFinancialsArray[quartleyReportCounter].PayoutRatio
		currentStock.RevenuePerShare = historicStockFinancialsArray[quartleyReportCounter].RevenuePerShare
		currentStock.NetIncomePerShare = historicStockFinancialsArray[quartleyReportCounter].NetIncomePerShare
		currentStock.BookValuePerShare = historicStockFinancialsArray[quartleyReportCounter].BookValuePerShare
		currentStock.ShareholdersEquityPerShare = historicStockFinancialsArray[quartleyReportCounter].ShareholdersEquityPerShare
		currentStock.InterestDebtPerShare = historicStockFinancialsArray[quartleyReportCounter].InterestDebtPerShare
		currentStock.MarketCap = historicStockFinancialsArray[quartleyReportCounter].MarketCap
		currentStock.EnterpriseValue = historicStockFinancialsArray[quartleyReportCounter].EnterpriseValue
		currentStock.PeRatio = historicStockFinancialsArray[quartleyReportCounter].PeRatio
		currentStock.Pocfratio = historicStockFinancialsArray[quartleyReportCounter].Pocfratio
		currentStock.PfcfRatio = historicStockFinancialsArray[quartleyReportCounter].PfcfRatio
		currentStock.Pbratio = historicStockFinancialsArray[quartleyReportCounter].Pbratio
		currentStock.PtbRatio = historicStockFinancialsArray[quartleyReportCounter].PtbRatio
		currentStock.EvToSales = historicStockFinancialsArray[quartleyReportCounter].EvToSales
		currentStock.EnterpriseValueOverEBITDA = historicStockFinancialsArray[quartleyReportCounter].EnterpriseValueOverEBITDA
		currentStock.EvToOperatingCashFlow = historicStockFinancialsArray[quartleyReportCounter].EvToOperatingCashFlow
		currentStock.EarningsYield = historicStockFinancialsArray[quartleyReportCounter].EarningsYield
		currentStock.FreeCashFlowYield = historicStockFinancialsArray[quartleyReportCounter].FreeCashFlowYield
		currentStock.DebtToEquity = historicStockFinancialsArray[quartleyReportCounter].DebtToEquity
		currentStock.DebtToAssets = historicStockFinancialsArray[quartleyReportCounter].DebtToAssets
		currentStock.NetDebtToEBITDA = historicStockFinancialsArray[quartleyReportCounter].NetDebtToEBITDA
		currentStock.CurrentRatio = historicStockFinancialsArray[quartleyReportCounter].CurrentRatio
		currentStock.InterestCoverage = historicStockFinancialsArray[quartleyReportCounter].InterestCoverage
		currentStock.IncomeQuality = historicStockFinancialsArray[quartleyReportCounter].IncomeQuality
		currentStock.SalesGeneralAndAdministrativeToRevenue = historicStockFinancialsArray[quartleyReportCounter].SalesGeneralAndAdministrativeToRevenue
		currentStock.ResearchAndDevelopmentToRevenue = historicStockFinancialsArray[quartleyReportCounter].ResearchAndDevelopmentToRevenue
		currentStock.IntangiblesToTotalAssets = historicStockFinancialsArray[quartleyReportCounter].IntangiblesToTotalAssets
		currentStock.CapexToOperatingCashFlow = historicStockFinancialsArray[quartleyReportCounter].CapexToOperatingCashFlow
		currentStock.CapexToRevenue = historicStockFinancialsArray[quartleyReportCounter].CapexToRevenue
		currentStock.CapexToDepreciation = historicStockFinancialsArray[quartleyReportCounter].CapexToDepreciation
		currentStock.StockBasedCompensationToRevenue = historicStockFinancialsArray[quartleyReportCounter].StockBasedCompensationToRevenue
		currentStock.GrahamNumber = historicStockFinancialsArray[quartleyReportCounter].GrahamNumber
		currentStock.Roic = historicStockFinancialsArray[quartleyReportCounter].Roic
		currentStock.ReturnOnTangibleAssets = historicStockFinancialsArray[quartleyReportCounter].ReturnOnTangibleAssets
		currentStock.GrahamNetNet = historicStockFinancialsArray[quartleyReportCounter].GrahamNetNet
		currentStock.WorkingCapital = historicStockFinancialsArray[quartleyReportCounter].WorkingCapital
		currentStock.TangibleAssetValue = historicStockFinancialsArray[quartleyReportCounter].TangibleAssetValue
		currentStock.NetCurrentAssetValue = historicStockFinancialsArray[quartleyReportCounter].NetCurrentAssetValue
		currentStock.InvestedCapital = historicStockFinancialsArray[quartleyReportCounter].InvestedCapital
		currentStock.AverageReceivables = historicStockFinancialsArray[quartleyReportCounter].AverageReceivables
		currentStock.AveragePayables = historicStockFinancialsArray[quartleyReportCounter].AveragePayables
		currentStock.AverageInventory = historicStockFinancialsArray[quartleyReportCounter].AverageInventory
		currentStock.DaysSalesOutstanding = historicStockFinancialsArray[quartleyReportCounter].DaysSalesOutstanding
		currentStock.DaysPayablesOutstanding = historicStockFinancialsArray[quartleyReportCounter].DaysPayablesOutstanding
		currentStock.DaysOfInventoryOnHand = historicStockFinancialsArray[quartleyReportCounter].DaysOfInventoryOnHand
		currentStock.ReceivablesTurnover = historicStockFinancialsArray[quartleyReportCounter].ReceivablesTurnover
		currentStock.PayablesTurnover = historicStockFinancialsArray[quartleyReportCounter].PayablesTurnover
		currentStock.InventoryTurnover = historicStockFinancialsArray[quartleyReportCounter].InventoryTurnover
		currentStock.Roe = historicStockFinancialsArray[quartleyReportCounter].Roe
		currentStock.CapexPerShare = historicStockFinancialsArray[quartleyReportCounter].CapexPerShare

		stockDataArray = append(stockDataArray, currentStock)
	}

}

func CSVconverter(destinationFileName string, stockDataArray *[]StockData) error {
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
		"FreeCashFlowPerShare",
		"CashPerShare",
		"PriceToSalesRatio",
		"DividendYield",
		"PayoutRatio",
		"RevenuePerShare",
		"NetIncomePerShare",
		"BookValuePerShare",
		"TangibleBookValuePerShare",
		"ShareholdersEquityPerShare",
		"InterestDebtPerShare",
		"MarketCap",
		"EnterpriseValue",
		"PeRatio",
		"Pocfratio",
		"PfcfRatio",
		"Pbratio",
		"PtbRatio",
		"EvToSales",
		"EnterpriseValueOverEBITDA",
		"EvToOperatingCashFlow",
		"EarningsYield",
		"FreeCashFlowYield",
		"DebtToEquity",
		"DebtToAssets",
		"NetDebtToEBITDA",
		"CurrentRatio",
		"InterestCoverage",
		"IncomeQuality",
		"SalesGeneralAndAdministrativeToRevenue",
		"ResearchAndDevelopmentToRevenue",
		"IntangiblesToTotalAssets",
		"CapexToOperatingCashFlow",
		"CapexToRevenue",
		"CapexToDepreciation",
		"StockBasedCompensationToRevenue",
		"GrahamNumber",
		"Roic",
		"ReturnOnTangibleAssets",
		"GrahamNetNet",
		"WorkingCapital",
		"TangibleAssetValue",
		"NetCurrentAssetValue",
		"InvestedCapital",
		"AverageReceivables",
		"AveragePayables",
		"AverageInventory",
		"DaysSalesOutstanding",
		"DaysPayablesOutstanding",
		"DaysOfInventoryOnHand",
		"ReceivablesTurnover",
		"PayablesTurnover",
		"InventoryTurnover",
		"Roe",
		"CapexPerShare"}
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
			fmt.Sprintf("%f", r.FreeCashFlowPerShare),
			fmt.Sprintf("%f", r.CashPerShare),
			fmt.Sprintf("%f", r.PriceToSalesRatio),
			fmt.Sprintf("%f", r.DividendYield),
			fmt.Sprintf("%f", r.PayoutRatio),
			fmt.Sprintf("%f", r.RevenuePerShare),
			fmt.Sprintf("%f", r.NetIncomePerShare),
			fmt.Sprintf("%f", r.BookValuePerShare),
			fmt.Sprintf("%f", r.TangibleBookValuePerShare),
			fmt.Sprintf("%f", r.ShareholdersEquityPerShare),
			fmt.Sprintf("%f", r.InterestDebtPerShare),
			fmt.Sprintf("%f", r.MarketCap),
			fmt.Sprintf("%f", r.EnterpriseValue),
			fmt.Sprintf("%f", r.PeRatio),
			fmt.Sprintf("%f", r.Pocfratio),
			fmt.Sprintf("%f", r.PfcfRatio),
			fmt.Sprintf("%f", r.Pbratio),
			fmt.Sprintf("%f", r.PtbRatio),
			fmt.Sprintf("%f", r.EvToSales),
			fmt.Sprintf("%f", r.EnterpriseValueOverEBITDA),
			fmt.Sprintf("%f", r.EvToOperatingCashFlow),
			fmt.Sprintf("%f", r.EarningsYield),
			fmt.Sprintf("%f", r.FreeCashFlowYield),
			fmt.Sprintf("%f", r.DebtToEquity),
			fmt.Sprintf("%f", r.DebtToAssets),
			fmt.Sprintf("%f", r.NetDebtToEBITDA),
			fmt.Sprintf("%f", r.CurrentRatio),
			fmt.Sprintf("%f", r.InterestCoverage),
			fmt.Sprintf("%f", r.IncomeQuality),
			fmt.Sprintf("%f", r.SalesGeneralAndAdministrativeToRevenue),
			fmt.Sprintf("%f", r.ResearchAndDevelopmentToRevenue),
			fmt.Sprintf("%f", r.IntangiblesToTotalAssets),
			fmt.Sprintf("%f", r.CapexToOperatingCashFlow),
			fmt.Sprintf("%f", r.CapexToRevenue),
			fmt.Sprintf("%f", r.CapexToDepreciation),
			fmt.Sprintf("%f", r.StockBasedCompensationToRevenue),
			fmt.Sprintf("%f", r.GrahamNumber),
			fmt.Sprintf("%f", r.Roic),
			fmt.Sprintf("%f", r.ReturnOnTangibleAssets),
			fmt.Sprintf("%f", r.GrahamNetNet),
			fmt.Sprintf("%f", r.WorkingCapital),
			fmt.Sprintf("%f", r.TangibleAssetValue),
			fmt.Sprintf("%f", r.NetCurrentAssetValue),
			fmt.Sprintf("%f", r.InvestedCapital),
			fmt.Sprintf("%f", r.AverageReceivables),
			fmt.Sprintf("%f", r.AveragePayables),
			fmt.Sprintf("%f", r.AverageInventory),
			fmt.Sprintf("%f", r.DaysSalesOutstanding),
			fmt.Sprintf("%f", r.DaysPayablesOutstanding),
			fmt.Sprintf("%f", r.DaysOfInventoryOnHand),
			fmt.Sprintf("%f", r.ReceivablesTurnover),
			fmt.Sprintf("%f", r.PayablesTurnover),
			fmt.Sprintf("%f", r.InventoryTurnover),
			fmt.Sprintf("%f", r.Roe),
			fmt.Sprintf("%f", r.CapexPerShare),
		)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}
