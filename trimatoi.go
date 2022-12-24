package piscine

func IsRuneDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func TrimAtoi(s string) int {
	var neg bool = false
	var empty bool = true
	var res int = 0
	for _, v := range s {
		if empty && !neg && v == '-' {
			neg = true
		} else if IsRuneDigit(v) {
			res *= 10
			res += int(v - 48)
			empty = false
		}
	}
	if empty {
		return 0
	} else {
		if neg {
			return -res
		} else {
			return res
		}
	}
}
