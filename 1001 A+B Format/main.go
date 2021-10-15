package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	A, B int
	Res  []byte
)

func main() {
	file, _ := os.Open("./1001 A+B Format/data")
	fmt.Fscanf(file, "%d %d", &A, &B)
	sum := A + B

	if sum < 0 {
		Res = append(Res, '-')
		sum = -sum
	}

	sumBytes := []byte(strconv.Itoa(sum))

	l := len(sumBytes)

	cnt := (l - 1) / 3
	offset := l - cnt*3

	for i := 0; i < l; i++ {
		if offset != 0 {
			offset--
		} else {
			Res = append(Res, ',')
			offset = 2
		}
		Res = append(Res, sumBytes[i])
	}

	fmt.Printf("%s", string(Res))
}
