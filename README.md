# Advent of Go 2025

I decided to try [this year's Advent of Code](https://adventofcode.com/2025/)
exclusively using [Go](https://go.dev/doc/). While most of the programs here run
successfully, there were a few I had trouble with. Day 10, Part 2 and Day 12,
Part 1 simply run too slow to generate an answer (although I do think the
algorithm I went with is correct) and Day 12, Part 2 is left unimplemented since
I couldn't complete Part 1 for that day.

## Building

To build any file, just use the `go build` command. The program reads its input
from a file named `input.txt`. To specify whether you'd like to process the
input according to the first or second part of that day's challenge, you can
either pass a `1` or a `2` as a command line argument.

Each challenge comes with additional input used for basic testing that can be
pasted into a file named `test.txt`.  Passing `test` as a command line argument
will have the program read input from `test.txt` instead of `input.txt`.

If you find this code at all useful, I've included a template, `template.go`,
that can be used for any other Advent of Code challenges.
