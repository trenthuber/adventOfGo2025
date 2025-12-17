// https://adventofcode.com/2025/day/10

// NOTE: Part 2 runs too slow to check for correctness

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	total := 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		lights := tokens[0][1 : len(tokens[0])-1]
		var goal uint16
		for i := range lights {
			goal <<= 1
			if lights[len(lights)-i-1] == '#' {
				goal |= 1
			}
		}

		schemes := tokens[1 : len(tokens)-1]
		buttons := make([]uint16, len(schemes))
		for i, scheme := range schemes {
			digits := strings.Split(scheme[1:len(scheme)-1], ",")
			for _, digit := range digits {
				n, _ := strconv.Atoi(digit)
				buttons[i] |= 1 << n
			}
		}

		for i := 1; i <= len(buttons); i++ {
			if configureLights(goal, buttons, i) {
				total += i
				break
			}
		}
	}
	println(total)
}

func configureLights(goal uint16, buttons []uint16, depth int) bool {
	if goal == 0 {
		return true
	}
	if depth == 0 {
		return false
	}
	for i := range buttons {
		temp := make([]uint16, len(buttons))
		copy(temp, buttons)
		slices.Delete(temp, i, i+1)
		temp = temp[:len(temp)-1]
		if configureLights(goal^buttons[i], temp, depth-1) {
			return true
		}
	}
	return false
}

func part2(scanner *bufio.Scanner) {
	total := 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		goalToken := tokens[len(tokens)-1]
		var goal []int
		for _, digit := range strings.Split(goalToken[1:len(goalToken)-1], ",") {
			n, _ := strconv.Atoi(digit)
			goal = append(goal, n)
		}

		buttonTokens := tokens[1 : len(tokens)-1]
		buttons := make([][]int, len(buttonTokens))
		for i, buttonToken := range buttonTokens {
			buttons[i] = make([]int, len(goal))
			for _, digit := range strings.Split(buttonToken[1:len(buttonToken)-1], ",") {
				n, _ := strconv.Atoi(digit)
				buttons[i][n] = 1
			}
		}

		total += distanceFromOrigin(goal, buttons)
	}
	println(total)
}

func distanceFromOrigin(position []int, buttons [][]int) (distance int) {
	if len(buttons) == 0 {
		return -1
	}

	m := math.MaxInt
	for i, n := range buttons[0] {
		if n == 1 && position[i] < m {
			m = position[i]
		}
	}
	for i := range buttons[0] {
		position[i] -= buttons[0][i] * m
	}
	if slices.ContainsFunc(position, func(a int) bool { return a != 0 }) {
		distance = -1
		for i := m; i >= 0; i-- {
			d := distanceFromOrigin(position, buttons[1:])
			if d != -1 && (distance == -1 || i+d < distance) {
				distance = i + d
			}
			for j := range buttons[0] {
				position[j] += buttons[0][j]
			}
		}
		for i := range buttons[0] {
			position[i] -= buttons[0][i]
		}
	} else {
		distance = m
		for i := range buttons[0] {
			position[i] += buttons[0][i] * m
		}
	}
	return
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
