package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type PointsWithDistance struct {
	a        Point
	b        Point
	distance float64
}

func distanceBetween(a, b Point) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2))
}

func calculateDistances(points []Point) []PointsWithDistance {
	result := []PointsWithDistance{}
	for i, outerPoint := range points {
		for j, innerPoint := range points {
			if i != j {
				distance := distanceBetween(outerPoint, innerPoint)
				// ? outerPoint,innerPoint,distance
				result = append(result, PointsWithDistance{a: outerPoint, b: innerPoint, distance: distance})
			}
		}
	}
	return result
}

func reduce(distances []PointsWithDistance, f func(PointsWithDistance, PointsWithDistance) bool) PointsWithDistance {
	result := distances[0]
	for _, points := range distances {
		if f(points, result) {
			result = points
		}
	}
	return result
}

func findMinDistance(distances []PointsWithDistance) PointsWithDistance {
	return reduce(distances, func(a, b PointsWithDistance) bool {
		return a.distance < b.distance
	})
}

func findMaxDistance(distances []PointsWithDistance) PointsWithDistance {
	return reduce(distances, func(a, b PointsWithDistance) bool {
		return a.distance > b.distance
	})
}

func main() {

	points := []Point{{x: 1.0, y: 4}, {x: 4, y: 4}, {x: 3, y: 2}, {x: 5, y: 1}}
	distances := calculateDistances(points)
	min := findMinDistance(distances)
	fmt.Println(min)

	max := findMaxDistance(distances)
	fmt.Println(max)

}
