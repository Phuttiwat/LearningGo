package main

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data1 := []byte("Hi\n borndoDev")
	err := os.WriteFile("data.txt", data1, 0644)
	check(err)

	f, err := os.Create("employeeName")
	check(err)

	defer f.Close()

	data2 := []byte("Sira\n Manee")
	os.WriteFile("employeeName.txt", data2, 0644)
}
