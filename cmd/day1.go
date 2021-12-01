package main

import (
	"fmt"
	"strconv"

	"github.com/Acksell/aoc2021/util"
)

const inputFilePath = "inputs/depths.txt"

type Counter struct {
	depth   *int
	counter int
}

type SlidingCounter struct {
	depth1  *int
	depth2  *int
	depth3  *int
	counter int
}

func (c *Counter) Load(line string) error {
	depth, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("failed to parse %s", line)
	}
	if c.depth != nil {
		if *c.depth < depth {
			c.counter += 1
		}
	}
	c.depth = &depth
	return nil
}

func (c *SlidingCounter) Load(line string) error {
	depth, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("failed to parse %s", line)
	}
	if c.depth1 == nil {
		c.depth1 = &depth
	} else if c.depth2 == nil {
		c.depth2 = &depth
	} else if c.depth3 == nil {
		c.depth3 = &depth
	} else {
		prevsum := *c.depth1 + *c.depth2 + *c.depth3
		newsum := *c.depth2 + *c.depth3 + depth
		c.depth1 = c.depth2
		c.depth2 = c.depth3
		c.depth3 = &depth
		if newsum > prevsum {
			c.counter += 1
		}
	}
	return nil
}

func main() {
	c := &Counter{}
	util.ReadLines(inputFilePath, c)
	fmt.Println(c.counter)

	c2 := &SlidingCounter{}
	util.ReadLines(inputFilePath, c2)
	fmt.Println(c2.counter)
}
