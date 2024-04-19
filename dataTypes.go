package main

// Struct for yahoofina
type BodyDailyStockData struct {
	Meta   MetaStockData
	Values []DailyStockData
}

type MetaStockData struct {
	Symbol            string
	Interval          string
	Currency          string
	Exchange_timezone string
	Exchange          string
	Mic_code          string
	Type              string
}

type DailyStockData struct {
	Datetime string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
}

type HistoricStockFinancials struct {
	Symbol                                 string
	Date                                   string
	Period                                 string
	OperatingCashFlowPerShare              float64
	FreeCashFlowPerShare                   float64
	CashPerShare                           float64
	PriceToSalesRatio                      float64
	DividendYield                          float64
	PayoutRatio                            float64
	RevenuePerShare                        float64
	NetIncomePerShare                      float64
	BookValuePerShare                      float64
	TangibleBookValuePerShare              float64
	ShareholdersEquityPerShare             float64
	InterestDebtPerShare                   float64
	MarketCap                              float64
	EnterpriseValue                        float64
	PeRatio                                float64
	Pocfratio                              float64
	PfcfRatio                              float64
	Pbratio                                float64
	PtbRatio                               float64
	EvToSales                              float64
	EnterpriseValueOverEBITDA              float64
	EvToOperatingCashFlow                  float64
	EarningsYield                          float64
	FreeCashFlowYield                      float64
	DebtToEquity                           float64
	DebtToAssets                           float64
	NetDebtToEBITDA                        float64
	CurrentRatio                           float64
	InterestCoverage                       float64
	IncomeQuality                          float64
	SalesGeneralAndAdministrativeToRevenue float64
	ResearchAndDevelopmentToRevenue        float64
	IntangiblesToTotalAssets               float64
	CapexToOperatingCashFlow               float64
	CapexToRevenue                         float64
	CapexToDepreciation                    float64
	StockBasedCompensationToRevenue        float64
	GrahamNumber                           float64
	Roic                                   float64
	ReturnOnTangibleAssets                 float64
	GrahamNetNet                           float64
	WorkingCapital                         float64
	TangibleAssetValue                     float64
	NetCurrentAssetValue                   float64
	InvestedCapital                        float64
	AverageReceivables                     float64
	AveragePayables                        float64
	AverageInventory                       float64
	DaysSalesOutstanding                   float64
	DaysPayablesOutstanding                float64
	DaysOfInventoryOnHand                  float64
	ReceivablesTurnover                    float64
	PayablesTurnover                       float64
	InventoryTurnover                      float64
	Roe                                    float64
	CapexPerShare                          float64
}

type StockData struct {
	Symbol                                 string
	Datetime                               string
	Close                                  float64
	High                                   float64
	Low                                    float64
	Open                                   float64
	Volume                                 float64
	Period                                 string
	OperatingCashFlowPerShare              float64
	FreeCashFlowPerShare                   float64
	CashPerShare                           float64
	PriceToSalesRatio                      float64
	DividendYield                          float64
	PayoutRatio                            float64
	RevenuePerShare                        float64
	NetIncomePerShare                      float64
	BookValuePerShare                      float64
	TangibleBookValuePerShare              float64
	ShareholdersEquityPerShare             float64
	InterestDebtPerShare                   float64
	MarketCap                              float64
	EnterpriseValue                        float64
	PeRatio                                float64
	Pocfratio                              float64
	PfcfRatio                              float64
	Pbratio                                float64
	PtbRatio                               float64
	EvToSales                              float64
	EnterpriseValueOverEBITDA              float64
	EvToOperatingCashFlow                  float64
	EarningsYield                          float64 // will have to calcualte this for every stock
	FreeCashFlowYield                      float64 // Not useful for model, but will calculate for future predictions
	DebtToEquity                           float64
	DebtToAssets                           float64
	NetDebtToEBITDA                        float64
	CurrentRatio                           float64
	InterestCoverage                       float64
	IncomeQuality                          float64
	SalesGeneralAndAdministrativeToRevenue float64
	ResearchAndDevelopmentToRevenue        float64
	IntangiblesToTotalAssets               float64
	CapexToOperatingCashFlow               float64
	CapexToRevenue                         float64
	CapexToDepreciation                    float64
	StockBasedCompensationToRevenue        float64
	GrahamNumber                           float64
	Roic                                   float64
	ReturnOnTangibleAssets                 float64
	GrahamNetNet                           float64
	WorkingCapital                         float64
	TangibleAssetValue                     float64
	NetCurrentAssetValue                   float64
	InvestedCapital                        float64
	AverageReceivables                     float64
	AveragePayables                        float64
	AverageInventory                       float64
	DaysSalesOutstanding                   float64
	DaysPayablesOutstanding                float64
	DaysOfInventoryOnHand                  float64
	ReceivablesTurnover                    float64
	PayablesTurnover                       float64
	InventoryTurnover                      float64
	Roe                                    float64
	CapexPerShare                          float64
}

// streamlined Data for generalized model
type StockDataImproved struct {
	Symbol                                    string
	Datetime                                  string
	Close                                     float64
	High                                      float64
	Low                                       float64
	Open                                      float64
	Volume                                    float64
	Period                                    string
	OperatingCashFlowPerShare                 float64
	YearOverYearRateOperatingCashFlowPerShare float64
	FreeCashFlowPerShare                      float64
	YearOverYearRateFreeCashFlowPerShare      float64
	CashPerShare                              float64
	PriceToSalesRatio                         float64
	PayoutRatio                               float64
	RevenuePerShare                           float64
	YearOverYearRateRevenuePerShare           float64
	BookValuePerShare                         float64
	YearOverYearRateBookValuePerShare         float64
	MarketCap                                 float64
	PeRatio                                   float64
	PfcfRatio                                 float64
	EvToOperatingCashFlow                     float64
	YearOverYearRateEvToOperatingCashFlow     float64
	NetDebtToEBITDA                           float64
	YearOverYearRateNetDebtToEBITDA           float64
	StockBasedCompensationToRevenue           float64
	GrahamNumber                              float64
	YearOverYearRateGrahamNumber              float64
	Roic                                      float64
	YearOverYearRateRoic                      float64
	Roe                                       float64
	CapexPerShare                             float64
	MovingAverage50Days                       float64
	MovingAverage200Days                      float64
}

// streamlined Data for generalized model
type StockDataImproved2 struct {
	Symbol                                          string
	Datetime                                        string
	Close                                           float64
	High                                            float64
	Low                                             float64
	Open                                            float64
	Volume                                          float64
	Period                                          string
	OperatingCashFlowPerShare                       float64
	YearOverYearRateOperatingCashFlowPerShare       float64
	FreeCashFlowPerShare                            float64
	YearOverYearRateFreeCashFlowPerShare            float64
	CashPerShare                                    float64
	YearOverYearRateCashPerShare                    float64
	PriceToSalesRatio                               float64
	YearOverYearRatePriceToSalesRatio               float64
	PayoutRatio                                     float64
	YearOverYearRatePayoutRatio                     float64
	RevenuePerShare                                 float64
	YearOverYearRateRevenuePerShare                 float64
	BookValuePerShare                               float64
	YearOverYearRateBookValuePerShare               float64
	MarketCap                                       float64
	PeRatio                                         float64
	YearOverYearRatePeRatio                         float64
	PfcfRatio                                       float64
	YearOverYearRatePfcfRatio                       float64
	EvToOperatingCashFlow                           float64
	YearOverYearRateEvToOperatingCashFlow           float64
	NetDebtToEBITDA                                 float64
	YearOverYearRateNetDebtToEBITDA                 float64
	StockBasedCompensationToRevenue                 float64
	YearOverYearRateStockBasedCompensationToRevenue float64
	GrahamNumber                                    float64
	YearOverYearRateGrahamNumber                    float64
	Roic                                            float64
	YearOverYearRateRoic                            float64
	Roe                                             float64
	YearOverYearRateRoe                             float64
	CapexPerShare                                   float64
	YearOverYearRateCapexPerShare                   float64
	MovingAverage50Days                             float64
	MovingAverage200Days                            float64
}
