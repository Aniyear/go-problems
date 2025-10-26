package main

import (
  "os"
  "github.com/01-edu/z01"
)

func Print(s string) {
  for _, ch := range s {
    z01.PrintRune(ch)
  }
}

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

func main() {
  if len(os.Args) != 2 {
    return
  }

  words := Split(os.Args[1:][0])

  if len(words) == 0 {
    Print("\n")
    return
  }

  ans := ""
  for i := len(words)-1; i >= 0; i-- {
    ans += words[i]+" "
  }

  Print(ans[:len(ans)-1]+"\n")
}