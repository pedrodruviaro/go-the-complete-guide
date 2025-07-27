package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User // ou `user User`
}

func (user User) OutputUserDetail() {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

/* - retorna um endereço de memória
   - isso significa que quem chamar essa função terá acesso direto ao objeto criado na memória, e poderá modificá-lo
*/
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("invalid body")
	}

	// ao invés de guardar em uma variável e retornar ela, enviamos o endereço de memória onde os valores foram guardados
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("invalid body")
	}

	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "01/01/1001",
			createdAt: time.Now(),
		},
	}, nil
}