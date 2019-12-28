package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbers := make([]int, 3)
	numbers = nil
	var s string

	for {
		fmt.Println("Type a integer or 'X' to quit: ")
		fmt.Scan(&s)

		if s == "X" {
			break
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}

		numbers = append(numbers, n)
		sort.Ints(numbers)
		fmt.Println(numbers)
	}
}
