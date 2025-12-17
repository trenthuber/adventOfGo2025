package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(scanner *bufio.Scanner) {
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func part2(scanner *bufio.Scanner) {
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func main() {
	var part func(*bufio.Scanner)
	filename := "input.txt"
	usage := func() {
		fmt.Printf("usage: %v {1|2} [test]\n", os.Args[0])
		os.Exit(1)
	}
	switch len(os.Args) {
	case 3:
		if os.Args[2] == "test" {
			filename = "test.txt"
		} else {
			usage()
		}
		fallthrough
	case 2:
		switch os.Args[1] {
		case "1":
			part = part1
		case "2":
			part = part2
		default:
			usage()
		}
	case 1:
		usage()
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open %v\n", filename)
		os.Exit(1)
	}
	defer file.Close()
	part(bufio.NewScanner(file))
}
