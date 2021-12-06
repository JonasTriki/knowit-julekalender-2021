package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Package struct {
	pos int
	len int
}

const maxPackageLen = 20

func any(bools []bool) bool {
	for i := 0; i < len(bools); i++ {
		if bools[i] {
			return true
		}
	}
	return false
}

func all(bools []bool) bool {
	for i := 0; i < len(bools); i++ {
		if !bools[i] {
			return false
		}
	}
	return true
}

func main() {

	// Les og parse data
	dataBytes, _ := ioutil.ReadFile("pakker.txt")
	dataRaw := strings.Split(string(dataBytes), "\n")
	packageLines := make([]*Package, 0)
	for i := 0; i < len(dataRaw); i++ {
		line := dataRaw[i]
		if len(line) == 0 {
			continue
		}
		lineSplit := strings.Split(dataRaw[i], ",")
		pPos, _ := strconv.Atoi(lineSplit[0])
		pLen, _ := strconv.Atoi(lineSplit[1])
		packageLines = append(packageLines, &Package{pos: pPos, len: pLen})
	}
	numPackages := len(packageLines)

	// Init pakke-matrise
	packages := make([][]bool, numPackages)
	for i := 0; i < numPackages; i++ {
		packages[i] = make([]bool, maxPackageLen)
	}

	// Stable pakker
	numFallenPacks := 0
	for i := 0; i < numPackages; i++ {
		pkg := packageLines[i]
		pkgPos := pkg.pos
		pkgLen := pkg.len
		pkgHalfLen := int(math.Floor(float64(pkgLen) / 2))
		middleBelowStart := 0
		middleBelowEnd := 0
		if pkgLen%2 == 0 {
			middleBelowStart = pkgHalfLen - 1
			middleBelowEnd = middleBelowStart + 2
		} else {
			middleBelowStart = pkgHalfLen
			middleBelowEnd = middleBelowStart + 1
		}
		pkgMinSupportLen := int(math.Ceil(float64(pkgLen) / 2))

		// Søk nedover fra toppen for å finne ledig plass i pakkematrisen
		curPkgRow := numPackages
		hasMcFallen := false
		for nextRow := numPackages - 1; nextRow >= 0; nextRow-- {
			pkgRow := packages[nextRow][pkgPos : pkgPos+pkgLen]

			if any(pkgRow) {

				// For at pakken skal stoppe må en av følgende være sann:
				// 1. Har støtte under midten av pakken
				// 2. Er støttet opp på begge sider (har minst 1 pakke under på hver halvside)
				// Hvis ikke, vil pakken falle av og forsvinne på magiskt vis.
				middlePart := pkgRow[middleBelowStart:middleBelowEnd]
				leftHalf := pkgRow[:pkgMinSupportLen]
				rightHalf := pkgRow[pkgMinSupportLen:]
				hasMiddleSupport := all(middlePart)
				hasSupportUnder := any(leftHalf) && any(rightHalf)
				if hasMiddleSupport || hasSupportUnder {
					break
				} else {

					// Pakken faller og forsvinner
					numFallenPacks++
					hasMcFallen = true // Relevant: https://youtu.be/lWKd5xquliU
					break
				}
			} else {

				// Hele raden under er ledig => vi flytter oss nedover ett hakk.
				curPkgRow = nextRow
			}
		}
		if hasMcFallen {
			continue
		}

		// Plasser pakke dersom vi har lov
		if curPkgRow <= numPackages-1 {
			for i := pkgPos; i < pkgPos+pkgLen; i++ {
				packages[curPkgRow][i] = true
			}
		}
	}

	fmt.Println(numFallenPacks)
}
