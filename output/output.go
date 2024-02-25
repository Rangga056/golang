package main

import "fmt"

func main() {
	//NOTE Print() will print its arguments with their default format
	var first_name string = "muhammad"
	var last_name string = "falah"

	//NOTE you can use "\n" to print a new line
	//Print will add space between argumentsif neither are string

	fmt.Print(first_name, "\n", last_name, "\n")
	fmt.Print(first_name, last_name)

}
