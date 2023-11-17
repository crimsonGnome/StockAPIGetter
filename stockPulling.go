package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	api "stockpulling/main/API"
	"time"
)

func main() {
	// Call both API
	resultHistorical := api.GetHistoricDataFinancials("AMZN")

	var historicStockFinancialsArray []HistoricStockFinancials
	json.Unmarshal(resultHistorical, &historicStockFinancialsArray)

	fmt.Printf("%+v", historicStockFinancialsArray)

	resultDaily := api.GetDailyStockData("AMZN")

	var dailyStockPriceArray []DailyStockData
	json.Unmarshal(resultDaily, &dailyStockPriceArray)

	fmt.Printf("%+v", dailyStockPriceArray)

	// TODO:
	// create new master data structure
	var stockDataArray []StockData
	quartleyReportCounter := 0
	// forLoop over dailyStock Price
	for i := 0; i < len(dailyStockPriceArray); i++ {
		// - add Daily stock values to meta data
		dailyStockPrice := DailyStockData{
			Symbol:   i.Meta.Symbol,
			Date:     i.Body.Date,
			Close:    i.Body.Close,
			Datetime: i.Body.Datetime,
			High:     i.Body.High,
			Low:      i.Body.Low,
			Open:     i.Body.Open,
			Volume:   i.Body.Volume,
		}
		// check if dailyStock date is >= Historic Financial Date
		dateStringDaily := dailyStockPrice.date
		dateStringQuartley := historicStockFinancialsArray[quartleyReportCounter].Date

		dateDaily, error := time.Parse("02/28/2016 9:03:46 PM", dateStringDaily)
		if err != nil {
			fmt.Println(err)
		}

		dateQuartley, error := time.Parse("02/28/2016 9:03:46 PM", dateStringQuartley)
		if err != nil {
			fmt.Println(err)
		}

		if dateDaily.Unix() < dateQuartley {
			quartleyReportCounter++
		}

		dailyStockPrice.Period = historicStockFinancialsArray[quartleyReportCounter].Period
		dailyStockPrice.OperatingCashFlowPerShare = historicStockFinancialsArray[quartleyReportCounter].operatingCashFlowPerShare
		dailyStockPrice.FreeCashFlowPerShare = historicStockFinancialsArray[quartleyReportCounter].FreeCashFlowPerShare
		dailyStockPrice.CashPerShare = historicStockFinancialsArray[quartleyReportCounter].CashPerShare
		dailyStockPrice.DividendYield = historicStockFinancialsArray[quartleyReportCounter].DividendYield
		dailyStockPrice.PayoutRatio = historicStockFinancialsArray[quartleyReportCounter].PayoutRatio
		dailyStockPrice.RevenuePerShare = historicStockFinancialsArray[quartleyReportCounter].RevenuePerShare
		dailyStockPrice.NetIncomePerShare = historicStockFinancialsArray[quartleyReportCounter].NetIncomePerShare
		dailyStockPrice.BookValuePerShare = historicStockFinancialsArray[quartleyReportCounter].BookValuePerShare
		dailyStockPrice.ShareholdersEquityPerShare = historicStockFinancialsArray[quartleyReportCounter].ShareholdersEquityPerShare
		dailyStockPrice.InterestDebtPerShare = historicStockFinancialsArray[quartleyReportCounter].interestDebtPerShare
		dailyStockPrice.MarketCap = historicStockFinancialsArray[quartleyReportCounter].MarketCap
		dailyStockPrice.EnterpriseValue = historicStockFinancialsArray[quartleyReportCounter].EnterpriseValue
		dailyStockPrice.PeRatio = historicStockFinancialsArray[quartleyReportCounter].PeRatio
		dailyStockPrice.Pocfratio = historicStockFinancialsArray[quartleyReportCounter].Pocfratio
		dailyStockPrice.PfcfRatio = historicStockFinancialsArray[quartleyReportCounter].PfcfRatio
		dailyStockPrice.Pbratio = historicStockFinancialsArray[quartleyReportCounter].Pbratio
		dailyStockPrice.PtbRatio = historicStockFinancialsArray[quartleyReportCounter].PtbRatio
		dailyStockPrice.EvToSales = historicStockFinancialsArray[quartleyReportCounter].EvToSales
		dailyStockPrice.EnterpriseValueOverEBITDA = historicStockFinancialsArray[quartleyReportCounter].EnterpriseValueOverEBITDA
		dailyStockPrice.EvToOperatingCashFlow = historicStockFinancialsArray[quartleyReportCounter].EvToOperatingCashFlow
		dailyStockPrice.EarningsYield = historicStockFinancialsArray[quartleyReportCounter].EarningsYield
		dailyStockPrice.FreeCashFlowYield = historicStockFinancialsArray[quartleyReportCounter].FreeCashFlowYield
		dailyStockPrice.DebtToEquity = historicStockFinancialsArray[quartleyReportCounter].DebtToEquity
		dailyStockPrice.DebtToAssets = historicStockFinancialsArray[quartleyReportCounter].DebtToAssets
		dailyStockPrice.NetDebtToEBITDA = historicStockFinancialsArray[quartleyReportCounter].NetDebtToEBITDA
		dailyStockPrice.CurrentRatio = historicStockFinancialsArray[quartleyReportCounter].currentRatio
		dailyStockPrice.InterestCoverage = historicStockFinancialsArray[quartleyReportCounter].InterestCoverage
		dailyStockPrice.IncomeQuality = historicStockFinancialsArray[quartleyReportCounter].IncomeQuality
		dailyStockPrice.SalesGeneralAndAdministrativeToRevenue = historicStockFinancialsArray[quartleyReportCounter].SalesGeneralAndAdministrativeToRevenue
		dailyStockPrice.ResearchAndDevelopmentToRevenue = historicStockFinancialsArray[quartleyReportCounter].ResearchAndDevelopmentToRevenue
		dailyStockPrice.IntangiblesToTotalAssets = historicStockFinancialsArray[quartleyReportCounter].IntangiblesToTotalAssets
		dailyStockPrice.CapexToOperatingCashFlow = historicStockFinancialsArray[quartleyReportCounter].CapexToOperatingCashFlow
		dailyStockPrice.CapexToRevenue = historicStockFinancialsArray[quartleyReportCounter].CapexToRevenue
		dailyStockPrice.CapexToDepreciation = historicStockFinancialsArray[quartleyReportCounter].CapexToDepreciation
		dailyStockPrice.StockBasedCompensationToRevenue = historicStockFinancialsArray[quartleyReportCounter].StockBasedCompensationToRevenue
		dailyStockPrice.GrahamNumber = historicStockFinancialsArray[quartleyReportCounter].GrahamNumber
		dailyStockPrice.Roic = historicStockFinancialsArray[quartleyReportCounter].Roic
		dailyStockPrice.ReturnOnTangibleAssets = historicStockFinancialsArray[quartleyReportCounter].ReturnOnTangibleAssets
		dailyStockPrice.GrahamNetNet = historicStockFinancialsArray[quartleyReportCounter].GrahamNetNet
		dailyStockPrice.WorkingCapital = historicStockFinancialsArray[quartleyReportCounter].WorkingCapital
		dailyStockPrice.TangibleAssetValue = historicStockFinancialsArray[quartleyReportCounter].TangibleAssetValue
		dailyStockPrice.NetCurrentAssetValue = historicStockFinancialsArray[quartleyReportCounter].NetCurrentAssetValue
		dailyStockPrice.InvestedCapital = historicStockFinancialsArray[quartleyReportCounter].InvestedCapital
		dailyStockPrice.AverageReceivables = historicStockFinancialsArray[quartleyReportCounter].AverageReceivables
		dailyStockPrice.AveragePayables = historicStockFinancialsArray[quartleyReportCounter].AveragePayables
		dailyStockPrice.AverageInventory = historicStockFinancialsArray[quartleyReportCounter].AverageInventory
		dailyStockPrice.DaysSalesOutstanding = historicStockFinancialsArray[quartleyReportCounter].DaysSalesOutstanding
		dailyStockPrice.DaysPayablesOutstanding = historicStockFinancialsArray[quartleyReportCounter].DaysPayablesOutstanding
		dailyStockPrice.DaysOfInventoryOnHand = historicStockFinancialsArray[quartleyReportCounter].DaysOfInventoryOnHand
		dailyStockPrice.ReceivablesTurnover = historicStockFinancialsArray[quartleyReportCounter].ReceivablesTurnover
		dailyStockPrice.PayablesTurnover = historicStockFinancialsArray[quartleyReportCounter].PayablesTurnover
		dailyStockPrice.InventoryTurnover = historicStockFinancialsArray[quartleyReportCounter].InventoryTurnover
		dailyStockPrice.Roe = historicStockFinancialsArray[quartleyReportCounter].Roe
		dailyStockPrice.CapexPerShare = historicStockFinancialsArray[quartleyReportCounter].CapexPerShare

		stockDataArray = append(stockDataArray, dailyStockPrice)
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("AMZN.json", file, 0644)

}
