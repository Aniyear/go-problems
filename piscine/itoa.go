package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}	
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}	
	digits := ""
	for n > 0 {
		digit := string(n % 10+48)
		digits = digit + digits
		n /= 10
	}
	if isNegative {
		digits = "-" + digits
	}
	return digits
}