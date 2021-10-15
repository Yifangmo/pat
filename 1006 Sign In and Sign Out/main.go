package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Person struct {
	Id          string
	SignInTime  int
	SignOutTime int
}

type Person1 struct {
	Id          string
	SignInTime  string
	SignOutTime string
}

var (
	M       int
	Persons []Person
	Person1s []Person1
)

func main() {
	file, _ := os.Open("./1006 Sign In and Sign Out/data")
	fmt.Fscanf(file, "%d", &M)
	for i := 0; i < M; i++ {
		var p, i, o string
		fmt.Fscanf(file, "%s %s %s", &p, &i, &o)
		var iH, iM, iS, iRes int
		fmt.Sscanf(i, "%02d:%02d:%02d", &iH, &iM, &iS)
		iRes = iH*60*60 + iM*60 + iS
		var oH, oM, oS, oRes int
		fmt.Sscanf(o, "%02d:%02d:%02d", &oH, &oM, &oS)
		oRes = oH*60*60 + oM*60 + oS
		Persons = append(Persons, Person{p, iRes, oRes})
		Person1s = append(Person1s, Person1{p, i, o})
	}
	var min, max, minI, maxI int
	min = math.MaxInt64
	max = -1
	for i, p := range Persons {
		if p.SignInTime < min {
			min = p.SignInTime
			minI = i
		}
		if p.SignOutTime > max {
			max = p.SignOutTime
			maxI = i
		}
	}
	fmt.Printf("%s %s", Persons[minI].Id, Persons[maxI].Id)

	fmt.Println()
	minS := "9"
	maxS := ""

	for i, p := range Person1s {
		if strings.Compare(minS, p.SignInTime) == 1 {
			minS = p.SignInTime
			minI = i
		}
		if strings.Compare(maxS, p.SignOutTime) == -1 {
			maxS = p.SignOutTime
			maxI = i
		}
	}

	fmt.Printf("%s %s", Person1s[minI].Id, Person1s[maxI].Id)
}
