package api

import (
	"fmt"
	"io"
	"net/http"
	"stockpulling/main/env"
)

func GetDailyStockData(StockID string) []byte {

	url := fmt.Sprintf("https://twelve-data1.p.rapidapi.com/time_series?interval=1day&symbol=%s&format=json&outputsize=1500", StockID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body

}
