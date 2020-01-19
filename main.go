package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	x, y := 30, 10
	area := make([][]int, y)
	for i := 0; i < y; i++ {
		area[i] = make([]int, x)
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("%d", area[i][j])
		}
		fmt.Println()
	}
}
