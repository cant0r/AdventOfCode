package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/cant0r/AdventOfCode/2025/go/internal/utilities"
)

//go:embed input.txt
var inputData string

type Coordinate struct {
	Row int
	Col int
}

// Assuming coordinates are positive lol
func (c *Coordinate) areaWith(other *Coordinate) int {
	oneSide := math.Abs(float64(c.Row)-float64(other.Row)) + 1
	otherSide := math.Abs(float64(c.Col)-float64(other.Col)) + 1

	return int(oneSide * otherSide)
}

type Side struct {
	Origin Coordinate
	End    Coordinate
}

func (s Side) equalsTo(other Side) bool {
	if s.Origin == other.Origin && s.End == other.End {
		return true
	}

	if s.Origin == other.End && s.End == other.Origin {
		return true
	}
	return false
}

func main() {
	partOne()
	// Later, ray tracing implementation is shit, need to debug
	//partTwo()
}

func partOne() {
	var result int = 0
	var coordinates []Coordinate = nil

	for _, dataRow := range strings.Split(inputData, "\n") {
		coordinatePair := strings.Split(dataRow, ",")
		row, _ := strconv.Atoi(coordinatePair[0])
		col, _ := strconv.Atoi(coordinatePair[1])

		coordinates = append(coordinates, Coordinate{Row: row, Col: col})
	}

	areas := make([]int, len(coordinates)*len(coordinates))

	for i := 0; i < len(coordinates); i++ {
		for j := 0; j < len(coordinates); j++ {
			area := coordinates[i].areaWith(&coordinates[j])
			// fmt.Printf("area=%d, coordinateI=%v coordinateJ=%v\n", area, coordinates[i], coordinates[j])
			areas[i*len(coordinates)+j] = area
		}
	}

	result = slices.Max(areas)

	utilities.PrintResult(result)
}

func partTwo() {
	var result int = 0
	var coordinates []Coordinate = nil

	for _, dataRow := range strings.Split(inputData, "\n") {
		coordinatePair := strings.Split(dataRow, ",")
		row, _ := strconv.Atoi(coordinatePair[1])
		col, _ := strconv.Atoi(coordinatePair[0])

		coordinates = append(coordinates, Coordinate{Row: row, Col: col})
	}

	sides := make([]Side, len(coordinates))
	sidesIndex := 0

	for i := 0; i < len(coordinates); i++ {
		for j := 0; j < len(coordinates); j++ {
			area := coordinates[i].areaWith(&coordinates[j])
			if area != 1 && (coordinates[i].Row == coordinates[j].Row || coordinates[i].Col == coordinates[j].Col) {
				newSide := Side{Origin: coordinates[i], End: coordinates[j]}
				if !slices.ContainsFunc(sides, func(side Side) bool {
					return side.equalsTo(newSide)
				}) {
					sides[sidesIndex] = newSide
					sidesIndex++
				}

			}
		}
	}

	areas := make([]int, len(coordinates)*len(coordinates))

	for i := 0; i < len(coordinates); i++ {
		for j := 0; j < len(coordinates); j++ {
			otherRowPoint := Coordinate{Row: coordinates[j].Row, Col: coordinates[i].Col}
			otherColPoint := Coordinate{Row: coordinates[i].Row, Col: coordinates[j].Col}

			//fmt.Printf("Origin %v otherRowPoint %v otherColPoint %v End %v\n", coordinates[i], otherRowPoint, otherColPoint, coordinates[j])

			pointsInArea := true

			pointsInArea = pointsInArea && isPointInArea(coordinates[i], sides)
			pointsInArea = pointsInArea && isPointInArea(otherRowPoint, sides)
			pointsInArea = pointsInArea && isPointInArea(otherColPoint, sides)
			pointsInArea = pointsInArea && isPointInArea(coordinates[j], sides)

			fmt.Printf("Points in are %v\n", pointsInArea)

			if pointsInArea {
				area := coordinates[i].areaWith(&coordinates[j])
				// fmt.Printf("area=%d, coordinateI=%v coordinateJ=%v\n", area, coordinates[i], coordinates[j])
				areas[i*len(coordinates)+j] = area
			}
		}
	}

	result = slices.Max(areas)

	utilities.PrintResult(result)
}

func isPointInArea(coordinate Coordinate, sides []Side) bool {
	intersections := 0
	for i := 0; i < len(sides); i++ {
		// We want the line in linear equation standard form: A*x + B*y + C = 0
		// See: http://en.wikipedia.org/wiki/Linear_equation
		a1 := coordinate.Col - 0
		b1 := 0 - coordinate.Row
		c1 := (coordinate.Row * 0) - (0 * coordinate.Col)

		d1 := (a1 * sides[i].Origin.Col) + (b1 * sides[i].Origin.Row) + c1
		d2 := (a1 * sides[i].End.Col) + (b1 * sides[i].End.Row) + c1

		// If d1 and d2 both have the same sign, they are both on the same side
		// of our line 1 and in that case no intersection is possible.
		if d1 > 0 && d2 > 0 {
			continue
		}
		if d1 < 0 && d2 < 0 {
			continue
		}

		// The fact that vector 2 intersected the infinite line 1 above doesn't
		// mean it also intersects the vector 1. Vector 1 is only a subset of that
		// infinite line 1, so it may have intersected that line before the vector
		// started or after it ended. To know for sure, we have to repeat the
		// the same test the other way round.
		a2 := sides[i].End.Row - sides[i].Origin.Row
		b2 := sides[i].Origin.Col - sides[i].End.Col
		c2 := (sides[i].End.Col * sides[i].Origin.Row) - (sides[i].Origin.Col * sides[i].End.Row)

		// Calculate d1 and d2 again, this time using points of vector 1.
		d1 = (a2 * 0) + (b2 * 0) + c2
		d2 = (a2 * coordinate.Row) + (b2 * coordinate.Col) + c2

		// Again, if both have the same sign (and neither one is 0),
		// no intersection is possible.
		if d1 > 0 && d2 > 0 {
			continue
		}
		if d1 < 0 && d2 < 0 {
			continue
		}

		// If we get here, only two possibilities are left. Either the two
		// vectors intersect in exactly one point or they are collinear, which
		// means they intersect in any number of points from zero to infinite.
		if (a1*b2)-(a2*b1) == 0 {
			continue
		}

		// If they are not collinear, they must intersect in exactly one point.
		intersections++
	}

	fmt.Printf("Intersections %d\n", intersections)

	return utilities.Mod(intersections, 2) != 0
}
