package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func createTallMapBase(prefix string, startVal int) map[int]string {
	return map[int]string{
		startVal + 9: fmt.Sprintf("%sni", prefix),
		startVal + 8: fmt.Sprintf("%såtte", prefix),
		startVal + 7: fmt.Sprintf("%ssju", prefix),
		startVal + 6: fmt.Sprintf("%sseks", prefix),
		startVal + 5: fmt.Sprintf("%sfem", prefix),
		startVal + 4: fmt.Sprintf("%sfire", prefix),
		startVal + 3: fmt.Sprintf("%stre", prefix),
		startVal + 2: fmt.Sprintf("%sto", prefix),
		startVal + 1: fmt.Sprintf("%sen", prefix),
	}
}

func createTallMap() map[int]string {
	tallMap := map[int]string{}

	// 1 - 9
	for k, v := range createTallMapBase("", 0) {
		tallMap[k] = v
	}

	// 10 - 19
	tallMap[10] = "ti"
	tallMap[11] = "elleve"
	tallMap[12] = "tolv"
	tallMap[13] = "tretten"
	tallMap[14] = "fjorten"
	tallMap[15] = "femten"
	tallMap[16] = "seksten"
	tallMap[17] = "sytten"
	tallMap[18] = "atten"
	tallMap[19] = "nitten"

	// 20 - 29
	tallMap[20] = "tjue"
	for k, v := range createTallMapBase("tjue", 20) {
		tallMap[k] = v
	}

	// 30 - 39
	tallMap[30] = "tretti"
	for k, v := range createTallMapBase("tretti", 30) {
		tallMap[k] = v
	}

	// 40 - 49
	tallMap[40] = "førti"
	for k, v := range createTallMapBase("førti", 40) {
		tallMap[k] = v
	}

	// 50
	tallMap[50] = "femti"

	return tallMap
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func main() {
	tallData := readFile("tall.txt")
	tallDataLen := len(tallData)
	tallMap := createTallMap()

	// Regn ut sum
	dataPos := 0
	sum := 0
	for dataPos < tallDataLen-1 {
		besteTall := -1
		besteParSum := -1
		for tall := 50; tall > 0; tall-- {
			tallStr := tallMap[tall]
			if strings.HasPrefix(tallData[dataPos:], tallStr) {
				nyDataPos := dataPos + len(tallStr)
				if nyDataPos >= tallDataLen-1 {
					dataPos = nyDataPos
					sum += tall
					break
				} else {
					for nesteTall := 50; nesteTall > 0; nesteTall-- {
						nesteTallStr := tallMap[nesteTall]
						if strings.HasPrefix(tallData[nyDataPos:], nesteTallStr) {
							if tall+nesteTall > besteParSum {
								besteParSum = tall + nesteTall
								besteTall = tall
								break
							}
						}
					}
				}
			}
		}
		if besteTall > -1 {
			dataPos += len(tallMap[besteTall])
			sum += besteTall
		}
	}

	fmt.Printf("%v\n", sum)
}
