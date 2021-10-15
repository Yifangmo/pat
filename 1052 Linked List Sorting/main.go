package main

import (
	"fmt"
	"os"
)

type Node struct {
	Addr string
	Data int
	Next *Node
}

type MapV struct {
	Key      int
	NextAddr string
	Index    int
}

type LinkedList = *Node

var (
	HeadAddr string
	Map      map[string]*MapV
	N        int
	Nodes    []Node
	Link     LinkedList
)

func main() {
	file, _ := os.Open("./1052 Linked List Sorting/data")
	fmt.Fscanf(file, "%d %s", &N, &HeadAddr)
	Map = make(map[string]*MapV)
	for i := 0; i < N; i++ {
		var addr, nextAddr string
		var key int
		fmt.Fscanf(file, "%s %d %s", &addr, &key, &nextAddr)
		Map[addr] = &MapV{key, nextAddr, i}
		Nodes = append(Nodes, Node{addr, key, nil})
	}
	v := Map[HeadAddr]
	Link = &Nodes[v.Index]
	for {
		if v == nil {
			break
		}
		nextV := Map[v.NextAddr]
		if nextV != nil {
			Nodes[v.Index].Next = &Nodes[nextV.Index]
		}
		v = nextV
	}

	// Link.Print()
	var tail *Node
	for {
		node := Link
		if node == tail {
			break
		}
		for {
			next := node.Next
			if next == tail {
				tail = node
				break
			}
			if node.Data > next.Data {
				tmp := node.Data
				node.Data = next.Data
				next.Data = tmp
				tmpAddr := node.Addr
				node.Addr = next.Addr
				next.Addr = tmpAddr
			}
			node = next
		}
	}
	node := Link
	fmt.Printf("%d %s\n", N, Link.Addr)
	for {
		if node == nil {
			break
		}
		nextAddr := "-1"
		if node.Next != nil {
			nextAddr = node.Next.Addr
		}
		fmt.Printf("%s %d %s", node.Addr, node.Data, nextAddr)
		if node.Next != nil {
			fmt.Println()
		}
		node = node.Next
	}
}

func (l LinkedList) Print() {
	node := l
	for {
		if node == nil {
			break
		}
		next := node.Next
		nextAddr := "-1"
		if next != nil {
			nextAddr = next.Addr
		}
		fmt.Printf("(%s:%d:%s)", node.Addr, node.Data, nextAddr)
		if next != nil {
			fmt.Print("->")
		}
		node = next
	}
}
