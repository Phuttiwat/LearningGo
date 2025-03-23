package main

import "fmt"

func grade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	score := 85
	fmt.Println("your grade is", grade(score))
}
