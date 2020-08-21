package main

import (
	"fmt"
	"strconv"
)

type currency struct {
	currencyName string
	rate         float64
}

var m map[string]currency

func main() {

	m = make(map[string]currency)

	m["USD"] = currency{"US dollar", 1.1318}
	m["JPY"] = currency{"Japanese yen", 121.05}
	m["GBP"] = currency{"Pound Sterling", 0.90630}
	m["CNY"] = currency{"Chinese yuan renminbi", 7.9944}
	m["SGD"] = currency{"Singapore dollar", 1.5743}
	m["MYR"] = currency{"Malaysian ringgit", 4.8390}

	fmt.Println(m["USD"].currencyName, m["USD"].rate)

	var currencyFrom string
	var currencyAmt string
	var currencyTo string
	fmt.Println("Enter your currency: ")
	fmt.Scanln(&currencyFrom)
	fmt.Println("Enter your currency amount: ")
	fmt.Scanln(&currencyAmt)
	fmt.Println("Enter currency to convert to: ")
	fmt.Scanln(&currencyTo)

	currencyAmtValue, _ := strconv.ParseInt(currencyAmt, 10, 0)

	result := (float64(currencyAmtValue) / m[currencyFrom].rate) * m[currencyTo].rate

	fmt.Println(result)
}
