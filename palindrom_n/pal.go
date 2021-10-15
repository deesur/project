package main

import (
	"fmt"
)

func pal(st string) bool {
	result1 := []byte{}
	result2 := []byte{}
	for i := len(st) - 1; i >= 0; i-- {
		if st[i] >= 65 && st[i] <= 90 || st[i] >= 97 && st[i] <= 122 {
			result1 = append(result1, st[i])
		}
	}
	for i := 0; i <= len(st)-1; i++ {
		if st[i] >= 65 && st[i] <= 90 || st[i] >= 97 && st[i] <= 122 {
			result2 = append(result2, st[i])
		}
	}

	return string(result1) == string(result2)
}

func main() {
	fmt.Println(pal("qwertytrwq"))
}
