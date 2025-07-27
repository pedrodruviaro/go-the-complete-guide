package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func (user user) outputUserDetail() {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func (user *user) clearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == ""  || birthdate == "" {
		return nil, errors.New("invalid body")
	}

	return &user{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(
		firstName,
		lastName,
		birthdate,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(appUser)

	appUser.outputUserDetail()
	appUser.clearUserName()
	appUser.outputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
