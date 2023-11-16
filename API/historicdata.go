package api

import (
	"fmt"
	"io"
	"net/http"
	"stockpulling/main/env"
)

type GetHistoricDataFinancials struct {
	Symbol                                 string
	Date                                   string
	Period                                 string
	OperatingCashFlowPerShare              float32
	FreeCashFlowPerShare                   float32
	CashPerShare                           float32
	PriceToSalesRatio                      float32
	DividendYield                          float32
	PayoutRatio                            float32
	RevenuePerShare                        float32
	NetIncomePerShare                      float32
	BookValuePerShare                      float32
	TangibleBookValuePerShare              float32
	ShareholdersEquityPerShare             float32
	InterestDebtPerShare                   float32
	MarketCap                              float32
	EnterpriseValue                        float32
	PeRatio                                float32
	Pocfratio                              float32
	PfcfRatio                              float32
	Pbratio                                float32
	PtbRatio                               float32
	EvToSales                              float32
	EnterpriseValueOverEBITDA              float32
	EvToOperatingCashFlow                  float32
	EarningsYield                          float32
	FreeCashFlowYield                      float32
	DebtToEquity                           float32
	DebtToAssets                           float32
	NetDebtToEBITDA                        float32
	CurrentRatio                           float32
	InterestCoverage                       float32
	IncomeQuality                          float32
	SalesGeneralAndAdministrativeToRevenue float32
	ResearchAndDevelopmentToRevenue        float32
	IntangiblesToTotalAssets               float32
	CapexToOperatingCashFlow               float32
	CapexToRevenue                         float32
	CapexToDepreciation                    float32
	StockBasedCompensationToRevenue        float32
	GrahamNumber                           float32
	Roic                                   float32
	ReturnOnTangibleAssets                 float32
	GrahamNetNet                           float32
	WorkingCapital                         float32
	TangibleAssetValue                     float32
	NetCurrentAssetValue                   float32
	InvestedCapital                        float32
	AverageReceivables                     float32
	AveragePayables                        float32
	AverageInventory                       float32
	DaysSalesOutstanding                   float32
	DaysPayablesOutstanding                float32
	DaysOfInventoryOnHand                  float32
	ReceivablesTurnover                    float32
	PayablesTurnover                       float32
	InventoryTurnover                      float32
	Roe                                    float32
	CapexPerShare                          float32
}

func GetHistoricData(StockID string) []byte {
	// data := []HistoricStockFinancials{}
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

	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(data[i].symbol)
	// 	fmt.Println(data[i].date)
	// }
	// for i := 0; i < size(iot.body); i++ {
	// 	var singleData HistoricStockFinancials

	// 	singleData.symbol = iot.body[i].symbol
	// 	singleData.date = iot.body[i].date
	// 	singleData.operatingCashFlowPerShare = iot.body[i].operatingCashFlowPerShare
	// 	singleData.freeCashFlowPerShare = iot.body[i].freeCashFlowPerShare
	// 	singleData.cashPerShare = iot.body[i].cashPerShare
	// 	singleData.dividendYield = iot.body[i].dividendYield
	// 	singleData.payoutRatio = iot.body[i].payoutRatio
	// 	singleData.revenuePerShare = iot.body[i].revenuePerShare
	// 	singleData.netIncomePerShare = iot.body[i].netIncomePerShare
	// 	singleData.bookValuePerShare = iot.body[i].bookValuePerShare
	// 	singleData.shareholdersEquityPerShare = iot.body[i].shvareholdersEquityPerShare
	// 	singleData.interestDebtPerShare = iot.body[i].interestDebtPerShare
	// 	singleData.marketCap = iot.body[i].marketCap
	// 	singleData.enterpriseValue = iot.body[i].enterpriseValue
	// 	singleData.peRatio = iot.body[i].peRatio
	// 	singleData.pocfratio = iot.body[i].pocfratio
	// 	singleData.pfcfRatio = iot.body[i].pfcfRatio
	// 	singleData.pbratio = iot.body[i].pbratio
	// 	singleData.ptbRatio = iot.body[i].ptbRatio
	// 	singleData.evToSales = iot.body[i].evToSales
	// 	singleData.enterpriseValueOverEBITDA = iot.body[i].enterpriseValueOverEBITDA
	// 	singleData.evToOperatingCashFlow = iot.body[i].evToOperatingCashFlow
	// 	singleData.earningsYield = iot.body[i].earningsYield
	// 	singleData.freeCashFlowYield = iot.body[i].freeCashFlowYield
	// 	singleData.debtToEquity = iot.body[i].debtToEquity
	// 	singleData.debtToAssets = iot.body[i].debtToAssets
	// 	singleData.netDebtToEBITDA = iot.body[i].netDebtToEBITDA
	// 	singleData.currentRatio = iot.body[i].currentRatio
	// 	singleData.interestCoverage = iot.body[i].interestCoverage
	// 	singleData.incomeQuality = iot.body[i].incomeQuality
	// 	singleData.salesGeneralAndAdministrativeToRevenue = iot.body[i].salesGeneralAndAdministrativeToRevenue
	// 	singleData.researchAndDevelopmentToRevenue = iot.body[i].researchAndDevelopmentToRevenue
	// 	singleData.intangiblesToTotalAssets = iot.body[i].intangiblesToTotalAssets
	// 	singleData.capexToOperatingCashFlow = iot.body[i].capexToOperatingCashFlow
	// 	singleData.capexToRevenue = iot.body[i].capexToRevenue
	// 	singleData.capexToDepreciation = iot.body[i].capexToDepreciation
	// 	singleData.stockBasedCompensationToRevenue = iot.body[i].stockBasedCompensationToRevenue
	// 	singleData.grahamNumber = iot.body[i].grahamNumber
	// 	singleData.roic = iot.body[i].roic
	// 	singleData.returnOnTangibleAssets = iot.body[i].returnOnTangibleAssets
	// 	singleData.grahamNetNet = iot.body[i].grahamNetNet
	// 	singleData.workingCapital = iot.body[i].workingCapital
	// 	singleData.tangibleAssetValue = iot.body[i].tangibleAssetValue
	// 	singleData.netCurrentAssetValue = iot.body[i].netCurrentAssetValue
	// 	singleData.investedCapital = iot.body[i].investedCapital
	// 	singleData.averageReceivables = iot.body[i].averageReceivables
	// 	singleData.averagePayables = iot.body[i].averagePayables
	// 	singleData.averageInventory = iot.body[i].averageInventory
	// 	singleData.daysSalesOutstanding = iot.body[i].daysSalesOutstanding
	// 	singleData.daysPayablesOutstanding = iot.body[i].daysPayablesOutstanding
	// 	singleData.daysOfInventoryOnHand = iot.body[i].daysOfInventoryOnHand
	// 	singleData.receivablesTurnover = iot.body[i].receivablesTurnover
	// 	singleData.payablesTurnover = iot.body[i].payablesTurnover
	// 	singleData.inventoryTurnover = iot.body[i].inventoryT.urnover
	// 	singleData.roe = iot.body[i].roe
	// 	singleData.capexPerShare = iot.body[i].capexPerShare

	// 	fmt.Println(singleData)
	// 	data = append(data, singleData)

	// }

	// fmt.Println(balance)
	// fmt.Println(string(body))
	// fmt.Printf("\nbreak\n")

}
