package main

import (
	"fmt"
	"github.com/u110/areagen/cmd/area"
	"math/rand"
	"time"
)

func ReGenerateRoom(arr []*area.Area) {
	for _, i := range arr {
		i.GenRoom()
	}
}

func generateArea() {
	rand.Seed(time.Now().UnixNano()) // set random seed
	x, y := 100, 50

	for {
		mainArea := area.Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}

		areaArr := []*area.Area{}
		areaArr = append(areaArr, &mainArea)
		target := &mainArea

		countup := 0
		for countup < 6 {
			target.Sep()
			target.GenRoom()
			target.GenPath()
			target.Child.GenRoom()
			target.Child.GenPath()
			target = target.Child
			areaArr = append(areaArr, target)
			countup++
		}
		// mainArea.ShowRange(x, y)
		// time.Sleep(500 * time.Millisecond)
		for _, a := range areaArr {
			a.LinkPath()
		}
		mainArea.ShowRange(x, y)
		time.Sleep(1200 * time.Millisecond)
	}
}

func test() {
	rand.Seed(3)
	fmt.Println("start")
	x, y := 100, 50
	m := area.Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}
	m.SepV()
	m.GenRoom()
	m.GenPath()
	m.Child.GenRoom()
	m.Child.GenPath()
	m.LinkPath()
	m.ShowRange(x, y)
}

func main() {
	generateArea()
	// test()

}
