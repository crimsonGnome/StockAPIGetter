package api

import (
	"fmt"
	"io"
	"net/http"
	env "stockpulling/main/env"
)

func GetHistoricDataFinancials(StockID string) []byte {
	url := fmt.Sprintf("https://holistic-finance-stock-data.p.rapidapi.com/api/v1/keymetrics?symbol=%s&period=quarter", StockID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "holistic-finance-stock-data.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return responseData

}
