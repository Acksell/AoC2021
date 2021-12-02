package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2021/util"
)

type SimplePosition struct {
	x, y int
}

type AimedPosition struct {
	x, y, aim int
}

func (p *SimplePosition) Load(s string) error {
	l := strings.Split(s, " ")
	command := l[0]
	n, _ := strconv.Atoi(l[1])
	switch command {
	case "forward":
		p.y += n
	case "down":
		p.x += n
	case "up":
		p.x -= n
	}
	return nil
}

func (p *AimedPosition) Load(s string) error {
	l := strings.Split(s, " ")
	command := l[0]
	n, _ := strconv.Atoi(l[1])
	switch command {
	case "forward":
		p.y += n
		p.x += p.aim * n
	case "down":
		p.aim += n
	case "up":
		p.aim -= n
	}
	return nil
}

func main() {
	p := &SimplePosition{}
	util.ReadLines("inputs/course.txt", p)
	fmt.Println(p.x * p.y)

	p2 := &AimedPosition{}
	util.ReadLines("inputs/course.txt", p2)
	fmt.Println(p2.x * p2.y)
}
