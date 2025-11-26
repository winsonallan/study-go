package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	// Instantiating structs
	// Method 1a
	// var appUser user
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now(),
	// }

	// Method 1b
	// appUser := user{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now(),
	// }
	
	// Method 2 (But the order must be the same according to the defined type)
	// var appUser user
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	//  	time.Now(),
	// }
	
	// Method 3 - To create an empty null struct
	// var appUser user
	// appUser = user{}

	// Method 4 - Using Creation/Constructor Functions
	// appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	// Method 4b - From different package
	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetails()
	// If use regular function
	// outputUserDetails(&appUser)

	// If use method
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
