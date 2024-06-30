package main

import "fmt"

func sumAndDiff(a int, b int) (int, int) {
	var sum = a + b
	var diff = a - b
	return sum, diff
}

func main() {
	a, b := 10, 5
	var sum,diff = sumAndDiff(a,b)
	
	fmt.Println(sum)
	fmt.Println(diff)

}