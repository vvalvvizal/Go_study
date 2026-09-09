package main

import "fmt"

func return1(num1, num2 int) (int, int) {
	var quotient = (int)(num1 / num2)
	var remainder = num1 % num2
	return quotient, remainder
}

func main() {
	quotient, remainder := return1(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

}
