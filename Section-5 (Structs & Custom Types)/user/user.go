package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     // Embedding User struct
}

// Method with value receiver, it gets a copy of the struct
func (u User) OutputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt.Year())

}

func (u *User) PointerOutputUserDetails() {

	//*********************************
	//This is also valid and old syntax on cpp style but go allows to skip the dereferencing step and handles it internally
	//(*u).firstName = "ChangedName"
	//*********************************

	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt.Year())

}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName string, lastName string, birthDate string) (User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return User{}, errors.New("normal constructer first name, last name and birthdate are required fields")
	}

	return User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}
}

// Alternativly, we can make it pointer return type
// use New keyword to create the struct and return its pointer this time there is two ways to define constructer that is why we named it NewUserPointer
func NewUserPointer(firstName string, lastName string, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("pointer constructer first name, last name and birthdate are required fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
