package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, error1 := user.NewUser(firstName, lastName, birthDate)

	if error1 != nil {
		fmt.Println(error1)
		return
	}

	// Or using pointer return type
	var appUserPointer *user.User
	appUserPointer, error := user.NewUserPointer(firstName, lastName, birthDate)
	if error != nil {
		fmt.Println(error)
		return
	}

	admin := user.NewAdmin("test@example.com", "adminpass")

	admin.PointerOutputUserDetails()
	admin.ClearUserName()
	admin.PointerOutputUserDetails()
	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.PointerOutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()

	appUserPointer.PointerOutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	// Scanln skip whitespaces so user can't enter spaces in names
	fmt.Scanln(&value)
	return value
}
