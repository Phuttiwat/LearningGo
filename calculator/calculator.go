package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", prompt)
		panic(message)
	}

	return value
}

func getOperator() string {
	fmt.Print(" operator is ( + - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}
func minus(value1, value2 float64) float64 {
	return value1 - value2
}
func multiply(value1, value2 float64) float64 {
	return value1 * value2
}
func divide(value1, value2 float64) float64 {
	return value1 / value2
}
func main() {
	count := getInput(" How many number do you want: ")
	result := getInput(" Number 1: ")
	for i := 2; i <= int(count); i++ {
		operator := getOperator()
		value := getInput(" Number " + strconv.Itoa(i) + ": ")
		switch operator {
		case "+":
			result = add(result, value)
		case "-":
			result = minus(result, value)
		case "*":
			result = multiply(result, value)
		case "/":
			result = divide(result, value)
		default:
			panic("wrong operator")
		}
	}

	fmt.Printf("result = %v", result)
}
