package main

import (
	"fmt"
	"strings"

	"github.com/Acksell/aoc2021/util"
)

type Line struct {
	p1, p2 Point
}

func (l Line) Sweep() (points []Point) {
	if l.p1.x == l.p2.x {
		diff := l.p2.y - l.p1.y
		for i := 0; i <= util.Abs(diff); i++ {
			points = append(points, Point{l.p1.x, l.p1.y + util.Sign(diff)*i})
		}
	} else if l.p1.y == l.p2.y {
		diff := l.p2.x - l.p1.x
		for i := 0; i <= util.Abs(diff); i++ {
			points = append(points, Point{l.p1.x + util.Sign(diff)*i, l.p1.y})
		}
	} else {
		diffx := l.p2.x - l.p1.x
		diffy := l.p2.y - l.p1.y
		for i := 0; i <= util.Abs(diffx); i++ { // assume diffx=diffy
			points = append(points, Point{l.p1.x + util.Sign(diffx)*i, l.p1.y + util.Sign(diffy)*i})
		}
	}
	return
}

type Point struct {
	x, y int
}

type Counter struct {
	counts map[Point]int
	lines  []Line
}

func (c *Counter) Load(s string) error {
	points := strings.Split(s, " -> ")
	p1 := util.StringsToInt(strings.Split(points[0], ","))
	p2 := util.StringsToInt(strings.Split(points[1], ","))
	c.lines = append(c.lines, Line{Point{p1[0], p1[1]}, Point{p2[0], p2[1]}})
	return nil
}

func (c *Counter) Sweep() {
	for _, l := range c.lines {
		for _, p := range l.Sweep() {
			c.counts[p] += 1
		}
	}
}

func (c *Counter) Count() int {
	sum := 0
	for _, v := range c.counts {
		if v >= 2 {
			sum += 1
		}
	}
	return sum
}

func main() {
	c := &Counter{make(map[Point]int), nil}
	util.ReadLines("inputs/lines.txt", c)
	c.Sweep()
	fmt.Println(c.Count())
	// fmt.Println(c.counts)
}
