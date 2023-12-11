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

type Point struct {
    r int
    c int
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    colLength := len(lines)
    rowLength := len(lines[0])

    rowsOccupied := make([]bool, colLength)
    colsOccupied := make([]bool, rowLength)
    for r, line := range lines {
        for c := range line {
            if line[c] == '#' {
                rowsOccupied[r] = true
                colsOccupied[c] = true
            }
        }
    }

    expanded := make([]string, 0)
    for r, line := range lines {
        var row strings.Builder

        for c, char := range line {
            row.WriteRune(char)
            if !colsOccupied[c] {
                row.WriteRune(char)
            }
        }

        expanded = append(expanded, row.String())
        if !rowsOccupied[r] {
            expanded = append(expanded, row.String())
        }
    }

    galaxies := make([]*Point, 0)
    for r, line := range expanded {
        for c, char := range line {
            if char == '#' {
                galaxies = append(galaxies, &Point{r, c})
            }
        }
    }

    res := 0
    for i := 0; i < len(galaxies); i++ {
        for j := i + 1; j < len(galaxies); j++ {
            dr := galaxies[i].r - galaxies[j].r
            dc := galaxies[i].c - galaxies[j].c
            if dr < 0 { dr = -dr }
            if dc < 0 { dc = -dc }
            res += dr + dc
        }
    }
    fmt.Println("solution", res)
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    colLength := len(lines)
    rowLength := len(lines[0])

    rowsOccupied := make([]bool, colLength)
    colsOccupied := make([]bool, rowLength)
    for r, line := range lines {
        for c := range line {
            if line[c] == '#' {
                rowsOccupied[r] = true
                colsOccupied[c] = true
            }
        }
    }

    expanded := make([]string, 0)
    for r, line := range lines {
        var row strings.Builder

        for c, char := range line {
            if !colsOccupied[c] {
                row.WriteRune('m')
            } else {
                row.WriteRune(char)
            }
        }

        if !rowsOccupied[r] {
            expanded = append(expanded, strings.Repeat("m", rowLength))
        } else {
            expanded = append(expanded, row.String())
        }
    }

    galaxies := make([]*Point, 0)
    for r, line := range expanded {
        for c, char := range line {
            if char == '#' {
                galaxies = append(galaxies, &Point{r, c})
            }
        }
    }

    scale := 1000000
    res := 0
    for i := 0; i < len(galaxies); i++ {
        for j := i + 1; j < len(galaxies); j++ {

            smallRow := min(galaxies[i].r, galaxies[j].r)
            largeRow := max(galaxies[i].r, galaxies[j].r)
            smallCol := min(galaxies[i].c, galaxies[j].c)
            largeCol := max(galaxies[i].c, galaxies[j].c)

            dr := 0
            for r := smallRow+1; r <= largeRow; r++ {
                if expanded[r][smallCol] == 'm' {
                    dr += scale
                } else {
                    dr += 1
                }
            }

            dc := 0
            for c := smallCol+1; c <= largeCol; c++ {
                if expanded[smallRow][c] == 'm' {
                    dc += scale
                } else {
                    dc += 1
                }
            }

            res += dr + dc
        }
    }
    fmt.Println("solution", res)
}

