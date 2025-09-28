package piscine

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	if value == 0 {
		return "0\n"
	}

	var digits string = "0123456789ABCDEF"
	var result string = ""
	var falue uint
	negative := false		
	if value < 0 {
		negative = true
		falue = uint(-value)
	} else {
		falue = uint(value)
	}
	for falue > 0 {
		remainder := falue % uint(base)
		result = string(digits[remainder]) + result
		falue = falue / uint(base)
	}
	if negative {
		result = "-" + result
	}
	return result 
}