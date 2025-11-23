package main

import (
	"fmt"
	"os"
)

func toRoman(n int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for i := 0; i < len(vals); i++ {
		for n >= vals[i] {
			res += symb[i]
			n -= vals[i]
		}
	}
	return res
}

func expand(roman string) string {
	val := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	out := ""
	for i := 0; i < len(roman); {
		if i+1 < len(roman) && val[roman[i]] < val[roman[i+1]] {
			// вычитание
			out += fmt.Sprintf("(%c-%c)", roman[i+1], roman[i])
			i += 2
		} else {
			// обычный символ
			out += string(roman[i])
			i++
		}
		if i < len(roman) {
			out += "+"
		}
	}
	return out
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: <number>")
		return
	}

	var n int
	fmt.Sscanf(os.Args[1], "%d", &n)
	if n <= 0 || n >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	r := toRoman(n)
	fmt.Println(expand(r))
	fmt.Println(r)
}
