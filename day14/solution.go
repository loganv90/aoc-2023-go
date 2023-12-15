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

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    data := make([][]string, 0)
    for _, line := range lines {
        if len(line) == 0 { continue }

        row := make([]string, 0)
        for _, char := range line {
            row = append(row, string(char))
        }
        data = append(data, row)
    }

    for c := 0; c < len(data[0]); c++ {
        index := 0
        for r := 0; r < len(data); r++ {
            if data[r][c] == "#" {
                index = r + 1
            } else if data[r][c] == "O" {
                data[r][c] = "."
                data[index][c] = "O"
                index += 1
            } else {
                continue
            }
        }
    }

    res := 0
    for i, row := range data {
        for _, char := range row {
            if char == "O" {
                res += len(data) - i
            }
        }
    }

    fmt.Println("solution", res)
}

func rollNorth(data [][]string) {
    for c := 0; c < len(data[0]); c++ {
        index := 0
        for r := 0; r < len(data); r++ {
            if data[r][c] == "#" {
                index = r + 1
            } else if data[r][c] == "O" {
                data[r][c] = "."
                data[index][c] = "O"
                index += 1
            } else {
                continue
            }
        }
    }
}

func rollEast(data [][]string) {
    for r := 0; r < len(data); r++ {
        index := len(data[0]) - 1
        for c := len(data[0]) - 1; c >= 0; c-- {
            if data[r][c] == "#" {
                index = c - 1
            } else if data[r][c] == "O" {
                data[r][c] = "."
                data[r][index] = "O"
                index -= 1
            } else {
                continue
            }
        }
    }
}

func rollSouth(data [][]string) {
    for c := 0; c < len(data[0]); c++ {
        index := len(data) - 1
        for r := len(data) - 1; r >= 0; r-- {
            if data[r][c] == "#" {
                index = r - 1
            } else if data[r][c] == "O" {
                data[r][c] = "."
                data[index][c] = "O"
                index -= 1
            } else {
                continue
            }
        }
    }
}

func rollWest(data [][]string) {
    for r := 0; r < len(data); r++ {
        index := 0
        for c := 0; c < len(data[0]); c++ {
            if data[r][c] == "#" {
                index = c + 1
            } else if data[r][c] == "O" {
                data[r][c] = "."
                data[r][index] = "O"
                index += 1
            } else {
                continue
            }
        }
    }
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    data := make([][]string, 0)
    for _, line := range lines {
        if len(line) == 0 { continue }

        row := make([]string, 0)
        for _, char := range line {
            row = append(row, string(char))
        }
        data = append(data, row)
    }

    saves := make([][][]string, 0)
    exitLoop := false
    for !exitLoop {
        rollNorth(data)
        rollWest(data)
        rollSouth(data)
        rollEast(data)

        for i, save := range saves {
            if sameData(data, save) {
                exitLoop = true
                dataIndex := (999999999 - i) % (len(saves) - i)
                data = saves[i + dataIndex]
                break
            }
        }

        saves = append(saves, copyData(data))
    }

    res := 0
    for i, row := range data {
        for _, char := range row {
            if char == "O" {
                res += len(data) - i
            }
        }
    }

    fmt.Println("solution", res)
}

func sameData(data1 [][]string, data2 [][]string) bool {
    for r := 0; r < len(data1); r++ {
        for c := 0; c < len(data1[0]); c++ {
            if data1[r][c] != data2[r][c] {
                return false
            }
        }
    }

    return true
}

func copyData(data [][]string) [][]string {
    res := make([][]string, 0)
    for r := 0; r < len(data); r++ {
        row := make([]string, 0)
        for c := 0; c < len(data[0]); c++ {
            row = append(row, data[r][c])
        }
        res = append(res, row)
    }

    return res
}

