package main

import (
	"encoding/json"
	"fmt"
	api "stockpulling/main/API"
)

func generateCSV(stockSymbol string) {
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


	stockDataArray := apiToStockDataStruct(&historicStockFinancialsArray, &dailyStockPriceArray )

	fileName := fmt.Sprintf("_%s.csv", stockSymbol)

	// converts file into stock 
	CSVconverter(fileName, &stockDataArray)
}

func main() {
	// Call both API
	// for _, stockSymbol := range StockList {
	stockSymbol := "AAPL"
	generateCSV(stockSymbol)

}
