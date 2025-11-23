package main

import (
	"os"
	"fmt"
)

func Atoi(s string) int {
	if s == "" {
		return 4000
	}
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 4000
		}
		n = n*10 + int(ch-'0')
		if n > 4000 {
			return 4000
		}
	}
	return n
}

func toRoman(arg string) string {
	n := Atoi(arg)
	if n <= 0 || n >= 4000 {
		return "ERROR: cannot convert to roman digit"
	}

	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	r := thousands[(n/1000)%10] +
		hundreds[(n/100)%10] +
		tens[(n/10)%10] +
		ones[n%10]

	return r
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
		return
	}
	arg := os.Args[1]
	fmt.Println(toRoman(arg))
}