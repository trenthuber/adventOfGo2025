// https://adventofcode.com/2025/day/11

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	initDevices(scanner)

	println(numPaths("you", "out", make(map[string]int)))
}

func part2(scanner *bufio.Scanner) {
	initDevices(scanner)

	a, b := "fft", "dac"
	middle := numPaths(a, b, make(map[string]int))
	if middle == 0 {
		a, b = b, a
		middle = numPaths(a, b, make(map[string]int))
	}
	println(numPaths("svr", a, make(map[string]int)) *
		middle * numPaths(b, "out", make(map[string]int)))
}

var devices map[string][]string

func initDevices(scanner *bufio.Scanner) {
	devices = make(map[string][]string)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ":")
		devices[names[0]] = strings.Split(names[1], " ")[1:]
	}
}

func numPaths(start, end string, cache map[string]int) int {
	if start == end {
		return 1
	}
	if _, exists := cache[start]; !exists {
		for _, output := range devices[start] {
			cache[start] += numPaths(output, end, cache)
		}
	}
	return cache[start]
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
		if os.Args[2] != "test" {
			usage()
		}
		switch os.Args[1] {
		case "1":
			filename = "test1.txt"
		case "2":
			filename = "test2.txt"
		default:
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
