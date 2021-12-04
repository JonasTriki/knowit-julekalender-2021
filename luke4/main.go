package main

import (
	"fmt"
)

func tryGoSteps(numSteps int) {
	x := 0
	y := 0
	goNorth := true
	for i := 0; i < numSteps; i++ {
		if goNorth {
			y++
			if y%3 == 0 && y%5 != 0 {
				goNorth = false
			}
		} else {
			x++
			if x%5 == 0 && x%3 != 0 {
				goNorth = true
			}
		}
	}
	fmt.Printf("%v,%v\n", x, y)
}

func main() {
	tryGoSteps(10079)
	tryGoSteps(100079)
	tryGoSteps(1000079)
	tryGoSteps(10000079)

	/*
		Fra resultatene ser man et møster. Legger så til nødvendig antall 6-ere på x og 3-ere på y.
		For eksempel for numSteps=10000079:
		6666716,3333363 => 66666666666666666716,33333333333333333363 (la til til 13x 6 på x og 13x 6 på y)
	*/
}
