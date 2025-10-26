package main

import (
	// "strconv"
	"fmt"
	"os"
)

func isBalanced(s string) bool {
	stack := make([]rune, 0, len(s))

	match := map[rune]rune{
		')':'(',
		']':'[',
		'}':'{',
	}

	for _, ch := range s {
		switch ch {
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != match[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		case '(', '[', '{':
			stack = append(stack, ch)
		default:
			continue
		}
	}

	return len(stack) == 0
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println()
		return 
	}

	for _, arg := range os.Args[1:] {
		if isBalanced(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}