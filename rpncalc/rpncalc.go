package main

import (
	"os"
	"strconv"
	"fmt"
)

func Split(s string) []string{
  var words []string
  var str string
  for _, char := range s {
    if char != ' ' {
      str += string(char)
    } else {
      if str != ""{
      words = append(words, str)
      str = ""
      }
    }
  }

  if str != ""{
    words = append(words, str)
  }
  return words
}

func rpncalc(operation []string) (int, bool) {
	var stack []int
	for _, op := range operation {
		switch op {
		case "+":
			if len(stack) < 2 {
				return 0, false
			}
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)	
		case "-":
			if len(stack) < 2 {
				return 0,false
			}
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b-a)
		case "*":
			if len(stack) < 2 {
				return 0, false
			}
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			if len(stack) < 2 {
				return 0,false
			}
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			if a == 0 {
				return 0,false
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, b/a)
		case "%":
			if len(stack) < 2 {
				return 0,false
			}
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			if a == 0 {
				return 0,false
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, b%a)
		default:
			num, err := strconv.Atoi(op)
			if err != nil {
				return 0, false
			}
			stack = append(stack, num)
		}
	}
	if len(stack) != 1 {
		return 0,false
	}
	return stack[0],true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	arg := os.Args[1]
	operation := Split(arg)
	result, ok := rpncalc(operation)
	if !ok {
		fmt.Println("Error")
		return
	}
	fmt.Println(result)
}