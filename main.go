package main

import (
	"fmt"
	"math"
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
		fmt.Printf("\x1b[3%dm%d\x1b[0m", a.Id%6+1, a.Id)
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

func RandFilterIntn(num int, upto float64) int {
	x := float64(rand.Intn(num))
	x = math.Min(float64(num)-upto, x)
	x = math.Max(upto, x)
	return int(x)
}

var count int

// 交互の垂直、水平の分割を行う
func (a *Area) Sep() {
	if count%2 == 0 {
		a.SepH()
	} else {
		a.SepV()
	}
	count++
}

// 水平分割
func (a *Area) SepV() {
	sepX := RandFilterIntn(a.W(), 8)
	fmt.Println("sepX:", sepX)
	subA := Area{Id: a.Id + 1, TL: a.TL, BR: []int{a.BR[0] - sepX, a.BR[1]}}
	subB := Area{Id: a.Id + 1, TL: []int{a.TL[0] + (a.W() - sepX), a.TL[1]}, BR: a.BR}
	if subA.W() > subB.W() {
		// 大きい方を子とする
		a.Child = &subA
		a.TL, a.BR = subB.TL, subB.BR
	} else {
		a.Child = &subB
		a.TL, a.BR = subA.TL, subA.BR
	}
}

// 垂直分割
func (a *Area) SepH() {
	sepY := RandFilterIntn(a.H(), 8)
	fmt.Println("sepY:", sepY)
	subA := Area{Id: a.Id + 1, TL: a.TL, BR: []int{a.BR[0], a.BR[1] - sepY}}
	subB := Area{Id: a.Id + 1, TL: []int{a.TL[0], (a.H() - sepY) + a.TL[1]}, BR: a.BR}
	if subA.H() > subB.H() {
		// 大きい方を子とする
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
	x, y := 100, 50
	area := Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}
	area.ShowRange(x, y)

	countup := 0
	baseArea := &area
	for countup < 5 {
		baseArea.Sep()
		area.ShowRange(x, y)
		baseArea = baseArea.Child
		countup++
	}
}
