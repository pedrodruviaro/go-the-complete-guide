package main

import (
	"fmt"

	"example.com/structs/user"
)


func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User // crio uma variável que irá (atualmente é `nil`) guardar um endereço de memória

	appUser, err := user.New(
		firstName,
		lastName,
		birthdate,
	) // atribui o endereço de memória retornado pela função

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetail()
	appUser.ClearUserName()
	appUser.OutputUserDetail()

	var admin *user.Admin

	admin, adminError := user.NewAdmin("pedrodruviaro@gmail.com", "123456")

	if adminError != nil {
		fmt.Println(adminError)
		return
	}

	admin.OutputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
