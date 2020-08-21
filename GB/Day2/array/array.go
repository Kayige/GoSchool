package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var OperatingsystemList [6]string

	OperatingsystemList[0] = "Windwos XP"
	OperatingsystemList[1] = "Linux 1.0"
	OperatingsystemList[2] = "raspbian teensy"
	OperatingsystemList[3] = "MacOS"
	OperatingsystemList[4] = "IOS"
	OperatingsystemList[5] = "Google Android"

	for i := 0; i < len(OperatingsystemList); i++ {

		fmt.Println(OperatingsystemList[i])
	}

	OperatingsystemList[0] = "Windows 10"

	for i := 0; i < len(OperatingsystemList); i++ {
		fmt.Println(OperatingsystemList[i])
	}

	var NewOperatingSystemList [9]string

	for i := 0; i < len(OperatingsystemList); i++ {
		NewOperatingSystemList[i] = OperatingsystemList[i]
	}

	NewOperatingSystemList[6] = " Ubuntu"
	NewOperatingSystemList[7] = "Ms-DOS"
	NewOperatingSystemList[8] = "Solaris"
	fmt.Println(OperatingsystemList)
	fmt.Println(NewOperatingSystemList)
}
