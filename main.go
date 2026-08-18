package main

import "fmt"

func main() {
	/*
	   	//EXPERIEMNT 1
	   	//QUESTION 1
	   	//hello code
	   	//fmt.Println("Hello! Kabir")

	   	// Declare two integer and two float variables
	   	// and perform addition, subtraction, and multiplication
	   	//  using integer and floating-point types.

	   	//num1 := 22
	   	//num2 := 10

	   	//fmt.Printf("Addition: %d\n", num1+num2)
	   	//fmt.Printf("Subtraction: %d\n", num1-num2)
	   	//fmt.Printf("Multiplication: %d\n", num1*num2)
	   	//fmt.Printf("Division: %d\n", num1/num2)

	   	//Float function

	   	var num1 float64 = 34.2
	   	var num2 float64 = 67.7

	   	fmt.Printf("Addition: %.2f\n", num1+num2)
	   	fmt.Printf("Subtraction: %.2f\n", num1-num2)
	   	fmt.Printf("Multiplication: %.2f\n", num1*num2)
	   	fmt.Printf("Division: %.2f\n", num1/num2)
	   }
	*/
	//QUESTION 2
	//Accept numeric input.
	//Perform addition, subtraction, and multiplication using integer types.
	//Perform addition, subtraction, and multiplication using floating-point types.
	//Use if statements and loops to validate the input
	// (reject non-numeric or out-of-range values) and handle different scenarios.
	for {

		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Integer Operations")
		fmt.Println("2. Floating Operations")
		fmt.Println("3. Exit")

		var choice int

		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			return
		}

		if choice < 1 || choice > 3 {
			fmt.Println("Invalid choice! Please enter 1, 2, or 3.")
			continue
		}

		if choice == 1 {

			var num1, num2 int

			fmt.Print("Enter first integer: ")
			fmt.Scan(&num1)

			fmt.Print("Enter second integer: ")
			fmt.Scan(&num2)

			fmt.Println("Addition:", num1+num2)
			fmt.Println("Subtraction:", num1-num2)
			fmt.Println("Multiplication:", num1*num2)

		} else if choice == 2 {

			var num1, num2 float64

			fmt.Print("Enter first float number: ")
			fmt.Scan(&num1)

			fmt.Print("Enter second float number: ")
			fmt.Scan(&num2)

			fmt.Println("Addition:", num1+num2)
			fmt.Println("Subtraction:", num1-num2)
			fmt.Println("Multiplication:", num1*num2)

		} else if choice == 3 {

			fmt.Println("Program ended.")
			break
		}
	}

}
