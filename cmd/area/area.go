package area

import (
	"fmt"
	"math"
	"math/rand"
)

type Area struct {
	Id     int
	TL     []int
	BR     []int
	Room   *Area
	Child  *Area
	Path0  [][]int
	Path1  [][]int
	Path2  [][]int
	Path3  [][]int
	NextTo []int // 隣接している側 0,1,2,3 - 上右下左
}

// 隣接する部屋に向けて通路を生成する
func (a *Area) GenPath() {
	// TODO:(u110) fix
	a.Path0 = make([][]int, 0)
	a.Path1 = make([][]int, 0)
	a.Path2 = make([][]int, 0)
	a.Path3 = make([][]int, 0)
	for _, i := range a.NextTo {
		switch i {
		case 0:
			a.GenTopPath()
		case 1:
			a.GenRightPath()
		case 2:
			a.GenBottomPath()
		case 3:
			a.GenLeftPath()
		default:
			return
		}
	}
}

func (a *Area) LinkPath() {
	idx := len(a.NextTo)
	i := a.NextTo[idx-1]
	switch i {
	case 0:
		// childのBottom x標まで伸ばす
		if a.Child == nil || a.Child.Path2 == nil {
			return
		}
		for _, p := range a.Child.Path2 {
			if p[1] == a.Child.BR[1]-1 {
				x := p[0]
				start := int(math.Min(float64(x), float64(a.Path0[0][0])))
				end := int(math.Max(float64(x), float64(a.Path0[0][0])))
				length := end - start + 1
				paths := make([][]int, length)
				i := 0
				for i < length {
					paths[i] = []int{start + i, a.TL[1]}
					i++
				}
				a.Path0 = append(a.Path0, paths...)
				return
			}
		}
	case 1:
		// childのLeft x標まで伸ばす
		if a.Child == nil || a.Child.Path3 == nil {
			return
		}
		fmt.Println(a.Child.Path3, a.Child.TL)
		for _, p := range a.Child.Path3 {
			if p[0] == a.Child.TL[0] {
				y := p[1]
				start := int(math.Min(float64(y), float64(a.Path1[0][1])))
				end := int(math.Max(float64(y), float64(a.Path1[0][1])))
				length := end - start + 1
				paths := make([][]int, length)
				i := 0
				for i < length {
					paths[i] = []int{a.BR[0] - 1, start + i}
					i++
				}
				fmt.Println("here", paths)
				a.Path1 = append(a.Path1, paths...)
				return
			}
		}
	case 2:
		// childのTop x標まで伸ばす
		if a.Child == nil || a.Child.Path0 == nil {
			return
		}
		fmt.Println(a.Child.Path0, a.Child.TL)
		for _, p := range a.Child.Path0 {
			if p[1] == a.Child.TL[1] {
				x := p[0]
				start := int(math.Min(float64(x), float64(a.Path2[0][0])))
				end := int(math.Max(float64(x), float64(a.Path2[0][0])))
				length := end - start + 1
				paths := make([][]int, length)
				i := 0
				for i < length {
					paths[i] = []int{start + i, a.BR[1] - 1}
					i++
				}
				fmt.Println("here", paths)
				a.Path2 = append(a.Path2, paths...)
				return
			}
		}
	case 3:
		// childのRight x標まで伸ばす
		if a.Child == nil || a.Child.Path1 == nil {
			return
		}
		fmt.Println(a.Child.Path1, a.Child.BR)
		for _, p := range a.Child.Path1 {
			if p[0] == a.Child.BR[0]-1 {
				y := p[1]
				start := int(math.Min(float64(y), float64(a.Path3[0][1])))
				end := int(math.Max(float64(y), float64(a.Path3[0][1])))
				length := end - start + 1
				paths := make([][]int, length)
				i := 0
				for i < length {
					paths[i] = []int{a.TL[0], start + i}
					i++
				}
				fmt.Println("here", paths)
				a.Path3 = append(a.Path3, paths...)
				return
			}
		}
	default:
		return
	}
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
	a.Path0 = append(a.Path0, paths...)
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
	a.Path2 = append(a.Path2, paths...)
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
	a.Path1 = append(a.Path1, paths...)
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
	a.Path3 = append(a.Path3, paths...)
}

func (a *Area) InPath(x int, y int) bool {
	for _, p := range a.Path0 {
		if x == p[0] && y == p[1] {
			return true
		}
	}
	for _, p := range a.Path1 {
		if x == p[0] && y == p[1] {
			return true
		}
	}
	for _, p := range a.Path2 {
		if x == p[0] && y == p[1] {
			return true
		}
	}
	for _, p := range a.Path3 {
		if x == p[0] && y == p[1] {
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
	// mark := fmt.Sprint(a.Id)
	mark := fmt.Sprint(a.Id)
	if a.InPath(i, j) {
		mark = "_"
	}
	if a.InRange(i, j) {
		if a.IsRoom(i, j) {
			mark = " "
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

var UPTO_AREA float64 = 10

// 水平分割
func (a *Area) SepV() {
	sepX := RandFilterIntn(a.W(), UPTO_AREA)
	subA := Area{Id: a.Id + 1, TL: a.TL, BR: []int{a.BR[0] - sepX, a.BR[1]}}
	subB := Area{Id: a.Id + 1, TL: []int{a.TL[0] + (a.W() - sepX), a.TL[1]}, BR: a.BR}
	if subA.W() > subB.W() {
		// 大きい方を子とする
		a.Child = &subA
		a.TL, a.BR = subB.TL, subB.BR
		a.Child.NextTo = append(a.Child.NextTo, 1)
		a.NextTo = append(a.NextTo, 3)
	} else {
		a.Child = &subB
		a.TL, a.BR = subA.TL, subA.BR
		a.Child.NextTo = append(a.Child.NextTo, 3)
		a.NextTo = append(a.NextTo, 1)
	}
}

// 垂直分割
func (a *Area) SepH() {
	sepY := RandFilterIntn(a.H(), UPTO_AREA)
	subA := Area{Id: a.Id + 1, TL: a.TL, BR: []int{a.BR[0], a.BR[1] - sepY}}
	subB := Area{Id: a.Id + 1, TL: []int{a.TL[0], (a.H() - sepY) + a.TL[1]}, BR: a.BR}
	if subA.H() > subB.H() {
		// 大きい方を子とする
		a.Child = &subA
		a.TL, a.BR = subB.TL, subB.BR
		a.Child.NextTo = append(a.Child.NextTo, 2)
		a.NextTo = append(a.NextTo, 0)
	} else {
		a.Child = &subB
		a.TL, a.BR = subA.TL, subA.BR
		a.Child.NextTo = append(a.Child.NextTo, 0)
		a.NextTo = append(a.NextTo, 2)
	}
}
