package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    part1("input1")
    part1("input")
    part2("input1")
    part2("input")
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

type Space struct {
    fill string
    leavingUp bool
    leavingDown bool
    leavingLeft bool
    leavingRight bool
}

func (s *Space) Clear() {
    s.leavingUp = false
    s.leavingDown = false
    s.leavingLeft = false
    s.leavingRight = false
}

func (s *Space) nextBeams(beam *Beam) []*Beam {
    res := []*Beam{}

    if beam.movingUp {
        if s.fill == "." {
            if !s.leavingUp {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingUp: true})
                s.leavingUp = true
            }
        } else if s.fill == "|" {
            if !s.leavingUp {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingUp: true})
                s.leavingUp = true
            }
        } else if s.fill == "-" {
            if !s.leavingLeft {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingLeft: true})
                s.leavingLeft = true
            }
            if !s.leavingRight {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingRight: true})
                s.leavingRight = true
            }
        } else if s.fill == "/" {
            if !s.leavingRight {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingRight: true})
                s.leavingRight = true
            }
        } else if s.fill == "\\" {
            if !s.leavingLeft {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingLeft: true})
                s.leavingLeft = true
            }
        } else {
            panic("unknown space fill")
        }
    } else if beam.movingDown {
        if s.fill == "." {
            if !s.leavingDown {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingDown: true})
                s.leavingDown = true
            }
        } else if s.fill == "|" {
            if !s.leavingDown {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingDown: true})
                s.leavingDown = true
            }
        } else if s.fill == "-" {
            if !s.leavingLeft {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingLeft: true})
                s.leavingLeft = true
            }
            if !s.leavingRight {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingRight: true})
                s.leavingRight = true
            }
        } else if s.fill == "/" {
            if !s.leavingLeft {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingLeft: true})
                s.leavingLeft = true
            }
        } else if s.fill == "\\" {
            if !s.leavingRight {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingRight: true})
                s.leavingRight = true
            }
        } else {
            panic("unknown space fill")
        }
    } else if beam.movingLeft {
        if s.fill == "." {
            if !s.leavingLeft {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingLeft: true})
                s.leavingLeft = true
            }
        } else if s.fill == "|" {
            if !s.leavingUp {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingUp: true})
                s.leavingUp = true
            }
            if !s.leavingDown {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingDown: true})
                s.leavingDown = true
            }
        } else if s.fill == "-" {
            if !s.leavingLeft {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingLeft: true})
                s.leavingLeft = true
            }
        } else if s.fill == "/" {
            if !s.leavingDown {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingDown: true})
                s.leavingDown = true
            }
        } else if s.fill == "\\" {
            if !s.leavingUp {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingUp: true})
                s.leavingUp = true
            }
        } else {
            panic("unknown space fill")
        }
    } else if beam.movingRight {
        if s.fill == "." {
            if !s.leavingRight {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingRight: true})
                s.leavingRight = true
            }
        } else if s.fill == "|" {
            if !s.leavingUp {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingUp: true})
                s.leavingUp = true
            }
            if !s.leavingDown {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingDown: true})
                s.leavingDown = true
            }
        } else if s.fill == "-" {
            if !s.leavingRight {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingRight: true})
                s.leavingRight = true
            }
        } else if s.fill == "/" {
            if !s.leavingUp {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingUp: true})
                s.leavingUp = true
            }
        } else if s.fill == "\\" {
            if !s.leavingDown {
                res = append(res, &Beam{row: beam.row, col: beam.col, movingDown: true})
                s.leavingDown = true
            }
        } else {
            panic("unknown space fill")
        }
    } else {
        panic("beam not moving")
    }

    return res
}

type Beam struct {
    row int
    col int
    movingUp bool
    movingDown bool
    movingLeft bool
    movingRight bool
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    spaces := [][]*Space{}
    for _, line := range lines {
        row := []*Space{}
        for _, char := range line {
            row = append(row, &Space{fill: string(char)})
        }
        spaces = append(spaces, row)
    }

    res := calculateEnergizedSpaces(spaces, &Beam{row: 0, col: -1, movingRight: true})
    fmt.Println("solution", res)
}

func calculateEnergizedSpaces(spaces [][]*Space, startBeam *Beam) int {
    beams := []*Beam{ startBeam }
    for len(beams) > 0 {
        newBeams := []*Beam{}

        for _, beam := range beams {
            if beam.movingUp {
                beam.row--
            } else if beam.movingDown {
                beam.row++
            } else if beam.movingLeft {
                beam.col--
            } else if beam.movingRight {
                beam.col++
            } else {
                panic("beam not moving")
            }

            if beam.row < 0 || beam.col < 0 || beam.row >= len(spaces) || beam.col >= len(spaces[beam.row]) {
                continue
            }

            space := spaces[beam.row][beam.col]
            nextBeams := space.nextBeams(beam)
            newBeams = append(newBeams, nextBeams...)
        }

        beams = newBeams
    }

    res := 0
    for _, row := range spaces {
        for _, space := range row {
            if space.leavingUp || space.leavingDown || space.leavingLeft || space.leavingRight {
                res++
            }
            space.Clear()
        }
    }

    return res
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    spaces := [][]*Space{}
    for _, line := range lines {
        row := []*Space{}
        for _, char := range line {
            row = append(row, &Space{fill: string(char)})
        }
        spaces = append(spaces, row)
    }

    res := 0
    for row := 0; row < len(spaces); row++ {
        res = max(res, calculateEnergizedSpaces(spaces, &Beam{row: row, col: -1, movingRight: true}))
        res = max(res, calculateEnergizedSpaces(spaces, &Beam{row: row, col: len(spaces[0]), movingLeft: true}))
    }
    for col := 0; col < len(spaces[0]); col++ {
        res = max(res, calculateEnergizedSpaces(spaces, &Beam{row: -1, col: col, movingDown: true}))
        res = max(res, calculateEnergizedSpaces(spaces, &Beam{row: len(spaces), col: col, movingUp: true}))
    }

    fmt.Println("solution", res)
}

