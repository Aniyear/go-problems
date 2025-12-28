package piscine

import (
	"strconv"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	sign := 1

	if dec[0] == '-' {
		sign = -1
		dec = dec[1:]
	}
	num := 0

	if v, ok := isNumber(dec); ok {
		num = v
	} else {
		return map[bool]string{true: "", false: "-"}[sign == 1] + dec + "\n"
	}

	return transfer(num, sign)

}

func isNumber(dec string) (int, bool) {
	n := 0
	for _, ch := range dec {
		if !(ch >= '0' && ch <= '9') && ch != '.' {
			return 0, false
		} else {
			if ch != '.' {
				n = n*10 + int(ch-48)
			}
		}
	}
	return n, true
}

func transfer(num, sign int) string {
	return strconv.Itoa(num*sign) + "\n"
}
