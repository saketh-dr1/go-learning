package main

import (
	"fmt"
	"os"
	"strconv"
)

func expense_category() {
	if len(os.Args) < 2 {
		fmt.Println("Enter an amount to start categorizing...")
		os.Exit(1)
	}

	expenseStr := os.Args[1]

	expense, err := strconv.Atoi(expenseStr)

	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}

	// if expense < 500 {
	// 	fmt.Println("small")
	// } else if expense > 500 && expense < 5000 {
	// 	fmt.Println("medium")
	// } else {
	// 	fmt.Println("Large")
	// }

	// swtich case logic
	switch {
	case expense < 500:
		fmt.Println("small")
	case expense > 500 && expense < 5000:
		fmt.Println("medium")
	case expense > 5000:
		fmt.Println("large")
	}
}
