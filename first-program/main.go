package main

import "fmt"

func main() {
	var name string //string variable.

	welcome := "Welcome"

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name)

	fmt.Printf("%s, %s!\n", welcome, name)

}
