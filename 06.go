// https://adventofcode.com/2025/day/6

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Fields(scanner.Text()))
	}

	ops, total := lines[len(lines)-1], 0
	for j, op := range ops {
		var (
			opFunc   func(int, int) int
			subtotal int
		)
		switch op {
		case "+":
			opFunc = func(a, b int) int { return a + b }
			subtotal = 0
		case "*":
			opFunc = func(a, b int) int { return a * b }
			subtotal = 1
		}

		for i := 0; i < len(lines)-1; i++ {
			num, _ := strconv.Atoi(lines[i][j])
			subtotal = opFunc(subtotal, num)
		}
		total += subtotal
	}
	println(total)
}

func part2(scanner *bufio.Scanner) {
	var lines [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, len(line))
		copy(row, line)
		lines = append(lines, row)
	}

	nums, index := [][]int{[]int{}}, 0
	for j := 0; j < len(lines[0]); j++ {
		num := 0
		for i := 0; i < len(lines)-1; i++ {
			if lines[i][j] != byte(' ') {
				num *= 10
				num += int(lines[i][j] - byte('0'))
			}
		}
		if num == 0 {
			nums = append(nums, []int{})
			index++
			continue
		}
		nums[index] = append(nums[index], num)
	}

	ops, total := strings.Fields(string(lines[len(lines)-1])), 0
	for j, op := range ops {
		var (
			opFunc   func(int, int) int
			subtotal int
		)
		switch op {
		case "+":
			opFunc = func(a, b int) int { return a + b }
			subtotal = 0
		case "*":
			opFunc = func(a, b int) int { return a * b }
			subtotal = 1
		}

		for _, num := range nums[j] {
			subtotal = opFunc(subtotal, num)
		}
		total += subtotal
	}
	println(total)
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
