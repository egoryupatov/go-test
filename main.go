package main

import (
	"flag"
	"fmt"
	"geometry/geometry"
	"strconv"
	"strings"
)

type pointList []geometry.Point

func (pl *pointList) String() string { return "" }

func (pl *pointList) Set(value string) error {
	parts := strings.Split(value, ",")          
	x, _ := strconv.ParseFloat(parts[0], 64) 
	y, _ := strconv.ParseFloat(parts[1], 64)
	*pl = append(*pl, geometry.Point{X: x, Y: y})
	return nil
}

func main() {
	distance := flag.Bool("distance", false, "расстояние")
	area := flag.Bool("area", false, "площадь")
	contains := flag.Bool("contains", false, "вхождение точки")
	circle := flag.String("circle", "", "круг X,Y,R")
	polygon := flag.Bool("polygon", false, "многоугольник")

	var points pointList
	flag.Var(&points, "point", "точка X,Y")

	flag.Parse()

	switch {
	case *distance:
		d := points[0].DistanceTo(points[1])
		fmt.Println(d)

	case *area:
		fig := makeFigure(*circle, *polygon, points)
		fmt.Println(fig.Area())

	case *contains:
		fig := makeFigure(*circle, *polygon, points[1:])
		fmt.Println(fig.Contains(points[0]))
	}
}

func makeFigure(circle string, polygon bool, points pointList) geometry.IFigure {
	if polygon {
		return geometry.Polygon{PointsArr: points}
	}
	
	parts := strings.Split(circle, ",")
	x, _ := strconv.ParseFloat(parts[0], 64)
	y, _ := strconv.ParseFloat(parts[1], 64)
	r, _ := strconv.ParseFloat(parts[2], 64)
	return geometry.Circle{Center: geometry.Point{X: x, Y: y}, Radius: r}
}