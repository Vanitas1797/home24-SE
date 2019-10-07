package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	text := read()
	compile(text)
	execute(text)
}

func read() string {
	args := os.Args
	if len(args) != 2 {
		fmt.Print("Usage: ", args[0], " <file name>\n")
		os.Exit(1)
	}
	fileName := args[1]
	file, err := ioutil.ReadFile(fileName)
	text := string(file)
	text = strings.ReplaceAll(text, "\r\n", "")
	if err != nil {
		fmt.Printf("Error reading %s\n", fileName)
		os.Exit(1)
	}
	return text
}

func compile(text string) {
	a := 0
	ptr := 0
	mem := make([]int, 30000)
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '[':
			a++
		case ']':
			a--
		case '<':
			ptr--
		case '>':
			ptr++
		}
	}
	if a != 0 {
		fmt.Println("Syntax error: One or more parantheses are not closed!")
	}
	if ptr < 0 {
		fmt.Println("Syntax error: The program pointer is out of range [-]!")
	} else if ptr > len(mem)-1 {
		fmt.Println("Syntax error: The program pointer is out of range [+]!")
	}
	if a != 0 || ptr < 0 || ptr > len(mem)-1 {
		os.Exit(0)
	}
}

func execute(text string) {
	ptr := 0
	mem := make([]int, 30000)
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '<':
			ptr--
		case '>':
			ptr++
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '[':
			if mem[ptr] == 0 {
				for j := 1; j > 0; i++ {
					if text[i+1] == '[' {
						j++
					} else if text[i+1] == ']' {
						j--
					}
				}
			}
		case ']':
			for j := 1; j > 0; i-- {
				if text[i-1] == '[' {
					j--
				} else if text[i-1] == ']' {
					j++
				}
			}
			i--
		case '.':
			fmt.Print(string(mem[ptr]))
		case ',':
			fmt.Scan(&mem[ptr])
		}
	}
}
