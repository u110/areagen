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
	fmt.Println("1. create area")
	x, y := 100, 50

	for {
		mainArea := area.Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}
		mainArea.GenRoom()
		// mainArea.ShowRange(x, y)
		areaArr := []*area.Area{}

		areaArr = append(areaArr, &mainArea)
		target := &mainArea

		countup := 0
		for countup < 5 {
			target.Sep()
			target.GenRoom()
			target.Child.GenRoom()
			target = target.Child
			areaArr = append(areaArr, target)
			mainArea.ShowRange(x, y)
			time.Sleep(500 * time.Millisecond)
			countup++
		}
		mainArea.ShowRange(x, y)
		time.Sleep(500 * time.Millisecond)
	}
}
