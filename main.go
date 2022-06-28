package main

import (
	"fmt"
	"warwuck/warwuck"
)

func main() {
	//for infinite loop
	for {
		fmt.Print("Enter champion name ")
		var str string
		fmt.Scanln(&str)
		fmt.Println(warwuck.FindChampion(str))
	}
}
