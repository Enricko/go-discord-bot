package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
