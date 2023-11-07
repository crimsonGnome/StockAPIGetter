package main

import (
	"fmt"
	"io"
	"net/http"
	"stockpulling/main/env"
)

func yahooFnPricing() {

	url := "https://yh-finance-complete.p.rapidapi.com/yhfhistorical?ticker=AMZN&sdate=11-2-2020&edate=11-2-2023"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "yh-finance-complete.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func yahooFn() {

	url := "https://yh-finance-complete.p.rapidapi.com/balanceSheetHistoryQuarterly?symbol=AMZN"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "yh-finance-complete.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
