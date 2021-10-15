package main

import (
	"container/list"
	"fmt"
	"os"
)

type Node struct {
	Data   int
	LChild *Node
	RChild *Node
}

type Tree = *Node

type Op struct {
	IsPush bool
	Data   int
}

const (
	PUSH string = "Push"
	POP  string = "Pop"
)

var (
	N       int
	Input   []Op
	NodeArr []*Node
	Imap    map[int]int
)

func main() {
	file, _ := os.Open("./1086 Tree Traversals Again/data")
	fmt.Fscanf(file, "%d", &N)
	for i := 0; i < 2 * N; i++ {
		tmpStr := ""
		tmpInt := 0
		fmt.Fscanf(file, "%s %d", &tmpStr, &tmpInt)
		switch tmpStr {
		case PUSH:
			Input = append(Input, Op{true, tmpInt})
		case POP:
			Input = append(Input, Op{false, -1})
		}
	}

	Imap = make(map[int]int)
	lst := list.New()
	for i, v := range Input {
		if v.IsPush {
			NodeArr = append(NodeArr, &Node{Data: v.Data})
			lst.PushBack(i)
		} else {
			NodeArr = append(NodeArr, nil)
			e := lst.Back()
			lst.Remove(e)
			Imap[e.Value.(int)] = i
		}
	}
	var tree Tree
	for i := 0; i < 2*N; i++ {
		if Input[i].IsPush {
			node := NodeArr[i]
			if i == 0 {
				tree = node
			}
			if i+1 < 2*N {
				node.LChild = NodeArr[i+1]
			}
			v := Imap[i]
			if v+1 < 2*N {
				node.RChild = NodeArr[v+1]
			}
		}
	}

	output := make([]int, 0)
	tree.PostTraversal(func(data int) {
		output = append(output, data)
	})
	for i := range output {
		fmt.Printf("%d", output[i])
		if i != N-1 {
			fmt.Print(" ")
		}
	}
}

func (t Tree) PostTraversal(f func(int)) {
	if t.LChild != nil {
		t.LChild.PostTraversal(f)
	}
	if t.RChild != nil {
		t.RChild.PostTraversal(f)
	}
	f(t.Data)
}
