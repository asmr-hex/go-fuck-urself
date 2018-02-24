package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get filename from cmd args
	fn := os.Args[1]

	// open file
	fd, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	// put file into string
	reader := bufio.NewReader(fd)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	// scan through each token
	InterpreteProgram(bytes)

	fmt.Println("")
}

// create empty int variable array
var memory = make([]int, 512)
var ptr = 0

func InterpreteProgram(bytes []byte) {
	loop := []byte{}
	looping := false
	jump2end := false

	for _, ch := range bytes {
		c := string(ch)

		if c != "]" && jump2end {
			continue
		}

		if looping {
			loop = append(loop, ch)
		}

		switch c {
		case ">":
			ptr += 1
		case "<":
			if ptr != 0 {
				ptr -= 1
			}
		case ".":
			fmt.Printf("%s", string(memory[ptr]))
		case "+":
			memory[ptr] += 1
		case "-":
			memory[ptr] -= 1
		case "[":
			if memory[ptr] == 0 {
				jump2end = true
				continue
			}

			looping = true
			loop = append(loop, ch)
		case "]":
			if jump2end {
				jump2end = false
			}

			if memory[ptr] != 0 {
				InterpreteProgram(loop)
			}

			looping = false
		}
	}
}
