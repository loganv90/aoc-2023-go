package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Println("solution", part1("input1", 6))
    fmt.Println("solution", part1("input", 64))
    part2("input1", 6)
    part2("input1", 10)
    part2("input1", 50)
    part2("input1", 100)
    part2("input1", 500)
    part2("input1", 1000)
    part2("input1", 5000)
    part2("input", 26501365)
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

func linesToGrid(lines []string) [][]string {
    grid := [][]string{}
    for _, line := range lines {
        if len(line) == 0 { continue }
        row := []string{}
        for _, c := range line {
            row = append(row, string(c))
        }
        grid = append(grid, row)
    }
    return grid
}

type Point struct {
    row int
    col int
}

var directions = []Point{
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0},
}

func part1(inputFilename string, maxSteps int) int {
    lines := getLines(inputFilename)
    grid := linesToGrid(lines)
    visited := map[Point]bool{}

    start := Point{-1, -1}
    for row, line := range grid {
        for col, c := range line {
            if c == "S" {
                start = Point{row, col}
            }
        }
    }
    if start.row == -1 || start.col == -1 {
        panic("no start")
    }

    step := 0
    endCount := 0
    var queue chan Point = make(chan Point, 100000)
    queue <- start

    for len(queue) > 0 && step <= maxSteps {
        lenQueue := len(queue)
        for i := 0; i < lenQueue; i++ {
            select {
            case p := <-queue:
                if _, ok := visited[p]; ok {
                    continue
                }
                visited[p] = true

                if step % 2 == 0 && maxSteps % 2 == 0 {
                    endCount++
                } else if step % 2 == 1 && maxSteps % 2 == 1 {
                    endCount++
                }

                for _, dir := range directions {
                    next := Point{p.row + dir.row, p.col + dir.col}

                    row := next.row
                    col := next.col
                    for row < 0 {
                        row += len(grid)
                    }
                    for col < 0 {
                        col += len(grid[0])
                    }
                    row = row % len(grid)
                    col = col % len(grid[0])

                    if grid[row][col] == "#" {
                        continue
                    }
                    queue <- next
                }
            default:
                panic("no more")
            }
        }
        step++
    }

    return endCount
}

func part2(inputFilename string, maxSteps int) {
    lines := getLines(inputFilename)
    size := len(lines[0])
    halfSize := size/2

    y0 := part1(inputFilename, halfSize + 0*size)
    y1 := part1(inputFilename, halfSize + 1*size)
    y2 := part1(inputFilename, halfSize + 2*size)

    n := (maxSteps - halfSize)/size
    a := (y2 - 2*y1 + y0)/2
    b := y1 - y0 - a
    c := y0
    res := a*n*n + b*n + c

    fmt.Println("solution", res)
}

