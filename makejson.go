package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	var address string

	data := make(map[string]string)

	input := bufio.NewReader(os.Stdin)

	fmt.Print("Type a name: ")
	name, err := input.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	name = strings.TrimSuffix(name, "\n")
	data["name"] = name

	fmt.Print("Type an address: ")
	address, err = input.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	address = strings.TrimSuffix(address, "\n")
	data["address"] = address

	output, err := json.Marshal(data)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(output))
}
