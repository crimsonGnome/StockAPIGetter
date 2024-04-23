package main

import (
	"encoding/json"
	api "stockpulling/main/API"
	env "stockpulling/main/env"
)

func generateSingleStockCSV(stockSymbol string) {
	// Call Historical Finical API
	resultHistoricalFinancial := api.GetHistoricDataFinancials(stockSymbol)

	// Store datain Array
	var historicStockFinancialsArray []HistoricStockFinancials
	json.Unmarshal(resultHistoricalFinancial, &historicStockFinancialsArray)

	// Call Minute to Minute Stock data
	resultDailyPrice := api.GetDailyStockData(stockSymbol)

	//Store in Result
	var dailyStockPriceArray BodyDailyStockData
	json.Unmarshal(resultDailyPrice, &dailyStockPriceArray)

	// store API into Stock Data Struct
	stockDataArray := apiToStockDataStruct(&historicStockFinancialsArray, &dailyStockPriceArray, stockSymbol)

	// Improve Stock Data
	improvedStockDataArray := generateImprovedStockArray2(stockDataArray)

	// converts file into stock
	CSVconverter2(stockSymbol, improvedStockDataArray)

	// Orginal
	// improvedStockDataArray := generateImprovedStockArray(stockDataArray)
	// CSVconverter(stockSymbol, improvedStockDataArray)
}

func main() {
	// top25 generates 25 different csv files of the the top25a stocks in the S&P 500
	// top25()

	// Generates single csv file for single stock
	generateSingleStockCSV(env.ENV_STOCK_SYMBOL)

}
