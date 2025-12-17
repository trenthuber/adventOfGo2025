// https://adventofcode.com/2025/day/8

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
	_, cables, adj := getData(scanner)

	for _, c := range cables[:numConnections] {
		adj[c.i][c.j], adj[c.j][c.i] = 1, 1
	}

	total := 1
	for i, sizes := 0, getCompSizes(adj); i < 3; i++ {
		total *= sizes[i]
	}
	println(total)
}

func part2(scanner *bufio.Scanner) {
	vecs, cables, adj := getData(scanner)

	c := 0
	for numComps := len(vecs); c < len(cables); {
		for n := 0; n < numComps-1; n, c = n+1, c+1 {
			adj[cables[c].i][cables[c].j], adj[cables[c].j][cables[c].i] = 1, 1
		}
		if numComps = len(getCompSizes(adj)); numComps == 1 {
			break
		}
	}
	println(vecs[cables[c-1].i][0] * vecs[cables[c-1].j][0])
}

type cable struct {
	i, j, d int
}

func getData(scanner *bufio.Scanner) (vecs [][3]int, cables []cable, adj [][]int) {
	for scanner.Scan() {
		var vec [3]int
		for i, num := range strings.Split(scanner.Text(), ",") {
			vec[i], _ = strconv.Atoi(num)
		}
		vecs = append(vecs, vec)
	}

	for i := 0; i < len(vecs); i++ {
		for j := i + 1; j < len(vecs); j++ {
			dx, dy, dz := vecs[j][0]-vecs[i][0], vecs[j][1]-vecs[i][1], vecs[j][2]-vecs[i][2]
			cables = append(cables, cable{i, j, dx*dx + dy*dy + dz*dz})
		}
	}
	slices.SortFunc(cables, func(a, b cable) int {
		return cmp.Compare(a.d, b.d)
	})

	adj = make([][]int, len(vecs))
	for i := range adj {
		adj[i] = make([]int, len(vecs))
	}

	return
}

func getCompSizes(adj [][]int) (sizes []int) {
	for i, checklist, queue := 0, make([]int, len(adj)), []int{}; i < len(adj); i++ {
		if checklist[i] == 1 {
			continue
		}
		checklist[i] = 1
		queue = append(queue, i)
		for n := 0; n < len(queue); n++ {
			for j := 0; j < len(adj); j++ {
				if adj[queue[n]][j] == 1 && checklist[j] == 0 {
					checklist[j] = 1
					queue = append(queue, j)
				}
			}
		}
		sizes = append(sizes, len(queue))
		queue = nil
	}
	slices.SortFunc(sizes, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	return
}

var numConnections int

func main() {
	var part func(*bufio.Scanner)
	filename := "input.txt"
	usage := func() {
		fmt.Printf("usage: %v {1|2} [test]\n", os.Args[0])
		os.Exit(1)
	}
	numConnections = 1000
	switch len(os.Args) {
	case 3:
		if os.Args[2] == "test" {
			filename = "test.txt"
			numConnections = 10
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
