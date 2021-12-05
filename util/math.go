package util

func Sign(i int) int {
	if i > 0 {
		return 1
	} else if i == 0 {
		return 0
	}
	return -1
}

func Abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
