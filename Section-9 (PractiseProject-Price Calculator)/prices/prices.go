package price

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

func (job TaxIncludedPricesJob) Process() {
	result := make(map[string]float64)

	for priceIndex, price := range job.InputPrices {
		result[priceIndex] = price * (1 + taxRate)
	}

	result[taxRate] = taxIncludedPrices

}
