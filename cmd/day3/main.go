package main

import (
	"fmt"
	"strconv"

	"github.com/Acksell/aoc2021/util"
)

const inputFilePath = "inputs/report.txt"

type PowerCounter struct {
	length, n int
	numzero   []int
}

func (b *PowerCounter) Load(bitstring string) error {
	if b.numzero == nil {
		b.length = len(bitstring)
		b.numzero = make([]int, len(bitstring))
	}
	for i, bit := range bitstring {
		if bit == '0' {
			b.numzero[i] += 1
		}
	}
	b.n += 1
	return nil
}

func (b *PowerCounter) Gamma() int {
	gamma := 0
	for i, count := range b.numzero {
		if count < b.n/2.0 {
			gamma += 1 << (b.length - i - 1)
		}
	}
	return gamma
}

func (b *PowerCounter) Epsilon() int {
	k := (1 << (b.length)) - 1
	return b.Gamma() ^ k
}

type LifeCounter struct {
	numbers []string
	length  int
}

func findRatings(numbers []string, bit int, o2 bool) string {
	if len(numbers) == 1 {
		return numbers[0]
	}
	var numzero int
	var ones, zeros []string
	for _, n := range numbers {
		if n[bit] == '0' {
			numzero += 1
			zeros = append(zeros, n)
		} else {
			ones = append(ones, n)
		}
	}
	if o2 { // can probably be simplified
		if numzero <= len(numbers)/2.0 {
			return findRatings(ones, bit+1, o2)
		} else {
			return findRatings(zeros, bit+1, o2)
		}
	} else {
		if numzero <= len(numbers)/2.0 {
			return findRatings(zeros, bit+1, o2)
		} else {
			return findRatings(ones, bit+1, o2)
		}
	}
}

func (b *LifeCounter) Load(bitstring string) error {
	b.length = len(bitstring)
	b.numbers = append(b.numbers, bitstring)
	return nil
}

func (b *LifeCounter) O2() int {
	n, _ := strconv.ParseInt(findRatings(b.numbers, 0, true), 2, 64)
	return int(n)
}

func (b *LifeCounter) CO2() int {
	n, _ := strconv.ParseInt(findRatings(b.numbers, 0, false), 2, 64)
	return int(n)
}

func main() {
	c := &PowerCounter{}
	util.ReadLines(inputFilePath, c)
	fmt.Println(c.Gamma() * c.Epsilon())

	c2 := &LifeCounter{}
	util.ReadLines(inputFilePath, c2)
	o2, co2 := c2.O2(), c2.CO2()
	fmt.Println(o2 * co2)
}
