package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	//var prices []float64 = []float64{10,20,30}

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()

	}

}
