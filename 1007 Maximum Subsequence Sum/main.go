package main

import (
	"fmt"
	"os"
)

var (
	k   int
	arr []int
)

func main() {
	file, _ := os.Open("1007 Maximum Subsequence Sum/data")
	fmt.Fscanf(file, "%d", &k)
	for i := 0; i < k; i++ {
		var tmp int
		fmt.Fscanf(file, "%d", &tmp)
		arr = append(arr, tmp)
	}
	var b, e, i, j, sum, tmpSum int
	for {
		if i >= k || arr[i] >= 0 {
			break
		}
		i++
	}
	j = i
	if i == k {
		fmt.Printf("%d %d %d", 0, 0, k)
		return
	}

	for {
		if j >= k {
			break
		}
		tmpSum += arr[j]
		if tmpSum > sum {
			b = i
			e = j
			sum = tmpSum
			fmt.Println(sum)
		}
		if tmpSum < 0 {
			i = j + 1
			tmpSum = 0
		}
		j++
	}

	fmt.Printf("%d %d %d", sum, b, e)
}
