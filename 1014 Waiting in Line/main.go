package main

import (
	"container/list"
	"fmt"
	"os"
)

type Customer struct {
	Spend      int
	FinSpend   int
	FinTimeStr string
}

type LstNode struct {
	Spend int
	Index int
}

const (
	LIMIT   int = 540
	INT_MAX     = int(^uint(0) >> 1)
)

func (c *Customer) ParseTime() {
	if c.FinSpend > LIMIT {
		c.FinTimeStr = "Sorry"
		return
	}
	hours := c.FinSpend/60 + 8
	minutes := c.FinSpend % 60
	c.FinTimeStr = fmt.Sprintf("%02d:%02d", hours, minutes)
}

var (
	N, M, K, Q int
	Customers  []Customer
	QueryArr   []int

	CurrentRemain  []int
	CurrentElement []*list.Element
	Lists          []*list.List
)

func main() {
	file, _ := os.Open("./1014 Waiting in Line/data")
	fmt.Fscanf(file, "%d %d %d %d", &N, &M, &K, &Q)
	for i := 0; i < K; i++ {
		tmp := 0
		fmt.Fscanf(file, "%d", &tmp)
		Customers = append(Customers, Customer{Spend: tmp})
	}
	for i := 0; i < Q; i++ {
		tmp := 0
		fmt.Fscanf(file, "%d", &tmp)
		QueryArr = append(QueryArr, tmp)
	}
	for i := 0; i < N; i++ {
		lst := list.New()
		for j := 0; j < M; j++ {
			if j*N+i < K {
				e := lst.PushBack(LstNode{Spend: Customers[j*N+i].Spend, Index: j*N + i})
				if j == 0 {
					CurrentRemain = append(CurrentRemain, Customers[i].Spend)
					CurrentElement = append(CurrentElement, e)
				}
			}
		}
		Lists = append(Lists, lst)
	}

	for i := M * N; i < K; {
		idxs, spend := GetMinIndexes()
		for j := range CurrentRemain {
			CurrentRemain[j] -= spend
		}
		for j, idx := range idxs {
			if i+j < K {
				Lists[idx].PushBack(LstNode{Customers[i+j].Spend, i + j})
			}
			CurrentElement[idx] = CurrentElement[idx].Next()
			spend := 0
			if CurrentElement[idx] != nil {
				spend = CurrentElement[idx].Value.(LstNode).Spend
			}
			CurrentRemain[idx] = spend
		}
		i += len(idxs)
	}

	for _, lst := range Lists {
		sum := 0
		node := lst.Front()
		for {
			if node == nil {
				break
			}
			n := node.Value.(LstNode)
			sum += n.Spend
			Customers[n.Index].FinSpend = sum
			Customers[n.Index].ParseTime()
			node = node.Next()
		}
	}

	for i := 0; i < Q; i++ {
		fmt.Printf("%s", Customers[QueryArr[i]-1].FinTimeStr)
		if i != Q-1 {
			fmt.Println()
		}
	}
}

func GetMinIndexes() (idxs []int, data int) {
	data = INT_MAX
	for i, v := range CurrentRemain {
		if v < data {
			idxs = idxs[0:0]
			idxs = append(idxs, i)
			data = v
		} else if v == data {
			idxs = append(idxs, i)
		}
	}
	// fmt.Println(idxs)
	return
}
