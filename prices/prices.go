package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println("**************************************")
	fmt.Printf("📊  Tax Calculation Report (Tax Rate: %.2f%%)\n", job.TaxRate*100)
	fmt.Println("**************************************")
	fmt.Printf("%-15s | %-15s\n", "Original Price", "Price w/ Tax")
	fmt.Println("--------------------------------------")

	for orig, taxed := range result {
		fmt.Printf("%-15s | %-15s\n", orig, taxed)
	}

	fmt.Println("--------------------------------------")
	fmt.Println("✅ Calculation complete!")
	fmt.Println("**************************************")
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		if err != nil {
			fmt.Println("Error closing file: ", err)
			return
		}
		return
	}
	job.InputPrices = prices
}
