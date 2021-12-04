package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkt"
)

func Map(vs [][]string, f func([]string) orb.Point) []orb.Point {
	vsm := make([]orb.Point, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func readData() []orb.Point {
	f, _ := os.Open("data/cities.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.Read() // Skip header
	records, _ := csvReader.ReadAll()
	return Map(records, func(s []string) orb.Point {
		point, _ := wkt.UnmarshalPoint(s[1])
		return point
	})
}

const earthRadius = 6371.0

func deg2rad(d float64) float64 {
	return d * math.Pi / 180.0
}

func Haversine(p, q orb.Point) (km float64) {
	lat1 := deg2rad(p.Lat())
	lon1 := deg2rad(p.Lon())
	lat2 := deg2rad(q.Lat())
	lon2 := deg2rad(q.Lon())

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRadius
}

func main() {
	data := readData()
	northPole := orb.Point{0.0, 90.0}
	currentPoint := northPole
	accDist := 0.0
	for len(data) > 0 {

		// Sorter data
		sort.Slice(data, func(i, j int) bool {
			return Haversine(currentPoint, data[i]) < Haversine(currentPoint, data[j])
		})

		// Regn ut avstand til neste destinasjon og legg til akkumulator
		accDist += Haversine(currentPoint, data[0])
		currentPoint = data[0]
		data = data[1:]
	}
	accDist += Haversine(currentPoint, northPole)

	fmt.Println(math.Round(accDist))
}
