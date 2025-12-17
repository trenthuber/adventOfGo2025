// https://adventofcode.com/2025/day/9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	getMax(scanner, func(r rect, tiles []tile) bool {
		return true
	})
}

func part2(scanner *bufio.Scanner) {
	getMax(scanner, func(r rect, tiles []tile) bool {
		for i := range tiles {
			bitcode := func(r rect, p tile) (b byte) {
				if p.x <= r.min.x {
					b |= 0b1000
				} else if p.x >= r.max.x {
					b |= 0b0100
				}
				if p.y <= r.min.y {
					b |= 0b0010
				} else if p.y >= r.max.y {
					b |= 0b0001
				}
				return
			}
			if bitcode(r, tiles[i])&bitcode(r, tiles[(i+1)%len(tiles)]) == 0 {
				return false
			}
		}
		return true
	})
}

type tile struct {
	x, y int
}

type rect struct {
	min, max tile
}

func getMax(scanner *bufio.Scanner, valid func(r rect, tiles []tile) bool) {
	var tiles []tile
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		var t tile
		t.x, _ = strconv.Atoi(coords[0])
		t.y, _ = strconv.Atoi(coords[1])
		tiles = append(tiles, t)
	}

	max := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			r := rect{tiles[i], tiles[j]}
			if r.min.x > r.max.x {
				r.min.x, r.max.x = r.max.x, r.min.x
			}
			if r.min.y > r.max.y {
				r.min.y, r.max.y = r.max.y, r.min.y
			}
			area := (r.max.x - r.min.x + 1) * (r.max.y - r.min.y + 1)
			if area > max && valid(r, tiles) {
				max = area
			}
		}
	}
	println(max)
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
