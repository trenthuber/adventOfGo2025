# Advent of Go 2025

Completing [this year's Advent of Code](https://adventofcode.com/2025/) to the
best of my ability entirely in [Go](https://go.dev/doc/). Day 10, Part 2 and Day
12, Part 1 remain unable to be checked for correctness since my code runs too
slow. Day 12, Part 2 remains unimplemented since I couldn't complete Part 1 for
that day.

## Building

To build a specific file, just `go build` it. The resulting executable expects
to take input from a file named `input.txt` located in the same directory as
you're running the executable in. To specify whether you'd like to process that
input according to the first or second part of that day, you have to pass either
a `1` or a `2` on the command line. Additionally, since each day comes with an
example input used for testing, you can optionally pass `test` at the end of the
command line to take input from `test.txt` instead, which is where I would paste
the test input copied directly from the webpage.

If you find this code at all useful, I've included a template, `template.go`,
that can be used for any other Advent of Code challenges.
