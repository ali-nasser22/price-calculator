package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJobs := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJobs.Process()

	}
}
