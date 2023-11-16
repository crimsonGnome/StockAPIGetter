package main

import (
	"encoding/json"
	"fmt"
	api "stockpulling/main/API"
)

func main() {
	// result := api.GetHistoricDataFinancials"AMZN")

	// var financials []api.HistoricStockFinancials
	// json.Unmarshal(result, &financials)

	// fmt.Printf("%+v", financials)

	result := api.GetDailyStockData("AMZN")

	var price []api.DailyStockData
	json.Unmarshal(result, &price)

	fmt.Printf("%+v", price)

	// TODO:
	// Call both API
	// create new master data structure
	// forLoop over dailyStock Price
	// add Daily stock values to meta data
	// if UTC date time is lower then date time of quarley reprot go to the next (older) report
	// add all data in
	// Add new value to s3 bucket

}
