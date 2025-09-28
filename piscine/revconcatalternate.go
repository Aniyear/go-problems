package piscine

func RevConcatAlternate(slice1,slice2 []int) []int {
  s1 := len(slice1)-1
  s2 := len(slice2)-1

  var result []int = []int{}

  for s1 > s2 {
	result = append(result, slice1[s1])
	s1--
  }

  for s2 > s1 {
	result = append(result, slice2[s2])
	s2--
  }

  for s1 >= 0 && s2 >= 0 {
	result = append(result, slice1[s1])
	result = append(result, slice2[s2])
	s1--
	s2--
  }

  for s1 >= 0 {
	result = append(result, slice1[s1])
	s1--
  }

  for s2 >= 0 {
	result = append(result, slice2[s2])
	s2--
  }

  return result
}