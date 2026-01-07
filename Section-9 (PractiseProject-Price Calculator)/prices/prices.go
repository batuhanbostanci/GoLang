package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
)

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPricesJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("An error occured!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("An error occured while reading the txt")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()

}

func (job *TaxIncludedPricesJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)

}
