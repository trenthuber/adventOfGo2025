# Advent of Go 2025

Completing [this year's Advent of Code](https://adventofcode.com/2025/) to the
best of my ability entirely in [Go](https://go.dev/doc/). Day 10, Part 2 and Day
12, Part 1 remain unable to be checked for correctness since my code runs too
slow. Day 12, Part 2 remains unimplemented since I couldn't complete Part 1 for
that day.

## Building

To build any file, just use the `go build` command. The program reads its input
from a file named `input.txt` located in the same directory that you're running
the program in. To specify whether you'd like to process the input according to
the first or second part of that day's challenge, you can either pass a `1` or a
`2` as a command line argument. Each challenge comes with additional smaller
input used for basic testing that can be copied to a file named `test.txt`.
Passing `test` as a command line argument will have the program read input from
`test.txt` instead.

If you find this code at all useful, I've included a template, `template.go`,
that can be used for any other Advent of Code challenges.
