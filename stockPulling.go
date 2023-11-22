package main

import (
	"encoding/json"
	"fmt"
	api "stockpulling/main/API"
	"strconv"
	"time"
)

func main() {
	// Call both API
	// for _, stockSymbol := range StockList {
	stockSymbol := "XOM"
	resultHistorical := api.GetHistoricDataFinancials(stockSymbol)

	var historicStockFinancialsArray []HistoricStockFinancials
	json.Unmarshal(resultHistorical, &historicStockFinancialsArray)

	fmt.Printf("%+v", historicStockFinancialsArray)

	resultDaily := api.GetDailyStockData(stockSymbol)

	var dailyStockPriceArray BodyDailyStockData
	json.Unmarshal(resultDaily, &dailyStockPriceArray)

	fmt.Printf("%+v", dailyStockPriceArray)

	// TODO:
	// create new master data structure
	var stockDataArray []StockData
	quartleyReportCounter := 0
	dateStringQuartley := historicStockFinancialsArray[quartleyReportCounter].Date
	dateQuartley, err := time.Parse("2006-01-02", dateStringQuartley)
	if err != nil {
		fmt.Println(err)
	}

	// forLoop over dailyStock Price
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

		dateDaily, err := time.Parse("2006-01-02", dateStringDaily)
		if err != nil {
			fmt.Println(err)
		}

		if dateDaily.Unix() < dateQuartley.Unix() {
			quartleyReportCounter = quartleyReportCounter + 1
			dateStringQuartley = historicStockFinancialsArray[quartleyReportCounter].Date
			dateQuartley, err = time.Parse("2006-01-02", dateStringQuartley)
			if err != nil {
				fmt.Println(err)
			}
		}

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

	// Saving as JSON format

	// file, _ := json.MarshalIndent(stockDataArray, "", " ")
	fileName := fmt.Sprintf("_%s.csv", stockSymbol)
	// _ = ioutil.WriteFile(fileName, file, 0644)

	CSVconverter(fileName, &stockDataArray)
}
