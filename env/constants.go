package env

import (
	"fmt"
	"os"
)

// When trying to loop over top 25 stocks
// var StockList = [...]string{"AAPL", "MSFT", "AMZN", "NVDA", "GOOGL", "META", "GOOG", "TSLA", "BRK-B", "UNH", "LLY",
//
//		"JPM", "XOM", "AVGO", "V", "JNJ", "PG", "MA", "HD", "ADBE", "COST", "MRK", "CVX", "ABBV",
//		"WMT",
//	}
//
// Missing: "AAL", "T", "WIX", "VZ", "UPS", "FDX", "KSS", "BA", "PFE", "HOOD"
var StockList = [...]string{"KSS", "BA", "PFE", "HOOD"}

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
