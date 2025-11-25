// The idea with packages is to organize the code.
package main

// 'main' is a special package name that tells GO about the main entry point of the program. Other than that, you can use mostly other words like 'hello', 'goodbye', etc. If we don't have any 'main' package, when building it, go won't be able to find the entry point of the program

import "fmt"

// This function is also called 'main' so that Go would know the function to execute by default.
// This means that there should not be 2 or more 'main' functions in the same package, not in the same file
func main() {
	fmt.Print("Hello World")
}
// Not all packages should have a 'main' function though, as if we only build a utility package to be used across projects, we don't need to execute the package itself. Instead, we want to use the functions available in the package in our project that imports the utility package


// Later when building, use
// go mod init 'url' -> This initializes a module. A module in go consists of 1 or more packages
// without that, we can't use go build

// When running a go program like app.go we can:
// 	go run app.go
// 	OR if there is already a module
// 	go run .