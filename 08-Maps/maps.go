package main

import "fmt"

func main(){
	//adding value to  map
	var myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	fmt.Println(myMap) //retunrs; map[one:1 three:3 two:2]

	//reading value from map
	fmt.Println(myMap["one"]) //returns; 1

	//deleting value from map
	delete(myMap,"one")
	fmt.Println(myMap) //returns; map[three:3 two:2]

	value1,exist := myMap["two"]
	if exist {
		fmt.Println("myMap is contains value1",value1) //returns; myMap is contains value1 2
	}else{
		fmt.Println("myMap is not contains value1")
	}
}