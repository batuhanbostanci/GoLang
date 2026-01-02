package main

import (
	"fmt"
	"math"
)
const inflationRate = 2.5
func main() {
	
	var invesmentAmount float64
	var years float64
	expectedReturnRate := 10.0

	//fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount: ")
	fmt.Scan(&invesmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	
	futureValue, futureRealvalue := calculateFutureValues(invesmentAmount, expectedReturnRate, years)

	

	formattedFV := fmt.Sprintf("Future Value of Investment: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value of Investment (adjusted for inflation): %.1f\n", futureRealvalue)

	fmt.Println(formattedFV, formattedFRV)

	//fmt.Println("Future Value of Investment:", futureValue)
	//fmt.Println("Future Real Value of Investment (adjusted for inflation):", futureRealValue)

	//fmt.Printf("Future Value of Investment: %.1f\nFuture Real Value of Investment (adjusted for inflation): %.1f\n", futureValue, futureRealValue)

	// fmt.Printf(`Future Value of Investment: %.1f
	// Future Real Value of Investment (adjusted for inflation): %.1f `, futureValue, futureRealValue)
}

func outputText(text string) {

	fmt.Print(text)
}



func calculateFutureValues(invesmentAmount , expectedReturnRate , years float64) (fv float64, rfv float64) {
	fv= invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow((1+inflationRate/100), years)
	return fv, rfv
	//return 
}