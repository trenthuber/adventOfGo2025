// https://adventofcode.com/2025/day/3

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func part1(scanner *bufio.Scanner) {
	total := 0
	for scanner.Scan() {
		digits := scanner.Bytes()
		left, right := getMaxIndex(digits), 0
		if left == len(digits)-1 {
			right = left
			left = getMaxIndex(digits[:right])
		} else {
			right = getMaxIndex(digits[left+1:]) + left + 1
		}
		total += 10*int(digits[left]-byte('0')) + int(digits[right]-'0')
	}
	println(total)
}

func getMaxIndex(digits []byte) int {
	if len(digits) == 0 {
		return -1
	}
	var maxDigit byte = '0'
	maxIndex := 0
	for i := 0; i < len(digits); i++ {
		if digits[i] > maxDigit {
			maxDigit = digits[i]
			maxIndex = i
		}
	}
	return maxIndex
}

func part2(scanner *bufio.Scanner) {
	total := 0
	for scanner.Scan() {
		numLeft = 12
		digits := scanner.Bytes()
		indices := getMaxIndices(digits)
		slices.Sort(indices)
		subtotal := 0
		for _, i := range indices {
			subtotal *= 10
			subtotal += int(digits[i] - '0')
		}
		total += subtotal
	}
	println(total)
}

var numLeft int

func getMaxIndices(digits []byte) []int {
	if numLeft == 0 {
		return nil
	}
	i := getMaxIndex(digits)
	if i == -1 {
		return nil
	}
	numLeft--
	right := getMaxIndices(digits[i+1:])
	for j := range right {
		right[j] += i + 1
	}
	var left []int
	if numLeft != 0 {
		left = getMaxIndices(digits[:i])
	}
	left = append(left, i)
	return append(left, right...)
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
