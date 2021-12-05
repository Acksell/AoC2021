package util

import "strconv"

func StringsToInt(strs []string) (ints []int) {
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return
}
