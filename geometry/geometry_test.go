package geometry_test

import (
	"geometry/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceTo(t *testing.T) {

    p1 := geometry.Point{X: 0, Y: 0}
    p2 := geometry.Point{X: 3, Y: 4}

    got := p1.DistanceTo(p2)

	assert.InDelta(t, 5.0, got, 0.001)
}

func TestCircleArea(t *testing.T) {
	circle := geometry.Circle{
		Center: geometry.Point{X: 2, Y: 5},
		Radius: 10,
	}

	got := circle.Area()

	assert.InDelta(t, 314.1592653589793, got, 0.0001)
}

func TestCirclePerimeter(t *testing.T) {
	circle := geometry.Circle{Radius: 10}

	got := circle.Perimeter()

	assert.Equal(t, 62.83185307179586, got)
}

func TestCircleContains(t *testing.T) {
	circle := geometry.Circle{
		Center: geometry.Point{X: 0, Y: 0},
		Radius: 5,
	}

	point := geometry.Point{X: 3, Y: 4}

	got := circle.Contains(point)

	assert.True(t, got)
}

func TestPolygonArea(t *testing.T) {
	p := geometry.Polygon{
		PointsArr: []geometry.Point{
			{X: 0, Y: 0},
			{X: 4, Y: 0},
			{X: 4, Y: 3},
			{X: 0, Y: 3},
		},
	}

	got := p.Area()

	assert.InDelta(t, 12.0, got, 0.0001)
}

func TestPolygonPerimeter(t *testing.T) {
	p := geometry.Polygon{
		PointsArr: []geometry.Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		},
	}

	got := p.Perimeter()

	assert.InDelta(t, 8.0, got, 0.0001)
}

func TestPolygonContains(t *testing.T) {
	p := geometry.Polygon{
		PointsArr: []geometry.Point{
			{X: 0, Y: 0},
			{X: 4, Y: 0},
			{X: 4, Y: 4},
			{X: 0, Y: 4},
		},
	}

	got := p.Contains(geometry.Point{X: 2, Y: 2})

	assert.True(t, got)
}