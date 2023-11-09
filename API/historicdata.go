package api

import (
	"fmt"
	"io"
	"net/http"
)

type HistoricStockFinancials struct {
	StockTag string
	operatingCashFlowPerShare float32
	freeCashFlowPerShare float32
	cashPerShare float32
	cashPerShare float32
	dividendYield float32
	payoutRatio float32
	revenuePerShare float32
	netIncomePerShare float32
	bookValuePerShare:float32 
	shareholdersEquityPerShare float32
	interestDebtPerShare float32
	marketCap float32
	enterpriseValue float32
	peRatio float32
	pocfratio float32
	pfcfRatio float32
	pbratio float32
	ptbRatio float32
	evToSales float32
	enterpriseValueOverEBITDA float32
	evToOperatingCashFlow float32
	earningsYield float32
	freeCashFlowYield float32
	debtToEquity float32
	debtToAssets float32
	netDebtToEBITDA float32
	currentRatio float32
	interestCoverage float32
	incomeQuality float32
	salesGeneralAndAdministrativeToRevenue: float32
	researchAndDevelopmentToRevenue:float32
	intangiblesToTotalAssets float32
	capexToOperatingCashFlow float32
	capexToRevenue float32
	capexToDepreciation float32
	stockBasedCompensationToRevenue float32
	grahamNumber float32
	roic float32
	returnOnTangibleAssets float32
	grahamNetNet float32
	workingCapital float32
	tangibleAssetValue float32
	netCurrentAssetValue float32
	investedCapital float32
	averageReceivables float32
	averagePayables float32
	averageInventory float32
	daysSalesOutstanding float32
	daysPayablesOutstanding float32
	daysOfInventoryOnHand float32
	receivablesTurnover float32
	payablesTurnover float32
	inventoryTurnover float32
	roe float32
	capexPerShare float32

}

func GetHistoricData(StockID string) {

	url := fmt.Sprintf("https://holistic-finance-stock-data.p.rapidapi.com/api/v1/keymetrics?symbol=%s&period=quarterly", StockID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", env.ENV_RAPID_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "holistic-finance-stock-data.p.rapidapi.com")

	balance, _ := http.DefaultClient.Do(req)

	defer balance.Body.Close()
	body, _ := io.ReadAll(balance.Body)

	fmt.Println(balance)
	fmt.Println(string(body))

	url := fmt.Sprintf("https://holistic-finance-stock-data.p.rapidapi.com/api/v1/keymetrics?symbol=%s&period=quarterly", StockID)

}
