package geometry

import (
	"math"
)

type IFigure interface {
	Area() float64
	Perimeter() float64
	Contains(point Point) bool
}

type IPoint interface {
	DistanceTo(other Point) float64
}

type Point struct {
	X float64
	Y float64
}

type Polygon struct {
	PointsArr []Point
}

type Circle struct {
	Center Point
	Radius float64
}

// Point

func (p Point) DistanceTo(other Point) float64 {

	distanceX := other.X - p.X
	distanceY := other.Y - p.Y

	distanceXSquare := math.Pow(distanceX, 2)
	distanceYSquare := math.Pow(distanceY, 2)

	distance := math.Sqrt(float64(distanceXSquare + distanceYSquare))

	return  distance
}

//  Circle

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Contains(point Point) bool {
	return c.Center.DistanceTo(point) <= c.Radius
}

// Polygon

func (p Polygon) Area() float64 {

	n := len(p.PointsArr)

	if n < 3 {
		return 0
	}

	sum := 0.0

	for i := 0; i < n; i++ {

		current := p.PointsArr[i]
		next := p.PointsArr[(i+1)%n]
		sum += current.X*next.Y - next.X*current.Y

	}

	return math.Abs(sum) / 2
}

func (p Polygon) Perimeter() float64 {
	n := len(p.PointsArr)

	if n < 2 {
		return 0
	}

	sum := 0.0

	for i := 0; i < n; i++ {
		current := p.PointsArr[i]
		next := p.PointsArr[(i+1)%n]
		sum += current.DistanceTo(next)
	}

	return sum
}

func (p Polygon) Contains(point Point) bool {

	n := len(p.PointsArr)

	inside := false

	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		pi := p.PointsArr[i]
		pj := p.PointsArr[j]
		if (pi.Y > point.Y) != (pj.Y > point.Y) &&
			point.X < (pj.X-pi.X)*(point.Y-pi.Y)/(pj.Y-pi.Y)+pi.X {
			inside = !inside
		}
	}

	return inside
}

