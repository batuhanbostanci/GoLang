package main

import (
	"fmt"
)

func main() {

	revenue := outputText("Enter the total revenue: ")
	exspenses := outputText("Enter the total expenses: ")
	taxRate := outputText("Enter the tax rate (in %): ")

	profitBeforeTax, taxAmount, profitAfterTax := calculateFinancials(revenue, exspenses, taxRate)

	fmt.Printf("Profit Before Tax: %.2f\n", profitBeforeTax)
	fmt.Printf("Tax Amount: %.2f\n", taxAmount)
	fmt.Printf("Profit After Tax: %.2f\n", profitAfterTax)

}

func calculateFinancials(revenue, exspenses, taxRate float64) (float64, float64, float64) {
	profitBeforeTax := revenue - exspenses

	taxAmount := profitBeforeTax * (taxRate / 100)
	profitAfterTax := profitBeforeTax - taxAmount

	return profitBeforeTax, taxAmount, profitAfterTax
}

func outputText(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}
