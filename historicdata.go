package main

import (
	"fmt"
	"io"
	"net/http"
	"stockpulling/main/env"
)

func historicData(stockID string) {

	url := fmt.Sprintf("https://twelve-data1.p.rapidapi.com/time_series?interval=1month&symbol=%s&format=json&outputsize=30", stockID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
