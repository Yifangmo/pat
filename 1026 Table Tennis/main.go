package main

import (
	"fmt"
	"os"
	"sort"
)

type Player struct {
	ArrivingTime       string
	ATInSeconds        int
	ReqSeconds         int
	ServeTimeInSeconds int
	IsVIP              bool
	ServedTableID      int
}

type Table struct {
	ID         int
	IsVIP      bool
	ServeCount int
}

var (
	N, K, M    int
	Players    []Player
	Tables     []Table
	Serving    []bool
	VIPTableID []int
	FinishTime []int
)

const LIMIT = 60 * 60 * 13

func ParseTime(sec int) string {
	if sec > LIMIT {
		return ""
	}
	hours := sec / 60 / 60
	minutes := sec / 60
	seconds := sec % 60
	res := fmt.Sprintf("%02d:%02d:%02d", hours+8, minutes, seconds)
	return res
}

func main() {
	file, _ := os.Open("./1026 Table Tennis/data")
	fmt.Fscanf(file, "%d", &N)
	for i := 0; i < N; i++ {
		var str string
		var i1, i2 int
		fmt.Fscanf(file, "%s %d %d", &str, &i1, &i2)
		isVIP := false
		if i2 == 1 {
			isVIP = true
		}
		var h, m, s int
		fmt.Sscanf(str, "%02d:%02d:%02d", &h, &m, &s)
		s += (h-8)*3600 + m*60
		Players = append(Players, Player{str, s, i1*60, -1, isVIP, -1})
	}
	fmt.Fscanf(file, "%d %d", &K, &M)
	for i := 0; i < K; i++ {
		Tables = append(Tables, Table{i, false, 0})
		Serving = make([]bool, K)
	}
	FinishTime = make([]int, K)
	for i := 0; i < M; i++ {
		var tmp int
		fmt.Fscanf(file, "%d", &tmp)
		Tables[i].IsVIP = true
		VIPTableID = append(VIPTableID, tmp)
	}

	sort.Slice(Players, func(i, j int) bool {
		return Players[i].ATInSeconds < Players[j].ATInSeconds
	})

	// fmt.Println(Players)

	// for i, p := range Players {
	// 	// if p.ATInSeconds >= 
	// }
	
}
