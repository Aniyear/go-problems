package main

import (
	"fmt"
	"os"
)

func Split(s string) []string {
	var words []string
	var str string
	for _, char := range s {
		if char != ' ' {
			str += string(char)
		} else {
			if str != "" {
				words = append(words, str)
				str = ""
			}
		}
	}

	if str != "" {
		words = append(words, str)
	}
	return words
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	arg := os.Args[1:]

	if arg[1] == "" {
		return
	}

	if len(arg[0]) < 3 {
		return
	}

	words := Split(arg[1])
	key1, key2 := defineKey(arg[0])
	res := []string{}

	for _, word := range words {
		if contains(word, key1) {
			res = append(res, clearSymbols(word))
		}
		if key2 != "" && contains(word, key2) {
			res = append(res, clearSymbols(word))
		}
	}

	for i, str := range res {
		fmt.Printf("%d: %s\n", i+1, str)
	}
}

func defineKey(s string) (string, string) {
	both := false
	key1, key2 := "", ""
	for _, char := range s[1 : len(s)-1] {
		if char == '|' {
			both = true
			continue
		} else if both == true {
			key2 += string(char)
		} else {
			key1 += string(char)
		}
	}
	return key1, key2
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func clearSymbols(s string) string {
	res := ""
	for _, char := range s {
		if char != ',' {
			res += string(char)
		}
	}
	return res
}
