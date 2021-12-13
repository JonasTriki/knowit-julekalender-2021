package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	// Read and parse names
	namesRaw, _ := ioutil.ReadFile("names.txt")
	namesMat := make([][]string, 0)
	for _, name := range strings.Split(string(namesRaw), "\n") {
		names := make([]string, 0)
		if len(name) == 0 {
			continue
		}
		names = append(names, name)
		for i := 0; i < len(name)-1; i++ {
			newName := ""
			for j := 0; j < len(name); j++ {
				if j == i {
					newName += string(name[i+1])
				} else if j == i+1 {
					newName += string(name[i])
				} else {
					newName += string(name[j])
				}
			}
			if !contains(names, newName) {
				names = append(names, newName)
			}
		}
		namesMat = append(namesMat, names)
	}

	// Read and decrypt locked names
	lockedNamesRaw, _ := ioutil.ReadFile("locked.txt")
	lockedNames := strings.Split(string(lockedNamesRaw), "\n")
	nameCount := map[string]int{}
	for _, lockedName := range lockedNames {
		if len(lockedName) == 0 {
			continue
		}

		nameOverflow := map[string]int{}
		for _, names := range namesMat {
			trueName := names[0]
			for _, name := range names {
				i := 0
				startPos := -1
				endPos := -1
				for j := 0; j < len(lockedName); j++ {
					if lockedName[j] == name[i] {
						if startPos == -1 {
							startPos = j
						}
						if i == len(name)-1 {
							endPos = j
							break
						} else {
							i++
						}
					}
				}
				if startPos != -1 && endPos != -1 {
					segLen := endPos - startPos + 1
					numOverflow := segLen - len(name)
					nameOverflow[trueName] = numOverflow
					break
				}
			}
		}

		smallestOverflow := -1
		bestName := ""
		for name, overflow := range nameOverflow {
			if smallestOverflow == -1 || overflow < smallestOverflow {
				smallestOverflow = overflow
				bestName = name
			} else if overflow == smallestOverflow {
				smallestOverflow = -1
				break
			}
		}
		if smallestOverflow != -1 {
			nameCount[bestName]++
		}
	}

	// Find max and print
	maxCount := -1
	maxName := ""
	for name, count := range nameCount {
		if count > maxCount {
			maxCount = count
			maxName = name
		}
	}

	// Svaret er mindy,26
	fmt.Printf("%v,%v\n", maxName, maxCount)
}
