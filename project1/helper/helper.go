package helper

import (
	"fmt"
)

func Arr(x, y int) [][]int {
	twoD := make([][]int, x)
	for i := 0; i < x; i++ {
		twoD[i] = make([]int, x)
		for j := 0; j < x; j++ {
			twoD[i][j] = y
		}
	}
	return twoD
}

func Arrprint(s [][]int) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func Change(s [][]int, f int) {
	for i := 0; i < len(s); i++ {
		s[i][i] = f
		s[i][len(s)-1-i] = f
	}
}
