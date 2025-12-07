package main 

import (
	"os"
	"github.com/01-edu/z01"
)

func Print(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func main() {
	if len(os.Args) != 3 {
		Print("\n")
		return
	}

	str := os.Args[1] + os.Args[2]

	result := Union(str)

	Print(result + "\n")
}

func Union(str string) string {
	result := ""
	seen := make(map[rune]bool)
	for _, char := range str {
		if !seen[char] {
			seen[char] = true
			result += string(char)
		}	
	}
	return result
}