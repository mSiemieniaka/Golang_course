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
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	// to jest metoda typu receiver
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""

}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}

}

func New(firstname, lastname, birthdate string) (*User, error) {
	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("nie poprawnie wpsiane dane ")
	}

	return &User{
		firstName: firstname,
		lastName:  lastname,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil

}
