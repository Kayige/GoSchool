package main

import "fmt"

func main() {
	slice := make([]string, 10, 11)
	slice[0] = "|1st slice position|"
	slice[1] = "|2nd Slice position|"
	slice[2] = "|3rd Slice position|"
	slice[3] = "|3rd Slice position|"

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	slice = append(slice, "|4th slice position|")
	slice = append(slice, "|5th slice position|")

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
}
