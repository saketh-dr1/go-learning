package main

import "fmt"

func basic_type() {
	username := "alice@rxample.com"
	userID := 1001
	balance := 2450.75
	isActive := true

	fmt.Println("User Report\n, username")
	fmt.Println("==========")
	fmt.Printf("Username : %s\n", username)
	fmt.Printf("User ID  : %d\n", userID)
	fmt.Printf("Balance  : %.2f\n", balance)
	fmt.Printf("Active   : %t\n", isActive)
}
