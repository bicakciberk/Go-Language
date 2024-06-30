package main

import "fmt"

func main() {
	var name string = "Arin" //thats no needed
	var secondName = "Arya" //that will work
	userName:="arya123" //that will work, its used when we assign a new variables
	userName="arya321" // and that is changing userName

	var age int = 19
	var isMarried bool = true



	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(secondName)
	fmt.Println(userName)
}