package main

import (
	"fmt"
	"strings"

	"github.com/Acksell/aoc2021/util"
)

type Population struct {
	fish map[int]int64
}

func (p *Population) Load(s string) error {
	fish := util.StringsToInt(strings.Split(s, ","))
	for _, f := range fish {
		p.fish[f] += 1
	}
	return nil
}

func (p *Population) Simulate(days int) int64 {
	var sum int64
	for d := 0; d < days; d++ {
		p.Step()
	}
	for _, v := range p.fish {
		sum += v
	}
	return sum
}

func (p *Population) Step() {
	newfish := p.fish[0]
	for i := 0; i < 8; i++ {
		p.fish[i] = p.fish[i+1]
	}
	p.fish[6] += newfish
	p.fish[8] = newfish
}

func main() {
	p := &Population{make(map[int]int64)}
	util.ReadLines("inputs/fish.txt", p)
	fmt.Println(p.Simulate(80))
	// NOTE: don't reuse the same struct for part 1 and 2 ever again.
	// fmt.Println(p.Simulate(256))
	p2 := &Population{make(map[int]int64)}
	util.ReadLines("inputs/fish.txt", p2)
	fmt.Println(p2.Simulate(256))
}
