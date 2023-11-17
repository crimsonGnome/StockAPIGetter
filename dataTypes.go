package main

// Struct for yahoofina
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

type HistoricStockFinancials struct {
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

type StockData struct {
	Symbol                                 string
	Date                                   string
	Close                                  string
	Datetime                               string
	High                                   int
	Low                                    float32
	Open                                   float32
	Volume                                 int
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
