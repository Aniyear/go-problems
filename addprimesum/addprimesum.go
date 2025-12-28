package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num := Atoi(os.Args[1])
	sum := 0

	if num < 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	for i := num; i >= 2; i-- {
		if isPrime(i) {
			sum += i
		}
	}

	Print(Itoa(sum) + "\n")

}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Itoa(n int) (s string) {
	for n > 0 {
		s = string(n%10+48) + s
		n /= 10
	}
	return
}

func Atoi(s string) (n int) {
	for _, c := range s {
		n = n*10 + int(c-48)
	}
	return
}

func Print(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}
