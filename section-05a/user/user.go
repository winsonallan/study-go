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

// Struct Embedding
type Admin struct {
	email string
	password string
	User
}

// 1. Regular Function using struct data
/*
func outputUserDetails(u *user) {

	// Pointers to structs are exception, in Go it's allowed to do u.firstName to read the value of the firstName, instead of having to dereference it first.

	// Although, we can dereference it first if we want to do it the technically correct way, and that is to do something like this:
	// (*u).firstName instead of just u.firstName
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
*/

// 2. Attaching the function to the struct, which changes the function into a method.
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// The (u user) above is called as 'receiver' which turns a function into a method

// 3. If you want to create a method that aims to change the value stored inside a struct, you have to use Pointers in the receiver
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin (email, password string) Admin {
	return Admin {
		email: email,
		password: password,
		User: User {
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

// Creation/Constructor Functions
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}