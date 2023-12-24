package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    part1("input1")
    part1("input")
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

func part1(inputFilename string) {
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

    for len(queue) > 0 && step <= 64 {
        lenQueue := len(queue)
        for i := 0; i < lenQueue; i++ {
            select {
            case p := <-queue:
                if _, ok := visited[p]; ok {
                    continue
                }
                visited[p] = true

                if step % 2 == 0 {
                    endCount++
                }

                for _, dir := range directions {
                    next := Point{p.row + dir.row, p.col + dir.col}
                    if next.row < 0 || next.row >= len(grid) {
                        continue
                    }
                    if next.col < 0 || next.col >= len(grid[next.row]) {
                        continue
                    }
                    if grid[next.row][next.col] == "#" {
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

    fmt.Println("solution", endCount)
}

