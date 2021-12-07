package main

import (
	"fmt"
	"math"
)

func main() {
	ant_dist := 1.0
	santa_dist := 20.0
	for ant_dist < santa_dist {

		// (a) Beveg nisse
		santa_dist += 20
		ant_dist *= santa_dist / (santa_dist - 20)

		// (b) Beveg maur
		ant_dist++
	}
	fmt.Println(int(math.Floor(ant_dist / 100)))
}
