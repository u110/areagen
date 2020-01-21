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

func main() {
	rand.Seed(time.Now().UnixNano()) // set random seed
	x, y := 100, 50

	for {
		mainArea := area.Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}
		mainArea.ShowRange(x, y)
		time.Sleep(500 * time.Millisecond)

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
			mainArea.ShowRange(x, y)
			fmt.Println(target.NextTo)
			time.Sleep(500 * time.Millisecond)

			target = target.Child
			areaArr = append(areaArr, target)
			countup++
		}
		mainArea.ShowRange(x, y)
		time.Sleep(500 * time.Millisecond)
	}
}
