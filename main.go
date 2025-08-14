package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//	cmdm := cmdmanager.New()
		priceJobs := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJobs.Process(doneChans[index])
	}
	for _, doneChan := range doneChans {
		<-doneChan
	}
}
