package main

import "fmt"

func main(){
	var names = [3]string{"Arya","Ned","Jon"}
	var brands = [3]string{"Google","Meta","Microsoft"}


	fmt.Println(names) //returns [Arya Ned Jon]
	fmt.Println("names", len(names)) //returns 3
	fmt.Println(names[1]) //returns Ned

	fmt.Println(brands) //returns [Google Meta Microsoft]


	var range1 = names[0:2]
	fmt.Println(range1) //returns [Arya Ned]

}