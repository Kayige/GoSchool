package main

import "fmt"

func main() {
	result := sum(10, 5)
	fmt.Println(result)
}

func sum(number1, number2 int) int {
	result := number1 + number2
	return result
}
