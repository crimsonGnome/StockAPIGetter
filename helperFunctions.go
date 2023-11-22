package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CSVconverter(destinationFileName string, stockDataArray *[]StockData) error {
	outputFile, err := os.Create(destinationFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{
		"Symbol",
		"Datetime",
		"Close",
		"High",
		"Low",
		"Open",
		"Volume",
		"Period",
		"OperatingCashFlowPerShare",
		"FreeCashFlowPerShare",
		"CashPerShare",
		"PriceToSalesRatio",
		"DividendYield",
		"PayoutRatio",
		"RevenuePerShare",
		"NetIncomePerShare",
		"BookValuePerShare",
		"TangibleBookValuePerShare",
		"ShareholdersEquityPerShare",
		"InterestDebtPerShare",
		"MarketCap",
		"EnterpriseValue",
		"PeRatio",
		"Pocfratio",
		"PfcfRatio",
		"Pbratio",
		"PtbRatio",
		"EvToSales",
		"EnterpriseValueOverEBITDA",
		"EvToOperatingCashFlow",
		"EarningsYield",
		"FreeCashFlowYield",
		"DebtToEquity",
		"DebtToAssets",
		"NetDebtToEBITDA",
		"CurrentRatio",
		"InterestCoverage",
		"IncomeQuality",
		"SalesGeneralAndAdministrativeToRevenue",
		"ResearchAndDevelopmentToRevenue",
		"IntangiblesToTotalAssets",
		"CapexToOperatingCashFlow",
		"CapexToRevenue",
		"CapexToDepreciation",
		"StockBasedCompensationToRevenue",
		"GrahamNumber",
		"Roic",
		"ReturnOnTangibleAssets",
		"GrahamNetNet",
		"WorkingCapital",
		"TangibleAssetValue",
		"NetCurrentAssetValue",
		"InvestedCapital",
		"AverageReceivables",
		"AveragePayables",
		"AverageInventory",
		"DaysSalesOutstanding",
		"DaysPayablesOutstanding",
		"DaysOfInventoryOnHand",
		"ReceivablesTurnover",
		"PayablesTurnover",
		"InventoryTurnover",
		"Roe",
		"CapexPerShare"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range *stockDataArray {
		var csvRow []string
		csvRow = append(csvRow,
			r.Symbol,
			r.Datetime,
			fmt.Sprintf("%f", r.Close),
			fmt.Sprintf("%f", r.High),
			fmt.Sprintf("%f", r.Low),
			fmt.Sprintf("%f", r.Open),
			fmt.Sprintf("%f", r.Volume),
			r.Period,
			fmt.Sprintf("%f", r.OperatingCashFlowPerShare),
			fmt.Sprintf("%f", r.FreeCashFlowPerShare),
			fmt.Sprintf("%f", r.CashPerShare),
			fmt.Sprintf("%f", r.PriceToSalesRatio),
			fmt.Sprintf("%f", r.DividendYield),
			fmt.Sprintf("%f", r.PayoutRatio),
			fmt.Sprintf("%f", r.RevenuePerShare),
			fmt.Sprintf("%f", r.NetIncomePerShare),
			fmt.Sprintf("%f", r.BookValuePerShare),
			fmt.Sprintf("%f", r.TangibleBookValuePerShare),
			fmt.Sprintf("%f", r.ShareholdersEquityPerShare),
			fmt.Sprintf("%f", r.InterestDebtPerShare),
			fmt.Sprintf("%f", r.MarketCap),
			fmt.Sprintf("%f", r.EnterpriseValue),
			fmt.Sprintf("%f", r.PeRatio),
			fmt.Sprintf("%f", r.Pocfratio),
			fmt.Sprintf("%f", r.PfcfRatio),
			fmt.Sprintf("%f", r.Pbratio),
			fmt.Sprintf("%f", r.PtbRatio),
			fmt.Sprintf("%f", r.EvToSales),
			fmt.Sprintf("%f", r.EnterpriseValueOverEBITDA),
			fmt.Sprintf("%f", r.EvToOperatingCashFlow),
			fmt.Sprintf("%f", r.EarningsYield),
			fmt.Sprintf("%f", r.FreeCashFlowYield),
			fmt.Sprintf("%f", r.DebtToEquity),
			fmt.Sprintf("%f", r.DebtToAssets),
			fmt.Sprintf("%f", r.NetDebtToEBITDA),
			fmt.Sprintf("%f", r.CurrentRatio),
			fmt.Sprintf("%f", r.InterestCoverage),
			fmt.Sprintf("%f", r.IncomeQuality),
			fmt.Sprintf("%f", r.SalesGeneralAndAdministrativeToRevenue),
			fmt.Sprintf("%f", r.ResearchAndDevelopmentToRevenue),
			fmt.Sprintf("%f", r.IntangiblesToTotalAssets),
			fmt.Sprintf("%f", r.CapexToOperatingCashFlow),
			fmt.Sprintf("%f", r.CapexToRevenue),
			fmt.Sprintf("%f", r.CapexToDepreciation),
			fmt.Sprintf("%f", r.StockBasedCompensationToRevenue),
			fmt.Sprintf("%f", r.GrahamNumber),
			fmt.Sprintf("%f", r.Roic),
			fmt.Sprintf("%f", r.ReturnOnTangibleAssets),
			fmt.Sprintf("%f", r.GrahamNetNet),
			fmt.Sprintf("%f", r.WorkingCapital),
			fmt.Sprintf("%f", r.TangibleAssetValue),
			fmt.Sprintf("%f", r.NetCurrentAssetValue),
			fmt.Sprintf("%f", r.InvestedCapital),
			fmt.Sprintf("%f", r.AverageReceivables),
			fmt.Sprintf("%f", r.AveragePayables),
			fmt.Sprintf("%f", r.AverageInventory),
			fmt.Sprintf("%f", r.DaysSalesOutstanding),
			fmt.Sprintf("%f", r.DaysPayablesOutstanding),
			fmt.Sprintf("%f", r.DaysOfInventoryOnHand),
			fmt.Sprintf("%f", r.ReceivablesTurnover),
			fmt.Sprintf("%f", r.PayablesTurnover),
			fmt.Sprintf("%f", r.InventoryTurnover),
			fmt.Sprintf("%f", r.Roe),
			fmt.Sprintf("%f", r.CapexPerShare),
		)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}
