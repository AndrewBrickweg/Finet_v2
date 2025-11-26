package analysis

import "fmt"

type Portfolios struct {
	BestPortfolio Portfolio
	Returns map[string][]float64
}

func OrchestratePortfolio(
	monthly []*StockDataMonthly,
	numPortfolios int,
	riskFreeRate float64,
	minWeight float64,
	maxWeight float64,
) (*Portfolios, error) {

	if len(monthly) == 0 {
		return nil, fmt.Errorf("no monthly data provided")
	}
	fmt.Println("DEBUG: Inside OrchestratePortfolio")

	adjClose := ExtractMonthlyAdjClosePrices(monthly)
	monthlyReturns := MonthlyStockReturns(adjClose)

	portfolios, bestPortfolio := OptimizePortfolio(monthlyReturns, numPortfolios, riskFreeRate, minWeight, maxWeight)

	if len(portfolios) == 0 {
		return nil, fmt.Errorf("no portfolios generated")
	}

	result := &Portfolios{
		BestPortfolio: bestPortfolio,
		Returns:      monthlyReturns,
	}

	return result, nil
}