package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2021/util"
)

type Square struct {
	Num, Row, Col int
	Seen          bool
}

type PlayerBoard struct {
	board  [][]*Square
	lookup map[int]*Square
}

type Bingo struct {
	Boards []*PlayerBoard
	Draws  []int
}

func (b *Bingo) Load(f string) error {
	in := strings.Split(f, "\n\n")
	nums := in[0]
	for _, n := range strings.Split(nums, ",") {
		i, _ := strconv.Atoi(string(n))
		b.Draws = append(b.Draws, i)
	}
	for _, s := range in[1:] {
		rows := strings.Split(s, "\n")
		pb := &PlayerBoard{board: make([][]*Square, len(rows)), lookup: make(map[int]*Square)}

		for i, row := range rows {
			cols := util.StringsToInt(strings.Fields(row))
			pb.board[i] = make([]*Square, len(cols))
			for j, num := range cols {
				s := &Square{Num: num, Row: i, Col: j}
				pb.board[i][j] = s
				pb.lookup[num] = s
			}
		}
		b.Boards = append(b.Boards, pb)
	}
	return nil
}

func (b *Bingo) Play() (int, *PlayerBoard) {
	for _, d := range b.Draws {
		for _, pb := range b.Boards {
			pb.Update(d)
			if pb.CheckRows() || pb.CheckColumns() {
				return d, pb
			}
		}
	}
	return 0, nil
}

func (b *Bingo) Play2() (int, *PlayerBoard) {
	winners := make(map[int]bool)
	for _, d := range b.Draws {
		for i, pb := range b.Boards {
			pb.Update(d)
			if pb.CheckRows() || pb.CheckColumns() {
				winners[i] = true
				if len(winners) == len(b.Boards) {
					return d, pb
				}
			}
		}
	}
	return 0, nil
}

func (b *PlayerBoard) CheckRows() bool {
	for _, row := range b.board {
		win := true
		for _, s := range row {
			if !s.Seen {
				win = false
				break
			}
		}
		if win {
			return win
		}
	}
	return false
}

func (b *PlayerBoard) CheckColumns() bool {
	for j := 0; j < len(b.board); j++ {
		win := true
		for i := 0; i < len(b.board); i++ {
			if !b.board[i][j].Seen {
				win = false
				break
			}
		}
		if win {
			return win
		}
	}
	return false
}

func (b *PlayerBoard) Update(draw int) {
	s, ok := b.lookup[draw]
	if !ok {
		return
	}
	s.Seen = true
	b.board[s.Row][s.Col] = s
	b.lookup[draw] = s
}

func (b *PlayerBoard) Sum() int {
	sum := 0
	for n, s := range b.lookup {
		if !s.Seen {
			sum += n
		}
	}
	return sum
}

func main() {
	bingo := &Bingo{}
	util.LoadFile("inputs/bingo.txt", bingo)
	draw, winner := bingo.Play()
	if winner != nil {
		fmt.Println(winner.Sum() * draw)
	}
	draw2, loser := bingo.Play2()
	if loser != nil {
		fmt.Println(loser.Sum() * draw2)
	}
}
