package main

import "fmt"

func main() {
out:
	for k := 1; k <= 2000; k++ {
		for j := 1; j <= 2000; j++ {
			for i := 1; i <= 2000; i++ {
				a := 2424154637*k + 1854803357
				b := 2807727397*j + 2787141611
				c := 2537380333*i + 1159251923
				if a == b && b == c {
					fmt.Println(a)
					break out
				}
			}
		}
	}
}
