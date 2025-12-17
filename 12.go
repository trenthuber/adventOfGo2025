// https://adventofcode.com/2025/day/12

// NOTE: Part 1 runs too slow to check for correctness, and so Part 2 remains inaccessable to me.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line[len(line)-1] == ':' {
			var s [3][3]bool
			for i := 0; ; i++ {
				scanner.Scan()
				line = scanner.Text()
				if len(line) == 0 {
					break
				}

				for j, c := range line {
					if c == '#' {
						s[i][j] = true
					}
				}
			}

			var syms [8][3][3]bool
			for rot := 0; rot < 4; rot++ {
				for refl := 0; refl < 2; refl++ {
					syms[2*rot+refl] = s
					for i := range s {
						s[i][0], s[i][2] = s[i][2], s[i][0]
					}
				}
				temp := s
				for i := range s {
					for j := range s[0] {
						s[i][j] = temp[2-j][i]
					}
				}
			}

			var shape [][3]uint
		nextShape:
			for i := 0; i < 8; i++ {
			nextPair:
				for j := i + 1; j < 8; j++ {
					for r := range syms[i] {
						for c := range syms[i][r] {
							if syms[i][r][c] != syms[j][r][c] {
								continue nextPair
							}
						}
					}
					continue nextShape
				}

				var data [3]uint
				for r, row := range syms[i] {
					for _, item := range row {
						data[r] <<= 1
						if item {
							data[r]++
						}
					}
					temp := data[r]
					for i := 0; i < 16; i++ {
						data[r] <<= 3
						data[r] |= temp
					}
				}
				shape = append(shape, data)
			}
			shapes = append(shapes, shape)
		} else {
			parts := strings.Split(line, ":")
			var dimensions, indices []uint8
			for _, value := range strings.Split(parts[0], "x") {
				n, _ := strconv.Atoi(value)
				dimensions = append(dimensions, uint8(n))
			}
			data := make([]uint, dimensions[1])
			for _, value := range strings.Split(parts[1], " ")[1:] {
				n, _ := strconv.Atoi(value)
				indices = append(indices, uint8(n))
			}
			region := region{dimensions[0], data, indices}

			if canFillRegion(region) {
				total++
			}
		}
	}
	println(total)
}

var shapes [][][3]uint

type region struct {
	width   uint8
	data    []uint
	indices []uint8
}

func canFillRegion(r region) bool {
	var n int
	for n = 0; n < len(r.indices); n++ {
		if r.indices[n] != 0 {
			break
		}
	}
	if n == len(r.indices) {
		return true
	}

	r.indices[n]--
	for h := 0; h <= len(r.data)-3; h++ {
		for _, shape := range shapes[n] {
			var invalids [3]uint
			for s := 0; s < 3; s++ {
				invalids[s] = shape[0]&r.data[h] | shape[1]&r.data[h+1] | shape[2]&r.data[h+2]
				shape[0] <<= 1
				shape[1] <<= 1
				shape[2] <<= 1
			}
			shape[0] >>= 3
			shape[1] >>= 3
			shape[2] >>= 3

			mask := uint(7)
			for w, n := 0, 0; w <= int(r.width)-3; w, n = w+1, n+1 {
				if mask&invalids[n] == 0 {
					r.data[h] ^= shape[0] << w & mask
					r.data[h+1] ^= shape[1] << w & mask
					r.data[h+2] ^= shape[2] << w & mask

					if canFillRegion(r) {
						return true
					}

					r.data[h] ^= shape[0] << w & mask
					r.data[h+1] ^= shape[1] << w & mask
					r.data[h+2] ^= shape[2] << w & mask
				}
				mask <<= 1
				if n == 2 {
					n = -1
				}
			}
		}
	}
	r.indices[n]++

	return false
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
