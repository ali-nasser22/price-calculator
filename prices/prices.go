package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/interfaces"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
	IOManager         interfaces.IOManagerInt `json:"-"`
}

func NewTaxIncludedPriceJob(io interfaces.IOManagerInt, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: io,
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		return
	}
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
	job.TaxIncludedPrices = result
	err = job.IOManager.WriteResult(job)
	if err != nil {
		errorChan <- err
		return
	}
	doneChan <- true

}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		if err != nil {
			fmt.Println("Error closing file: ", err)
			return err
		}
		return err
	}
	job.InputPrices = prices
	return nil
}
