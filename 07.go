// https://adventofcode.com/2025/day/7

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func part1(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, len(line))
		copy(row, line)
		grid = append(grid, row)
	}

	println(countSpliters(1, bytes.Index(grid[0], []byte{byte('S')})))
}

var grid [][]byte

func countSpliters(i, j int) int {
	for ; i < len(grid); i++ {
		if grid[i][j] == byte('|') {
			return 0
		}
		if grid[i][j] == byte('^') {
			return countSpliters(i, j-1) + countSpliters(i, j+1) + 1
		}
		grid[i][j] = byte('|')
	}
	return 0
}

func part2(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, len(line))
		copy(row, line)
		grid = append(grid, row)
	}
	cache = make([][]int, len(grid))
	for i := range cache {
		cache[i] = make([]int, len(grid[0]))
	}

	println(countTimelines(1, bytes.Index(grid[0], []byte{byte('S')})))
}

var cache [][]int

func countTimelines(i, j int) int {
	for ; i < len(grid); i++ {
		if grid[i][j] == byte('^') {
			if cache[i][j] == 0 {
				cache[i][j] = countTimelines(i, j-1) + countTimelines(i, j+1)
			}
			return cache[i][j]
		}
	}
	return 1
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
