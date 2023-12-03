package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "unicode"
)

func main() {
    inputFilename := "gear_input"
    part1(inputFilename)
    part2(inputFilename)
}

func part1(inputFilename string) {
    res := 0

    readFile, err := os.Open(inputFilename)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    grid := make([][]rune, 0)

    for scanner.Scan() {
        text := scanner.Text()

        row := make([]rune, 0) 
        for _, c := range text {
            row = append(row, c)
        }

        grid = append(grid, row)
    }

    res = sumAdjacentNumbers(grid)

    readFile.Close()

    fmt.Println("Solution", res)
}

func sumAdjacentNumbers(grid [][]rune) int {
    res := 0

    for y, row := range grid {
        for x, c := range row {
            if c == '.' {
                continue
            } else if unicode.IsNumber(c) {
                continue
            } else {
                res += sumAdjacentNumbersForPoint(grid, x, y)
            }
        }
    }

    return res
}

func sumAdjacentNumbersForPoint(grid [][]rune, x int, y int) int {
    res := 0

    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            if dy == 0 && dx == 0 {
                continue
            }

            if y + dy < 0 || y + dy >= len(grid) {
                continue
            }

            if x + dx < 0 || x + dx >= len(grid[y]) {
                continue
            }

            c := grid[y + dy][x + dx]
            if unicode.IsNumber(c) {
                res += findFullNumber(grid, x + dx, y + dy)
            }
        }
    }

    return res
}

func findFullNumber(grid [][]rune, x int, y int) int {
    xStart := x
    xEnd := x

    for xStart >= 0 && unicode.IsNumber(grid[y][xStart]) {
        xStart--
    }

    for xEnd <= len(grid[y]) - 1 && unicode.IsNumber(grid[y][xEnd]) {
        xEnd++
    }

    number := string(grid[y][xStart+1:xEnd])
    res, err := strconv.Atoi(number)

    for i := xStart+1; i < xEnd; i++ {
        grid[y][i] = '.'
    }

    if err != nil {
        panic(err)
    }

    return res
}

func part2(inputFilename string) {
    res := 0

    readFile, err := os.Open(inputFilename)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    grid := make([][]rune, 0)

    for scanner.Scan() {
        text := scanner.Text()

        row := make([]rune, 0) 
        for _, c := range text {
            row = append(row, c)
        }

        grid = append(grid, row)
    }

    res = sumGears(grid)

    readFile.Close()

    fmt.Println("Solution", res)
}

func sumGears(grid [][]rune) int {
    res := 0

    for y, row := range grid {
        for x, c := range row {
            if c == '*' {
                res += sumGearAtPoint(grid, x, y)
            }
        }
    }

    return res
}

func sumGearAtPoint(grid [][]rune, x int, y int) int {
    count := 0
    number1 := 0
    number2 := 0

    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            if dy == 0 && dx == 0 {
                continue
            }

            if y + dy < 0 || y + dy >= len(grid) {
                continue
            }

            if x + dx < 0 || x + dx >= len(grid[y]) {
                continue
            }

            c := grid[y + dy][x + dx]
            if unicode.IsNumber(c) {
                count += 1
                if count == 1 {
                    number1 = findFullNumber(grid, x + dx, y + dy)
                } else if count == 2 {
                    number2 = findFullNumber(grid, x + dx, y + dy)
                } else {
                    return 0
                }
            }
        }
    }

    if count == 2 {
        return number1 * number2
    } else {
        return 0
    }
}

