package prices

import (
	"fmt"

	"example.com/concu/conversion"
	"example.com/concu/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPricesJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)

		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)

		return err
	}

	job.InputPrices = prices
	return nil

}

func (job *TaxIncludedPricesJob) Process(doneChan chan bool, errorChan chan error) {

	err := job.LoadData()

	//errorChan <- errors.New("An error!")

	if err != nil {
		//return err
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)

	doneChan <- true

}
