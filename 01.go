// https://adventofcode.com/2025/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) {
	var pos, zeros int64 = 50, 0
	for scanner.Scan() {
		rot := scanner.Bytes()
		dir := rot[0]
		num, _ := strconv.ParseInt(string(rot[1:]), 0, 0)
		if dir == 'L' {
			num = 100 - num
		}
		if pos = (pos + num) % 100; pos == 0 {
			zeros++
		}
	}
	println(zeros)
}

func part2(scanner *bufio.Scanner) {
	var pos, zeros int64 = 50, 0
	for scanner.Scan() {
		rot := scanner.Bytes()
		dir := rot[0]
		num, _ := strconv.ParseInt(string(rot[1:]), 0, 0)
		zeros += num / 100
		num %= 100
		if dir == 'L' {
			num *= -1
		}
		if pos == 0 {
			pos = (num + 100) % 100
		} else {
			pos += num
			if dir == 'L' {
				if pos < 1 {
					pos = (pos + 100) % 100
					zeros++
				}
			} else {
				if pos > 99 {
					pos %= 100
					zeros++
				}
			}
		}
	}
	println(zeros)
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
