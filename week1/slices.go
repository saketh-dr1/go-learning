package main

import "fmt"

func slices_basics() {
	var fruitList = []string{}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "mango")

	fmt.Println(fruitList)
}
