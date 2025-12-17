// https://adventofcode.com/2025/day/5

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	var fresh [][2]int

	scanner.Scan()
	for line := scanner.Text(); line != ""; line = scanner.Text() {
		idRange := strings.Split(line, "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])
		fresh = append(fresh, [2]int{start, end})
		scanner.Scan()
	}

	total := 0
	for scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())
		for _, idRange := range fresh {
			if id >= idRange[0] && id <= idRange[1] {
				total++
				break
			}
		}
	}
	println(total)
}

func part2(scanner *bufio.Scanner) {
	var fresh [][2]int

	scanner.Scan()
	for line := scanner.Text(); line != ""; line = scanner.Text() {
		idRange := strings.Split(line, "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])
		fresh = append(fresh, [2]int{start, end})
		scanner.Scan()
	}

	slices.SortFunc(fresh, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})

	for i := 0; i < len(fresh)-1; i++ {
		if fresh[i+1][0] <= fresh[i][1] {
			fresh[i+1][0] = fresh[i][0]
			if fresh[i+1][1] < fresh[i][1] {
				fresh[i+1][1] = fresh[i][1]
			}
			fresh[i] = [2]int{1, 0}
		}
	}

	total := 0
	for _, idRange := range fresh {
		total += idRange[1] - idRange[0] + 1
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
