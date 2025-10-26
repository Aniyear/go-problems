package piscine

func CamelToSnakeCase(s string) string{
  if s == "" {
    return ""
  } else if !isCamelCase(s) {
    return s
  }
  return "Camel"
}

func isCamelCase(s string) bool{
	for i := 0; i < len(s); i++ {
		isLower := s[i] >= 'a' && s[i] <= 'z'
		isUpper := s[i] >= 'A' && s[i] <= 'Z'
		if !isLower && !isUpper {
			return false
		}

		if i > 0 {
			if isUpper && s[i-1] >= 'A' && s[i-1] <= 'Z' {
				return false
			}
		}
	}

	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return false
	}
	
	return true
}

