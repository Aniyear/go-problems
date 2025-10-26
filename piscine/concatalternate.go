package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	l1 := len(slice1) - 1
	l2 := len(slice2) - 1
	ans := []int{}

	if l1 > l2 {
		i := 0
		for l2 >= 0 {
			ans = append(ans, slice1[i])
			ans = append(ans, slice2[i])
			i++
			l2--
		}
		ans = append(ans, slice1[len(slice2):]...)
	} else if l2 > l1 {
		i := 0
		for l1 >= 0 {
			ans = append(ans, slice2[i])
			ans = append(ans, slice1[i])
			i++
			l1--
		}
		ans = append(ans, slice2[len(slice1):]...)
	} else {
		i := 0
		for l2 >= 0 {
			ans = append(ans, slice1[i])
			ans = append(ans, slice2[i])
			i++
			l2--
		}
	}

	if len(ans) == 0 {
		return nil
	}

	return ans
}
