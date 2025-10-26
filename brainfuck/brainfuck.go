package main 

import (
	"os"
	"github.com/01-edu/z01"
)

func buildBracketMap(code string) map[int]int {
	stack := []int{}
	pairs := make(map[int]int)

	for i, c := range code {
		if c == '[' {
			stack = append(stack, i)
		} else if c == ']' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pairs[j] = i
			pairs[i] = j
		}
	}
	return pairs
}

func main() {
	if len(os.Args) <= 1 {
		return
	}

	code := os.Args[1]

	var mem [2048]byte
	ptr := 0
	ip := 0

	brackets := buildBracketMap(code)

	for ip < len(code) {
		switch code[ip] {
		case '>':
			if ptr < len(mem)-1 {
				ptr++
			}
		case '<':
			if ptr > 0 {
				ptr--
			}
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			z01.PrintRune(rune(mem[ptr]))
		case '[':
			if mem[ptr] == 0 {
				ip = brackets[ip]
			}
		case ']':
			if mem[ptr] != 0 {
				ip = brackets[ip]
			}
		default:
		}
		ip++
	}
 

}