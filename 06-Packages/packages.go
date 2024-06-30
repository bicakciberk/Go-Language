package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	const greeting string = "hello there friends!"
	
	//strings
	fmt.Println(strings.Contains(greeting,"hello")) //returns; true
	fmt.Println(strings.ReplaceAll(greeting,"hello","hi")) //returns; hi there friends!
	fmt.Println(strings.ToUpper(greeting)) //returns; HELLO THERE FRIENDS!
	fmt.Println(strings.Index(greeting,"the")) //returns; 6


	//sort
	var nums = []int{5,99,2,5123,5632}
	var names = []string{"Jon","Arya","Brandon","Joffey"}

	sort.Ints(nums)
	sort.Strings(names)
	fmt.Println(nums) //returns; [2 5 99 5123 5632]
	fmt.Println(names) //returns; [Arya Brandon Joffey Jon]
}
