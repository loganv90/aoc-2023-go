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
    //part2("input")
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
        for _, char := range line {
            row = append(row, string(char))
        }
        grid = append(grid, row)
    }

    return grid
}

type Point struct {
    row int
    col int
}

type Direction struct {
    row int
    col int
    slope string
}

var directions = []Direction{
    {1, 0, "v"},
    {-1, 0, "^"},
    {0, 1, ">"},
    {0, -1, "<"},
}

func getStartPoint(grid [][]string) Point {
    startPoint := Point{-1, -1}

    for i, char := range grid[0] {
        if char == "." {
            startPoint = Point{0, i}
            break
        }
    }
    if startPoint.row == -1 || startPoint.col == -1 {
        panic("no start point found")
    }

    return startPoint
}

func getEndPoint(grid [][]string) Point {
    endPoint := Point{-1, -1}

    for i, char := range grid[len(grid) - 1] {
        if char == "." {
            endPoint = Point{len(grid) - 1, i}
            break
        }
    }
    if endPoint.row == -1 || endPoint.col == -1 {
        panic("no start point found")
    }

    return endPoint
}

func dfs(grid [][]string, currentPoint Point, endPoint Point, visited map[Point]bool) int {
    count := 0

    visited[currentPoint] = true

    for _, direction := range directions {
        nextRow := currentPoint.row + direction.row
        nextCol := currentPoint.col + direction.col
        if nextRow < 0 || nextRow >= len(grid) {
            continue
        }
        if nextCol < 0 || nextCol >= len(grid[nextRow]) {
            continue
        }

        nextPoint := Point{ row: nextRow, col: nextCol }
        if _, ok := visited[nextPoint]; ok {
            continue
        }

        nextChar := grid[nextPoint.row][nextPoint.col]
        if nextChar == "." || nextChar == direction.slope {
            count = max(count, dfs(grid, nextPoint, endPoint, visited))
        }
    }

    delete(visited, currentPoint)

    if count > 0 {
        return count + 1
    } else if currentPoint.row == endPoint.row && currentPoint.col == endPoint.col {
        return 1
    } else {
        return 0
    }
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)
    grid := linesToGrid(lines)

    startPoint := getStartPoint(grid)
    endPoint := getEndPoint(grid)

    longestPath := dfs(grid, startPoint, endPoint, map[Point]bool{}) - 1

    fmt.Println("solution", longestPath)
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)
    grid := linesToGrid(lines)
    fmt.Println(grid)

    fmt.Println("solution", 0)
}

