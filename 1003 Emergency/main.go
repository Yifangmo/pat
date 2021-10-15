package main

import (
	"fmt"
	"math"
	"os"
)

type City struct {
	Index int
	Hands int
}

var (
	N, M, Src, Dst int
	Matrix         [][]int
	Cities         []City
	Visited        []int
	MinWeight      []int
	MaxHands       []int
	PathCount      []int
)

const (
	UNREACHABLE int = -1
	MAXN        int = 500
)

func main() {
	file, _ := os.Open("./1003 Emergency/data")
	fmt.Fscanf(file, "%d %d %d %d", &N, &M, &Src, &Dst)
	for i := 0; i < N; i++ {
		s := make([]int, 0)
		for j := 0; j < N; j++ {
			if i == j {
				s = append(s, 0)
			} else {
				s = append(s, UNREACHABLE)
			}
		}
		Matrix = append(Matrix, s)
		c := City{Index: i}
		fmt.Fscanf(file, "%d", &c.Hands)
		Cities = append(Cities, c)
		MinWeight = append(MinWeight, math.MaxInt64)
	}
	for i := 0; i < M; i++ {
		j := 0
		k := 0
		v := 0
		fmt.Fscanf(file, "%d %d %d", &j, &k, &v)
		Matrix[j][k] = v
		Matrix[k][j] = v
	}
	Visited = make([]int, N)
	MaxHands = make([]int, N)
	PathCount = make([]int, N)

	MaxHands[Src] = Cities[Src].Hands
	MinWeight[Src] = 0
	currentNode := Src
	for {
		if currentNode == -1 {
			break
		}
		Visited[currentNode] = 1
		for i, v := range Matrix[currentNode] {
			if v > 0 && Visited[i] == 0 {
				if v+MinWeight[currentNode] < MinWeight[i] {
					MinWeight[i] = v + MinWeight[currentNode]
					MaxHands[i] = Cities[i].Hands + MaxHands[currentNode]
					PathCount[i] = 1
				} else if v+MinWeight[currentNode] == MinWeight[i] {
					PathCount[i]++
					if MaxHands[i] < Cities[i].Hands+MaxHands[currentNode] {
						MaxHands[i] = Cities[i].Hands + MaxHands[currentNode]
					}
				}
			}
		}
		minW := math.MaxInt64
		nextNode := -1
		for i := range MinWeight {
			if Visited[i] == 0 {
				if MinWeight[i] < minW {
					minW = MinWeight[i]
					nextNode = i
				} else if MinWeight[i] == minW && nextNode == Dst {
					nextNode = i
				}
			}
		}
		fmt.Println(MinWeight)
		fmt.Println(currentNode)
		currentNode = nextNode
	}

	fmt.Printf("%d %d", PathCount[Dst], MaxHands[Dst])
}
