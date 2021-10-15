package main

import (
	"container/list"
	"fmt"
	"os"
)

const ROOTID string = "01"

var (
	N int
	M int
	Input map[string][]string
)

func main() {
	Input = make(map[string][]string)
	file, _ := os.Open("./1004 Counting Leaves/data")
	fmt.Fscanf(file, "%d %d", &N, &M)
	for i := 0; i < M; i++ {
		pID := ""
		cNum := 0
		fmt.Fscanf(file, "%s %d", &pID, &cNum)
		for i := 0; i < cNum; i++ {
			cID := ""
			fmt.Fscanf(file, "%s", &cID)
			Input[pID] = append(Input[pID], cID)
		}
	}
	// fmt.Println(Input)
	lst := list.New()
	lst.PushBack(ROOTID)
	levelSize := lst.Len()
	leafNodeCnt := 0
	output := make([]int, 0)
	for {
		if lst.Len() == 0 {
			break
		}
		e := lst.Front()
		lst.Remove(e)
		levelSize--
		cIDs, ok := Input[e.Value.(string)]
		if !ok {
			leafNodeCnt++
		}
		for _, v := range cIDs {
			lst.PushBack(v)
		}
		if levelSize == 0 {
			output = append(output, leafNodeCnt)
			levelSize = lst.Len()
			leafNodeCnt = 0
		}
	}
	for i, v := range output {
		fmt.Printf("%d", v)
		if i != len(output)-1 {
			fmt.Print(" ")
		}
	}
}
