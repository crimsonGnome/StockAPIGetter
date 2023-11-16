package api

import (
	"fmt"
	"io"
	"net/http"
	"stockpulling/main/env"
)

type BodyDailyStockData struct {
	Meta MetaStockData
	Body []DailyStockData
}

type MetaStockData struct {
	Currency          string
	Exchange          string
	Exchange_timezone string
	Interval          string
	Symbol            string
	Type_             string
	Status            string
}

type DailyStockData struct {
	Close    string
	Datetime string
	High     int
	Low      float32
	Open     float32
	Volume   int
}

func GetDailyStockData(StockID string) []byte {

	url := "https://twelve-data1.p.rapidapi.com/time_series?interval=1day&symbol=AMZN&format=json&outputsize=1500"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return body

}
