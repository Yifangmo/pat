package main

import (
	"fmt"
	"strconv"
)

var Map map[byte]string = map[byte]string{
	'0': "zero",
	'1': "one",
	'2': "two",
	'3': "three",
	'4': "four",
	'5': "five",
	'6': "six",
	'7': "seven",
	'8': "eight",
	'9': "nine",
}

func main() {
	input := ""
	fmt.Scanf("%s", &input)
	digits := []byte(input)
	sum := 0
	for _, b := range digits {
		sum += int(b - '0')
	}
	sumStr := strconv.Itoa(sum)
	sumArr := []byte(sumStr)
	for i, b := range sumArr {
		fmt.Printf("%s", Map[b])
		if i != len(sumArr)-1 {
			fmt.Print(" ")
		}
	}
}
