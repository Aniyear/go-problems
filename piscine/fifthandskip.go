package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Input\n"
	}

	letters := ""
	for _, ch := range str {
		if ch != ' ' {
			letters += string(ch)
		}
	}

	ans := ""
	step := 0
	for _, ch := range letters {
		if step == 5 {
			ans += " "
			step = 0
			continue
		}
		ans += string(ch)
		step++
	}

	return ans + "\n"
}
