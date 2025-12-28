package main

import "fmt"

func main() {

	var revenue float64
	var exspenses float64
	var taxRate float64

	fmt.Print("Enter the total revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the total expenses: ")
	fmt.Scan(&exspenses)

	fmt.Print("Enter the tax rate (in %): ")
	fmt.Scan(&taxRate)

	profitBeforeTax := revenue - exspenses

	taxAmount := profitBeforeTax * (taxRate / 100)
	profitAfterTax := profitBeforeTax - taxAmount

	fmt.Println("Profit Before Tax:", profitBeforeTax)
	fmt.Println("Tax Amount:", taxAmount)
	fmt.Println("Profit After Tax:", profitAfterTax)
}
