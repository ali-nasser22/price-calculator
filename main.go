package main

import (
	"price-calculator/cmdmanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJobs := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJobs.Process()
		if err != nil {
			panic(err)
		}

	}
}
