package piscine

func LastWord(s string) string{
  var words []string
  var str string
  for _, char := range s {
    if char != ' ' {
      str += string(char)
    } else {
      words = append(words, str)
      str = ""
    }
  }

  if str != ""{
    words = append(words, str)
  }
  return words[len(words)-1]+"\n"
}