package main

import "fmt"

func main() {
	var number float64

	fmt.Print("Type a float point number: ")
	fmt.Scan(&number)
	fmt.Println(int64(number))
}
