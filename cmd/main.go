package main

import (
	"fmt"
	"github.com/u110/areagen/cmd/area"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // set random seed
	fmt.Println("1. create area")
	x, y := 100, 50
	mainArea := area.Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}
	mainArea.GenRoom()
	mainArea.ShowRange(x, y)

	countup := 0
	target := &mainArea
	for countup < 5 {
		target.Sep()
		target.GenRoom()
		target.Child.GenRoom()
		mainArea.ShowRange(x, y)
		target = target.Child
		countup++
	}
}
