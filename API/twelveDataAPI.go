package api

import (
	"fmt"
	"io"
	"net/http"
	env "stockpulling/main/env"
)

func GetDailyStockData(StockID string) []byte {

	// Format for time series - local run
	// timeSeries := end_date=2020-03-24 10:07:00
	// url := fmt.Sprintf("https://twelve-data1.p.rapidapi.com/time_series?%s&interval=1min&symbol=%s&format=json&outputsize=%s", timeSeries, StockID, os.Getenv("ENV_OUTPUT_SIZE"))

	// Lambda function call - want the most up to date data.
	url := fmt.Sprintf("https://twelve-data1.p.rapidapi.com/time_series?interval=1min&symbol=%s&format=json&outputsize=%s", StockID, env.ENV_OUTPUT_SIZE)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body

}
