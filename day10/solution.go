package main

import (
	"fmt"
	"os"
	"strings"
    "slices"
)

func main() {
    part1("input1")
    part1("input2")
    part1("input3")
    part1("input4")
    part1("input")
    part2("input5")
    part2("input6")
    part2("input7")
    part2("input8")
    part2("input")
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

type Direction struct {
    dRow int
    dCol int
    connectors []rune
    connectees []rune
}

var directions = []Direction{
    {dRow: -1, dCol: 0, connectors: []rune{'|', 'F', '7'}, connectees: []rune{'|', 'J', 'L'}},
    {dRow: 1, dCol: 0, connectors: []rune{'|', 'L', 'J'}, connectees: []rune{'|', 'F', '7'}},
    {dRow: 0, dCol: -1, connectors: []rune{'-', 'F', 'L'}, connectees: []rune{'-', 'J', '7'}},
    {dRow: 0, dCol: 1, connectors: []rune{'-', 'J', '7'}, connectees: []rune{'-', 'F', 'L'}},
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    startRow := 0
    startCol := 0
    pipes := [][]rune{}
    for r, line := range lines {
        if len(line) == 0 { continue }
        
        row := []rune{}
        for c, char := range line {
            if char == 'S' {
                startRow = r
                startCol = c
            }
            row = append(row, char)
        }
        
        pipes = append(pipes, row)
    }

    res := pipeLength(pipes, startRow, startCol, startRow, startCol)
    if res % 2 != 0 { panic("odd") }
    fmt.Println("solution", res/2)
}

func pipeLength(pipes [][]rune, row int, col int, prevRow int, prevCol int) int {
    var currChar rune = pipes[row][col]
    var nextChar rune
    var nextRow int
    var nextCol int

    for _, dir := range directions {
        if !slices.Contains(dir.connectees, currChar) && currChar != 'S' { continue }

        nextRow = row + dir.dRow
        nextCol = col + dir.dCol

        if nextRow == prevRow && nextCol == prevCol { continue }
        if nextRow < 0 || nextRow >= len(pipes) { continue }
        if nextCol < 0 || nextCol >= len(pipes[nextRow]) { continue }

        nextChar = pipes[nextRow][nextCol]

        if !slices.Contains(dir.connectors, nextChar) && nextChar != 'S' { continue }

        break
    }

    if nextChar == 'S' {
        return 1
    }

    return 1 + pipeLength(pipes, nextRow, nextCol, row, col)
}

type Point struct {
    row int
    col int
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    startRow := 0
    startCol := 0
    pipes := [][]rune{}
    for r, line := range lines {
        if len(line) == 0 { continue }
        
        row := []rune{}
        for c, char := range line {
            if char == 'S' {
                startRow = r
                startCol = c
            }
            row = append(row, char)
        }
        
        pipes = append(pipes, row)
    }

    pipeMap := map[Point]bool{}
    pipeFind(pipes, startRow, startCol, startRow, startCol, pipeMap)
    convertPipe(pipes, startRow, startCol, pipeMap)
    res := countInside(pipes, pipeMap)

    fmt.Println("solution", res)
}

func pipeFind(pipes [][]rune, row int, col int, prevRow int, prevCol int, pipeMap map[Point]bool) {
    var currChar rune = pipes[row][col]
    var nextChar rune
    var nextRow int
    var nextCol int

    pipeMap[Point{row, col}] = true

    for _, dir := range directions {
        if !slices.Contains(dir.connectees, currChar) && currChar != 'S' { continue }

        nextRow = row + dir.dRow
        nextCol = col + dir.dCol

        if nextRow == prevRow && nextCol == prevCol { continue }
        if nextRow < 0 || nextRow >= len(pipes) { continue }
        if nextCol < 0 || nextCol >= len(pipes[nextRow]) { continue }

        nextChar = pipes[nextRow][nextCol]

        if !slices.Contains(dir.connectors, nextChar) && nextChar != 'S' { continue }

        break
    }

    if nextChar == 'S' {
        return 
    }

    pipeFind(pipes, nextRow, nextCol, row, col, pipeMap)
}

func convertPipe(pipes [][]rune, row int, col int, pipeMap map[Point]bool) {
    up := false
    down := false
    left := false
    right := false

    for _, dir := range directions {
        newR := row + dir.dRow
        newC := col + dir.dCol
        if _, ok := pipeMap[Point{newR, newC}]; !ok { continue }
        newChar := pipes[newR][newC]
        if !slices.Contains(dir.connectors, newChar) { continue }

        if dir.dRow == -1 && dir.dCol == 0 { up = true }
        if dir.dRow == 1 && dir.dCol == 0 { down = true }
        if dir.dRow == 0 && dir.dCol == -1 { left = true }
        if dir.dRow == 0 && dir.dCol == 1 { right = true }
    }

    if up && down {
        pipes[row][col] = '|'
    }
    if left && right {
        pipes[row][col] = '-'
    }
    if up && right {
        pipes[row][col] = 'L'
    }
    if up && left {
        pipes[row][col] = 'J'
    }
    if down && right {
        pipes[row][col] = 'F'
    }
    if down && left {
        pipes[row][col] = '7'
    }
}

func countInside(pipes [][]rune, pipeMap map[Point]bool) int {
    for r, row := range pipes {
        for c := range row {
            if _, ok := pipeMap[Point{r, c}]; !ok {
                pipes[r][c] = '.'
            }
        }
    }

    for i := 0; i < len(pipes); i++ {
        for j := 0; j < len(pipes[i]); j++ {
        }
    }

    count := 0
    inside := false
    for i := 0; i < len(pipes); i++ {
        for j := 0; j < len(pipes[i]); j++ {
            if pipes[i][j] == '|' {
                inside = !inside
            } else if pipes[i][j] == '.' && inside {
                count++
            } else if pipes[i][j] == 'F' {
                for {
                    j++

                    if pipes[i][j] == '-' {
                        continue
                    } else if pipes[i][j] == 'J' {
                        inside = !inside
                        break
                    } else {
                        break
                    }
                }
            } else if pipes[i][j] == 'L' {
                for {
                    j++

                    if pipes[i][j] == '-' {
                        continue
                    } else if pipes[i][j] == '7' {
                        inside = !inside
                        break
                    } else {
                        break
                    }
                }
            }
        }
    }

    return count
}

