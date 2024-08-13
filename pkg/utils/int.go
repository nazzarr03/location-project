package utils

import (
	"strconv"
)

func UintToString(i uint) string {
	return strconv.Itoa(int(i))
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
