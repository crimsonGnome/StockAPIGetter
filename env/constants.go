package env

import (
	"fmt"
	"os"
)

// When trying to loop over top 25 stocks
// var StockList = [...]string{
//
//			"AAPL", "ABBV", "ABT", "ADBE", "AMZN", "AVGO", "AXP"
//			"CMG", "COST", "CRM", "CRWD", "FICO", "GE", "GOOG", "HLT"
//			"IDXX", "INTU", "JPM", "KKR", "LIN", "LLY"
//		    "MA", "MCO", "META", "MSFT", "MSTR", "NFLX", "NVDA"
//	        "QCOM", "SPGI", "SYK", "TMO", "UNH", "V", "WDAY", "WMT"
//
// Missing:
// Only can run
var StockList = [...]string{"QCOM", "SPGI", "SYK", "TMO", "UNH", "V", "WDAY", "WMT"}

func getEnvVariable(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		err := fmt.Errorf("environmentalVariable: %s does not exist", key)
		panic(err)
	}
	return value
}

var ENV_STOCK_SYMBOL = getEnvVariable("ENV_STOCK_SYMBOL")
var ENV_RAPID_API_KEY = getEnvVariable("ENV_RAPID_API_KEY")
var ENV_OUTPUT_SIZE = getEnvVariable("ENV_OUTPUT_SIZE")
