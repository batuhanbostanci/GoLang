package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {

	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func main() {

	var accountBalance = getBalanceFromFile()

	for {

		fmt.Println("Welcome to the Bank")
		fmt.Println("What do you want to do?")
		fmt.Println("Check balance")
		fmt.Println("Deposit money")
		fmt.Println("Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Enter your choice (1-4): ")
		fmt.Scan(&choice)

		// If else statement to handle user choice
		// ***********************************************
		// if choice == 1 {
		// 	fmt.Println("You balance is:", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
		// 	fmt.Print("Enter amount to deposit: ")
		// 	fmt.Scan(&depositAmount)
		// 	accountBalance += depositAmount
		// 	fmt.Println("New balance is:", accountBalance)
		// } else if choice == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("Enter amount to withdraw: ")
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Insufficient funds!")
		// 	} else {
		// 		accountBalance -= withdrawAmount
		// 		fmt.Println("New balance is:", accountBalance)
		// 	}
		// } else if choice == 4 {
		// 	fmt.Println("Thank you for banking with us. Goodbye!")
		// 	break
		// } else {
		// 	fmt.Println("Invalid choice. Please select a valid option.")
		// 	return
		// }

		// Switch case to handle user choice
		// ***********************************************

		// IN GO , SWITCH CASE DOES NOT REQUIRE BREAK STATEMENT
		switch choice {
		case 1:
			{
				fmt.Println("You balance is:", getBalanceFromFile())

			}
		case 2:
			{
				var depositAmount float64
				fmt.Print("Enter amount to deposit: ")
				fmt.Scan(&depositAmount)
				accountBalance += depositAmount
				fmt.Println("New balance is:", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 3:
			{
				var withdrawAmount float64
				fmt.Print("Enter amount to withdraw: ")
				fmt.Scan(&withdrawAmount)
				if withdrawAmount > accountBalance {
					fmt.Println("Insufficient funds!")
				} else {
					accountBalance -= withdrawAmount
					fmt.Println("New balance is:", accountBalance)
				}
				writeBalanceToFile(accountBalance)
			}
		case 4:
			{
				fmt.Println("Thank you for banking with us. Goodbye!")
				return

			}
		// Default case
		default:
			{
				fmt.Println("Invalid choice. Please select a valid option.")
			}

		}

		fmt.Println("You have selected option:", choice)

	}

}
