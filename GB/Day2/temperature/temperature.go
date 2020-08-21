package main

import "fmt"

func main() {
	temperature := [24]float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3}

	total := 0.00
	for i := 0; i < len(temperature); i++ {
		total += temperature[i]
	}
	average := total / float64(len(temperature))
	fmt.Printf("%.2f", average)

}
