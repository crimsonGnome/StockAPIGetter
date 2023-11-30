package api

import (
	"fmt"
	"io"
	"net/http"
	"stockpulling/main/env"
)

func GetDailyStockData(StockID string) []byte {

	url := fmt.Sprintf("https://twelve-data1.p.rapidapi.com/time_series?end_date=2020-03-24 10:07:00&interval=1min&symbol=%s&format=json&outputsize=5000", StockID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body

}
