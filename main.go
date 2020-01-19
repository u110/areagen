package main

import (
	"fmt"
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
	sepX := a.W() / 2
	area := Area{Id: 0, TL: []int{0, 0}, BR: []int{sepX, a.H()}}
	sub := Area{Id: 1, TL: []int{sepX, 0}, BR: []int{a.W(), a.H()}}
	a.Child = &sub
	a.TL, a.BR = area.TL, area.BR
}

func main() {
	fmt.Println("1. create area")
	x, y := 30, 10
	area := Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}
	area.ShowRange(x, y)
	fmt.Println("2. sep")
	area.Sep()
	area.ShowRange(x, y)
}
