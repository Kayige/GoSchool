package main

import "fmt"

func main() {
	array := [...]int{2, 4, 6, 8, 10, 0, 14, 16}

	slice := array[0:5]
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	slice = append(slice, 14)
	slice = append(slice, 16)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

}
