package main

import "fmt"

func hello() {
	fmt.Println("Hello BorntoDev")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}
func main() {
	hello()
	result := plus(5, 5)
	fmt.Println("result =", result)

	result = plus3value(5, 5, 10)
	fmt.Println("result =", result)
}
