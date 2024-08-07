package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserData() { // this function can be use thought a user variable due to the "(u user)"

	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstname string, lastname string, birthdate string) (*user, error) {
	if firstname == "" || lastname == "" {
		return nil, errors.New("firstname and lastname is required")
	}
	return &user{
		firstName: firstname,
		lastName:  lastname,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser, err = newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData() // The struct is empty here
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
