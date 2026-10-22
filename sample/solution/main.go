package main

import "fmt"

func solution(num1 int, num2 int) int {
	num3 := num1 * num2
	return num3
}

func main() {
	result := solution(5, 3)
	fmt.Println(result)
}
