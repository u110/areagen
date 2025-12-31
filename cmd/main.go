package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/u110/areagen/cmd/area"
)

// Config はダンジョン生成のパラメータを保持する
type Config struct {
	Loop     bool
	Interval int
	Width    int
	Height   int
	Depth    int
}

func ReGenerateRoom(arr []*area.Area) {
	for _, i := range arr {
		i.GenRoom()
	}
}

func generateArea(cfg Config) {
	rand.Seed(time.Now().UnixNano()) // set random seed
	x, y := cfg.Width, cfg.Height

	for {
		mainArea := area.Area{Id: 0, TL: []int{0, 0}, BR: []int{x, y}}

		areaArr := []*area.Area{}
		areaArr = append(areaArr, &mainArea)
		target := &mainArea

		countup := 0
		for countup < cfg.Depth {
			target.Sep()
			target.GenRoom()
			target.GenPath()
			target.Child.GenRoom()
			target.Child.GenPath()
			target = target.Child
			areaArr = append(areaArr, target)
			countup++
		}
		for _, a := range areaArr {
			a.LinkPath()
		}
		mainArea.ShowRange(x, y)

		if !cfg.Loop {
			break
		}
		time.Sleep(time.Duration(cfg.Interval) * time.Millisecond)
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
	cfg := Config{}
	flag.BoolVar(&cfg.Loop, "loop", false, "Enable continuous regeneration")
	flag.IntVar(&cfg.Interval, "interval", 1200, "Regeneration interval in milliseconds")
	flag.IntVar(&cfg.Width, "width", 100, "Map width")
	flag.IntVar(&cfg.Height, "height", 50, "Map height")
	flag.IntVar(&cfg.Depth, "depth", 6, "BSP split depth")
	flag.Parse()

	generateArea(cfg)
}
