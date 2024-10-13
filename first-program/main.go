package main

import "fmt"

func main() {
	var nome string //string variable.

	welcome := "Welcome"

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &nome)

	fmt.Printf("%s, %s!\n", welcome, nome)

}
