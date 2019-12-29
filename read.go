package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var path string
	var people = make([]Person, 0)

	fmt.Print("Type the path to the input file: ")
	fmt.Scan(&path)

	data, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var p = new(Person)

		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")

		p.fname = strings. Fields(line)[0]
		p.lname = strings.Fields(line)[1]

		people = append(people, *p)
	}

	for _, person := range people {
		fmt.Printf("fname: %-20s lname: %-20s\n", person.fname, person.lname)
	}
}
