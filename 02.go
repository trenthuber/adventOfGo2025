// https://adventofcode.com/2025/day/2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	scanner.Scan()
	idRanges := strings.Split(scanner.Text(), ",")
	total := 0
	for _, idRange := range idRanges {
		ids := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])
		for i := start; i <= end; i++ {
			idString := strconv.Itoa(i)
			l := len(idString) / 2
			if idString[:l] == idString[l:] {
				total += i
			}
		}
	}
	println(total)
}

func part2(scanner *bufio.Scanner) {
	scanner.Scan()
	idRanges := strings.Split(scanner.Text(), ",")
	total := 0
	for _, idRange := range idRanges {
		ids := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])
		for i := start; i <= end; i++ {
			idString := strconv.Itoa(i)
			l := len(idString)
		nextDivision:
			for numDiv := 2; numDiv <= l; numDiv++ {
				if l%numDiv == 0 {
					subLen := l / numDiv
					for j := 0; j < numDiv; j++ {
						if idString[:subLen] != idString[j*subLen:(j+1)*subLen] {
							continue nextDivision
						}
					}
					total += i
					break
				}
			}
		}
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
