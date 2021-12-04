package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	dataStr := string(data)
	bestStartPos := -1
	bestLength := -1
	for i := 0; i < len(dataStr)-1; i++ {
		for j := i + 1; j < len(dataStr); j++ {
			dataSub := dataStr[i:j]
			if strings.Count(dataSub, "J") == strings.Count(dataSub, "N") {
				if len(dataSub) > bestLength {
					bestStartPos = i
					bestLength = len(dataSub)
				}
			}
		}
	}
	fmt.Printf("%v, %v\n", bestLength, bestStartPos)
}
