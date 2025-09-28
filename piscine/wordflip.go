package piscine

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}
	  words := Split(str)
	  res := ""
	  for i := len(words)-1; i >= 0; i-- {
	    res += words[i] + " "
	  }
	  if res == "" {
	    return "\n"
	  }	
	  return res[:len(res)-1]+"\n"
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
