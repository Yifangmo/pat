package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
)

func main() {
	K := 0
	Map := map[int]*big.Float{}
	KeyArr := make([]int, 0)
	file, _ := os.Open("./1002 A+B for Polynomials/data")
	fmt.Fscanf(file, "%d", &K)
	for i := 0; i < K; i++ {
		exp1 := 0
		c1 := ""
		fmt.Fscanf(file, "%d %s", &exp1, &c1)
		f := new(big.Float)
		f.SetString(c1)
		Map[exp1] = f
	}
	fmt.Fscanf(file, "%d", &K)
	for i := 0; i < K; i++ {
		exp2 := 0
		c2 := ""
		fmt.Fscanf(file, "%d %s", &exp2, &c2)
		f := new(big.Float)
		f.SetString(c2)
		v, ok := Map[exp2]
		if ok {
			res := new(big.Float)
			res.Add(f, v)
			if res.String() == "0" {
				delete(Map, exp2)
			} else {
				Map[exp2] = res
			}
		} else {
			Map[exp2] = f
		}
	}
	for k := range Map {
		KeyArr = append(KeyArr, k)
	}

	sort.Slice(KeyArr, func(i, j int) bool {
		return KeyArr[j] < KeyArr[i]
	})

	fmt.Printf("%d", len(Map))

	for _, v := range KeyArr {
		fmt.Printf(" %d %s", v, Map[v].Text('f', 1))
	}
}
