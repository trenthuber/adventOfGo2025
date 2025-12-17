// https://adventofcode.com/2025/day/4

package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(scanner *bufio.Scanner) {
	var grid [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, len(line))
		copy(row, line)
		grid = append(grid, row)
	}

	println(len(getAccessableRolls(grid)))
}

func part2(scanner *bufio.Scanner) {
	var grid [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, len(line))
		copy(row, line)
		grid = append(grid, row)
	}

	total := 0
	for rolls := getAccessableRolls(grid); len(rolls) != 0; rolls = getAccessableRolls(grid) {
		total += len(rolls)
		for _, roll := range rolls {
			grid[roll[0]][roll[1]] = byte('.')
		}
	}
	println(total)
}

func getAccessableRolls(grid [][]byte) (result [][]int) {
	numRows, numCols, roll := len(grid), len(grid[0]), byte('@')
	for row, gridRow := range grid {
		for col, spot := range gridRow {
			if spot != roll {
				continue
			}
			numNeighbors := 0
			for i := -1; i <= 1; i++ {
				if row+i == -1 || row+i == numRows {
					continue
				}
				for j := -1; j <= 1; j++ {
					if col+j == -1 || col+j == numCols {
						continue
					}
					if grid[row+i][col+j] == roll {
						numNeighbors++
					}
				}
			}
			numNeighbors--
			if numNeighbors < 4 {
				result = append(result, []int{row, col})
			}
		}
	}
	return result
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
