package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to the Bank")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	presentOptions()

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error:", err)
		//panic("Cannot contiune, sorry!")
	}

	for {

		var choice int

		fmt.Print("Enter your choice (1-4): ")
		fmt.Scan(&choice)

		// Switch case to handle user choice
		// ***********************************************

		// IN GO , SWITCH CASE DOES NOT REQUIRE BREAK STATEMENT
		switch choice {
		case 1:
			{
				var totalBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

				if err != nil {
					fmt.Println("Error:", err)
				}
				fmt.Println("You balance is:", totalBalance)

			}
		case 2:
			{
				var depositAmount float64
				fmt.Print("Enter amount to deposit: ")
				fmt.Scan(&depositAmount)
				accountBalance += depositAmount
				fmt.Println("New balance is:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
