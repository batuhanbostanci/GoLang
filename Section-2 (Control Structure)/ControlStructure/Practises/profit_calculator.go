package main

import (
	"fmt"
	"os"
)

//Goals
// 1) Validate user input
// => Show error message & exit if invalid input is provided
//  -No negative numbers
//  -No 0
// 2) Store calculated results into file

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// There are 2 way to store data into file in this task
	// 1)creating a file and writing line by line
	// file, err := os.Create("financials.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // Write them line by line
	// fmt.Fprintf(file, "EBT: %.1f\n", ebt)
	// fmt.Fprintf(file, "Profit: %.1f\n", profit)
	// fmt.Fprintf(file, "Ratio: %.3f\n", ratio)

	// 2) Using Sprintf to format the entire content and writing it at once
	// Create one big string with newline characters (\n)
	data := fmt.Sprintf("%.1f\n%.1f\n%.3f\n", ebt, profit, ratio)

	// Write the entire string at once
	os.WriteFile("financials.txt", []byte(data), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {

		panic("Invalid input provided, Do not use negative numbers or zero")
	}

	return userInput
}
