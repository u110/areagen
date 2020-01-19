package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Area struct {
	Id    int
	TL    []int
	BR    []int
	Child *Area
}

func (a *Area) InRange(x int, y int) bool {
	return x >= a.TL[0] && y >= a.TL[1] && x < a.BR[0] && y < a.BR[1]
}

func (a *Area) W() int {
	return a.BR[0] - a.TL[0]
}
func (a *Area) H() int {
	return a.BR[1] - a.TL[1]
}

func (a *Area) Show(i int, j int) error {
	if a.InRange(i, j) {
		fmt.Print(a.Id)
		return nil
	}
	if a.Child != nil {
		return a.Child.Show(i, j)
	}
	return nil
}

func (a *Area) ShowRange(w int, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			a.Show(i, j)
		}
		fmt.Println()
	}
}

func (a *Area) Sep() {
	sepX := rand.Intn(a.W())
	fmt.Println("sepX:", sepX)
	subA := Area{Id: a.Id + 1, TL: []int{0, 0}, BR: []int{sepX, a.H()}}
	subB := Area{Id: a.Id + 1, TL: []int{sepX, 0}, BR: []int{a.W(), a.H()}}
	if sepX < a.W()/2 {
		// 大きい方を親区間とする
		a.Child = &subA
		a.TL, a.BR = subB.TL, subB.BR
	} else {
		a.Child = &subB
		a.TL, a.BR = subA.TL, subA.BR
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // set random seed
	fmt.Println("1. create area")
	x, y := 50, 10
	area := Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}
	area.ShowRange(x, y)
	fmt.Println("2. sep")
	area.Sep()
	area.ShowRange(x, y)
}
