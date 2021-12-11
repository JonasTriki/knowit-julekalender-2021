package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Loc struct {
	x int
	y int
}

func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

func maxHist(row []int) (int, int, int) {
	result := make([]int, 0)
	topVal := 0
	left := 0
	maxArea := 0
	maxLeft := -1
	maxRight := -1
	area := 0
	i := 0
	numCols := len(row)
	for i < numCols {
		if len(result) == 0 || row[result[0]] <= row[i] {
			result = prependInt(result, i)
			i++
			continue
		}

		left = result[0]
		topVal = row[left]
		result = result[1:]
		area = topVal * i

		if len(result) > 0 {
			left = result[0] + 1
			area = topVal * (i - left)
		}

		if area > maxArea {
			maxArea = area
			maxLeft = left
			maxRight = i - 1
		}
	}

	for len(result) > 0 {
		left = result[0]
		topVal = row[left]
		result = result[1:]
		area = topVal * i

		if len(result) > 0 {
			left = result[0] + 1
			area = topVal * (i - left)
		}

		if area > maxArea {
			maxArea = area
			maxLeft = left
			maxRight = numCols - 1
		}
	}
	return maxArea, maxLeft, maxRight
}

func findLargestRect(grid [][]int, targetNum int) (Loc, Loc) {
	top := 0
	bottom := 0
	result, left, right := maxHist(grid[0])
	for j := 1; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if grid[j][i] == 1 {
				grid[j][i] += grid[j-1][i]
			}
		}

		tmpResult, tmpLeft, tmpRight := maxHist(grid[j])

		if tmpResult > result {
			left = tmpLeft
			right = tmpRight
			bottom = j
			result = tmpResult
			top = bottom - (result / (right - left + 1)) + 1
		}
	}

	return Loc{x: left, y: top}, Loc{x: right, y: bottom}
}

func solve(filename string, numLocs int, gridSize int) {
	dataRaw, _ := ioutil.ReadFile(filename)
	dataLines := strings.Split(string(dataRaw), "\n")

	// Parse locations
	locs := make([]Loc, numLocs)
	for i := 0; i < numLocs; i++ {
		line := dataLines[i]
		lineSplitted := strings.Split(dataLines[i][1:len(line)-1], ",")
		x, _ := strconv.Atoi(lineSplitted[0])
		y, _ := strconv.Atoi(lineSplitted[1])
		locs[i] = Loc{x: x, y: y}
	}

	// Parse route
	dataRoute := dataLines[numLocs : len(dataLines)-1]
	route := make([]int, len(dataRoute))
	for i := 0; i < len(dataRoute); i++ {
		idx, _ := strconv.Atoi(dataRoute[i])
		route[i] = idx
	}

	// Initialize grid
	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

	// Loop through route
	maxPkgCount := 0
	for i := 0; i < len(route)-1; i++ {
		startLoc := locs[route[i]]
		endLoc := locs[route[i+1]]
		xSteps := int(math.Abs(float64(endLoc.x - startLoc.x)))
		ySteps := int(math.Abs(float64(endLoc.y - startLoc.y)))
		xPos := true
		if endLoc.x < startLoc.x {
			xPos = false
		}
		yPos := true
		if endLoc.y < startLoc.y {
			yPos = false
		}
		for xStep := 0; xStep < xSteps; xStep++ {
			for yStep := 0; yStep < ySteps; yStep++ {
				x := startLoc.x
				if xPos {
					x += xStep
				} else {
					x -= xStep
				}
				y := startLoc.y
				if yPos {
					y += yStep
				} else {
					y -= yStep
				}
				grid[y][x] += 1
				if grid[y][x] > maxPkgCount {
					maxPkgCount = grid[y][x]
				}
			}
		}
	}

	// Make 0-1 grid where 1 == maxPkgCount and 0 < maxPkgCount
	binaryGrid := make([][]int, gridSize)
	for j := 0; j < gridSize; j++ {
		binaryGrid[j] = make([]int, gridSize)
		for i := 0; i < gridSize; i++ {
			if grid[j][i] == maxPkgCount {
				binaryGrid[j][i] = 1
			} else {
				binaryGrid[j][i] = 0
			}
		}
	}

	bottomLeft, topRight := findLargestRect(binaryGrid, maxPkgCount)
	fmt.Printf("%v,%v %v,%v\n", bottomLeft.x, bottomLeft.y, topRight.x, topRight.y)
}

func main() {
	solve("input.txt", 200, 1000)
}
