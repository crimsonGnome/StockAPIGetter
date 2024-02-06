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
	dateQuartley, err := time.Parse("2006-01-02", dateStringQuarterly)
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
		dateTimeDaily, err := time.Parse("2006-01-02 15:04:05", dateStringDaily)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Sprintln(dateTimeDaily)

		// Compare if the quartely report is active
		for dateTimeDaily.Unix() < dateQuartley.Unix() {
			quarterlyReportCounter = quarterlyReportCounter + 1
			dateStringQuarterly = (*historicStockFinancialsArray)[quarterlyReportCounter].Date
			dateQuartley, err = time.Parse("2006-01-02", dateStringQuarterly)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Copy data metrics into currentStock
		currentStock.Period = (*historicStockFinancialsArray)[quarterlyReportCounter].Period
		currentStock.OperatingCashFlowPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].OperatingCashFlowPerShare
		currentStock.FreeCashFlowPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].FreeCashFlowPerShare
		currentStock.CashPerShare = (*historicStockFinancialsArray)[quarterlyReportCounter].CashPerShare
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

func CSVconverter(stockSymbol string, stockDataArray *[]StockData) error {
	destinationFileName := fmt.Sprintf("_%s.csv", stockSymbol)

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
