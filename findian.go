package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string

	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Type a string: ")

	str, err := input.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	str = strings.TrimSuffix(str, "\n")
	str = strings.ToLower(str)

	if strings.HasPrefix(str, "i") &&
		strings.ContainsAny("a", str[1:len(str)-1]) &&
		strings.HasSuffix(str, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
