package piscine

func CanJump(path []uint) bool {
	if len(path) == 0 {
		return false
	} else if len(path) == 1 {
		return true
	} else {
		return Jump(0, path)
	}
}

func Jump(i int, path []uint) bool {
	if i == len(path)-1 {
		return true
	} else if i >= len(path) {
		return false
	} else if path[i] == 0 {
		return false
	} else {
		return Jump(i+int(path[i]), path)
	}
}
