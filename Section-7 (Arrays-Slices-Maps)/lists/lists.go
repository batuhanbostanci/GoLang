package lists

import "fmt"

func main() {

	prices := []float64{10.5, 20.0, 30.75, 40.25, 50.0}

	fmt.Println(prices[0:1])

	prices = append(prices, 60.5)
	fmt.Println(prices)

	discountedPrices := []float64{5.0, 15.0, 25.0}

	// Append multiple values to the slice with special syntax `...` that unpacks the slice and adds each element individually
	prices = append(prices, discountedPrices...)

	fmt.Println(prices)

}
