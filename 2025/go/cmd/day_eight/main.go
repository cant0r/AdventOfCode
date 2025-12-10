package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/cant0r/AdventOfCode/2025/go/internal/utilities"
)

//go:embed input.txt
var inputData string

// float64 cause go.math :)
type Coordinate struct {
	X float64
	Y float64
	Z float64
}

func (c Coordinate) distanceFrom(other Coordinate) float64 {
	if c == other {
		return math.Inf(1)
	}

	dx := math.Pow(c.X-other.X, 2)
	dy := math.Pow(c.Y-other.Y, 2)
	dz := math.Pow(c.Z-other.Z, 2)

	return math.Sqrt(dx + dy + dz)
}

func loadCoordinates() []Coordinate {
	var coordinates []Coordinate

	for _, junctionBox := range strings.Split(inputData, "\n") {
		coordinate := strings.Split(junctionBox, ",")

		x, _ := strconv.ParseFloat(coordinate[0], 64)
		y, _ := strconv.ParseFloat(coordinate[1], 64)
		z, _ := strconv.ParseFloat(coordinate[2], 64)

		coordinates = append(coordinates, Coordinate{
			X: x,
			Y: y,
			Z: z,
		})
	}

	slices.SortFunc(coordinates, func(a, b Coordinate) int {
		return cmp.Compare(a.X, b.X)
	})

	return coordinates
}

func findJunctionBoxesByDistance(junctionBoxes []Coordinate, distanceMap map[Coordinate][]float64, distance float64) (junktionBoxOne, junctionBox2 Coordinate) {
	for coordinate, distances := range distanceMap {
		for i, dist := range distances {
			if dist == distance {
				return coordinate, junctionBoxes[i]
			}
		}
	}

	return Coordinate{}, Coordinate{}
}

func buildDistanceMap(junctionBoxes []Coordinate) map[Coordinate][]float64 {
	distanceMap := make(map[Coordinate][]float64)
	for _, junctionBox := range junctionBoxes {
		for _, otherJunctionBox := range junctionBoxes {
			if distance, ok := distanceMap[junctionBox]; ok == false {
				distanceMap[junctionBox] = []float64{junctionBox.distanceFrom(otherJunctionBox)}
			} else {
				distanceMap[junctionBox] = append(distance, junctionBox.distanceFrom(otherJunctionBox))
			}
		}
	}
	return distanceMap
}

func find(circuits map[Coordinate]Coordinate, coordinate Coordinate) Coordinate {
	if circuits[coordinate] == coordinate {
		return coordinate
	}
	circuits[coordinate] = find(circuits, circuits[coordinate])
	return circuits[coordinate]
}

func union(circuits map[Coordinate]Coordinate, coordinate1, coordinate2 Coordinate) {
	a := find(circuits, coordinate1)
	b := find(circuits, coordinate2)
	if a != b {
		circuits[a] = b
	}
}

func unionSetCount(circuits map[Coordinate]Coordinate, coordinate Coordinate) int {
	nodes := slices.Collect(maps.Values(circuits))
	count := 0
	if !slices.Contains(nodes, coordinate) {
		return 0
	} else {
		children := make([]Coordinate, 0)

		for child, origin := range circuits {
			if origin == coordinate && child != coordinate {
				children = append(children, child)
			}
		}

		count += len(children)

		for _, child := range children {
			count += unionSetCount(circuits, child)
		}

		return count
	}
}

func printCircuits(circuits map[Coordinate]Coordinate) {
	for node, origin := range circuits {
		fmt.Printf("%v <= %v\n", node, origin)
	}
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	var result int = 0
	var _distances []float64 = nil
	var distances []float64 = nil
	var junctionBoxes []Coordinate = loadCoordinates()
	var circuits map[Coordinate]Coordinate = make(map[Coordinate]Coordinate)

	distanceMap := buildDistanceMap(junctionBoxes)
	for _, junktionBoxDistances := range distanceMap {
		_distances = append(_distances, junktionBoxDistances...)
	}

	slices.SortFunc(_distances, func(a, b float64) int {
		return cmp.Compare(a, b)
	})

	for i := 0; i < len(_distances); i++ {
		if utilities.Mod(i, 2) == 0 {
			distances = append(distances, _distances[i])
		}
	}

	for _, junkctionBox := range junctionBoxes {
		circuits[junkctionBox] = junkctionBox
	}

	for _, distance := range distances[:1000] {
		j1, j2 := findJunctionBoxesByDistance(junctionBoxes, distanceMap, distance)
		union(circuits, j1, j2)
	}

	//printCircuits(circuits)

	var rootJuntionBoxes []Coordinate = nil

	for junctionBox, _ := range circuits {
		if circuits[junctionBox] == junctionBox {
			rootJuntionBoxes = append(rootJuntionBoxes, junctionBox)
		}
	}

	var setCounts []int = make([]int, len(rootJuntionBoxes))

	for i, rootJunctionBox := range rootJuntionBoxes {
		setCounts[i] = unionSetCount(circuits, rootJunctionBox) + 1
	}

	slices.Sort(setCounts)

	result += setCounts[len(setCounts)-1] * setCounts[len(setCounts)-2] * setCounts[len(setCounts)-3]

	utilities.PrintResult(result)
}

func partTwo() {
	var result int = 0
	var _distances []float64 = nil
	var distances []float64 = nil
	var junctionBoxes []Coordinate = loadCoordinates()
	var circuits map[Coordinate]Coordinate = make(map[Coordinate]Coordinate)

	distanceMap := buildDistanceMap(junctionBoxes)
	for _, junktionBoxDistances := range distanceMap {
		_distances = append(_distances, junktionBoxDistances...)
	}

	slices.SortFunc(_distances, func(a, b float64) int {
		return cmp.Compare(a, b)
	})

	for i := 0; i < len(_distances); i++ {
		if utilities.Mod(i, 2) == 0 {
			distances = append(distances, _distances[i])
		}
	}

	for _, junkctionBox := range junctionBoxes {
		circuits[junkctionBox] = junkctionBox
	}

	for _, distance := range distances {
		if math.IsInf(distance, 1) {
			break
		}
		j1, j2 := findJunctionBoxesByDistance(junctionBoxes, distanceMap, distance)
		union(circuits, j1, j2)

		var rootJuntionBoxes []Coordinate = nil

		for junctionBox, _ := range circuits {
			if circuits[junctionBox] == junctionBox {
				rootJuntionBoxes = append(rootJuntionBoxes, junctionBox)
			}
		}

		var setCounts []int = make([]int, len(rootJuntionBoxes))

		for i, rootJunctionBox := range rootJuntionBoxes {
			setCounts[i] = unionSetCount(circuits, rootJunctionBox) + 1
		}

		if len(setCounts) == 1 {
			result += int(j1.X * j2.X)
			break
		}
	}

	//printCircuits(circuits)

	utilities.PrintResult(result)
}
