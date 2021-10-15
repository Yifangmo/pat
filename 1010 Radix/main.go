package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	A, B, ReqStr string
	Tag, R1      int
	R2           string = "Impossible"
)

const MAX_RADIX = 36

func main() {
	file, _ := os.Open("./1010 Radix/data")
	fmt.Fscanf(file, "%s %s %d %d", &A, &B, &Tag, &R1)
	input := new(big.Int)
	switch Tag {
	case 1:
		input.SetString(A, R1)
		ReqStr = B
	case 2:
		input.SetString(B, R1)
		ReqStr = A
	}
	byts := []byte(ReqStr)
	minR := 2
	var maxB byte = '1'
	for _, b := range byts {
		if b > maxB {
			maxB = b
			if maxB >= 'a' {
				minR = 11 + int(maxB-'a')
			} else {
				minR = 1 + int(maxB-'0')
			}
		}
	}

	for i := minR; i <= MAX_RADIX; i++ {
		req := new(big.Int)
		req.SetString(ReqStr, i)
		res := req.Cmp(input)
		if res == 0 {
			R2 = strconv.Itoa(i)
		}else if res == 1 {
			break
		}
	}
	fmt.Printf("%s", R2)
}
