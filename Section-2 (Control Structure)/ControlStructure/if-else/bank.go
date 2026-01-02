package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	fmt.Println("Welcome to the Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("Check balance")
	fmt.Println("Deposit money")
	fmt.Println("Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("You balance is:", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter amount to deposit: ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("New balance is:", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds!")
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("New balance is:", accountBalance)
		}
	} else if choice == 4 {
		fmt.Println("Thank you for banking with us. Goodbye!")
	} else {
		fmt.Println("Invalid choice. Please select a valid option.")
	}

	fmt.Println("You have selected option:", choice)

}
