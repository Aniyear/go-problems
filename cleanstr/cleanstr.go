package main

import (
  "os"
  "github.com/01-edu/z01"
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

func Print(s string) {
  for _, char := range s{
    z01.PrintRune(char)
  }
}

func main() {
  if len(os.Args) != 2 {
    Print("\n")
    return
  }
  arg := os.Args[1]
  // if arg == "" {
  //   Print("\n")
  //   return
  // }
  
  res := ""


  for _, str := range Split(arg) {
    res += str + " "
  }
  
   if res == "" {
    Print("\n")
    return
  }

  Print(res[:len(res)-1]+"\n")
}