package area

import (
	"fmt"
	"math"
	"math/rand"
)

type Area struct {
	Id    int
	TL    []int
	BR    []int
	Room  *Area
	Child *Area
	Paths [][]int
}

// 隣接する部屋に向けて通路を生成する
func (a *Area) GenPath() {
	// TODO:(u110) fix
	a.Paths = make([][]int, 0)
	a.GenBottomPath()
	a.GenTopPath()
	a.GenRightPath()
	a.GenLeftPath()

}

// 上方に通路を生成
func (a *Area) GenTopPath() {
	pathlen := a.Room.TL[1] - a.TL[1]
	x := a.Room.TL[0] + rand.Intn(a.Room.W())

	paths := make([][]int, pathlen)
	countup := 0
	for countup < pathlen {
		paths[countup] = []int{
			x,
			a.TL[1] + countup,
		}
		countup++
	}
	a.Paths = append(a.Paths, paths...)
}

// 下方に通路を生成
func (a *Area) GenBottomPath() {
	pathlen := a.BR[1] - a.Room.BR[1]
	x := a.Room.TL[0] + rand.Intn(a.Room.W())

	paths := make([][]int, pathlen)
	countup := 0
	for countup < pathlen {
		paths[countup] = []int{
			x,
			a.Room.BR[1] + countup,
		}
		countup++
	}
	a.Paths = append(a.Paths, paths...)
}

// 右方に通路を生成
func (a *Area) GenRightPath() {
	pathlen := a.BR[0] - a.Room.BR[0]
	y := a.Room.TL[1] + rand.Intn(a.Room.H())

	paths := make([][]int, pathlen)
	countup := 0
	for countup < pathlen {
		paths[countup] = []int{
			a.Room.BR[0] + countup,
			y,
		}
		countup++
	}
	a.Paths = append(a.Paths, paths...)
}

// 左方に通路を生成
func (a *Area) GenLeftPath() {
	pathlen := a.Room.TL[0] - a.TL[0]
	y := a.Room.TL[1] + rand.Intn(a.Room.H())

	paths := make([][]int, pathlen)
	countup := 0
	for countup < pathlen {
		paths[countup] = []int{
			a.TL[0] + countup,
			y,
		}
		countup++
	}
	a.Paths = append(a.Paths, paths...)
}

func (a *Area) InPath(x int, y int) bool {
	pathlen := len(a.Paths)
	for i := 0; i < pathlen; i++ {
		path := a.Paths[i]
		if x == path[0] && y == path[1] {
			return true
		}
	}
	return false
}

func (a *Area) GenRoom() {
	x1 := RandFilterIntn(a.W()/2, 2)
	y1 := RandFilterIntn(a.H()/2, 2)
	x2 := RandFilterIntn(a.W()/2, 2)
	y2 := RandFilterIntn(a.H()/2, 2)
	a.Room = &Area{
		TL: []int{a.TL[0] + x1, a.TL[1] + y1},
		BR: []int{a.BR[0] - x2, a.BR[1] - y2},
	}
}

func (a *Area) IsRoom(x int, y int) bool {
	room := a.Room
	if room == nil {
		return false
	}
	return x >= room.TL[0] && y >= room.TL[1] && x < room.BR[0] && y < room.BR[1]
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
	mark := fmt.Sprint(a.Id)
	if a.InRange(i, j) {
		if a.InPath(i, j) {
			mark = "_"
		}
		if a.IsRoom(i, j) {
			mark = "."
		}
		fmt.Printf("\x1b[3%dm%s\x1b[0m", a.Id%6+1, mark)
		return nil
	}
	if a.Child != nil {
		return a.Child.Show(i, j)
	}
	return nil
}

func (a *Area) ShowRange(w int, h int) {
	fmt.Print("\033[H\033[2J") // カーソル移動、再描画で erase
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
