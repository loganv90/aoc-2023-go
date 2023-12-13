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

type Mirror struct {
    rows []string
    cols []string
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    mirrors := make([]*Mirror, 0)
    currentMirror := &Mirror{make([]string, 0), make([]string, 0)}
    for _, line := range lines {
        if len(line) <= 0 && len(currentMirror.rows) > 0 {
            mirrors = append(mirrors, currentMirror)
            currentMirror = &Mirror{make([]string, 0), make([]string, 0)}
        } else if len(line) <= 0 {
            continue
        } else {
            currentMirror.rows = append(currentMirror.rows, line)
        }
    }

    for _, mirror := range mirrors {
        for i := 0; i < len(mirror.rows[0]); i++ {
            col := ""
            for j := 0; j < len(mirror.rows); j++ {
                col += string(mirror.rows[j][i])
            }
            mirror.cols = append(mirror.cols, col)
        }
    }

    res := 0
    for _, mirror := range mirrors {
        index, err := findMirrorIndex(mirror.rows)
        if err == nil {
            res += index * 100
            continue
        }

        index, err = findMirrorIndex(mirror.cols)
        if err == nil {
            res += index
            continue
        }

        panic("no mirror found")
    }

    fmt.Println("solution", res)
}

func findMirrorIndex(data []string) (int, error) {
    for i := 0; i < len(data)-1; i++ {

        left := i
        right := i+1
        for {
            if left < 0 || right >= len(data) {
                return i+1, nil
            }

            if data[left] != data[right] {
                break
            }

            left--
            right++
        }
    }

    return -1, fmt.Errorf("no mirror found")
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    mirrors := make([]*Mirror, 0)
    currentMirror := &Mirror{make([]string, 0), make([]string, 0)}
    for _, line := range lines {
        if len(line) <= 0 && len(currentMirror.rows) > 0 {
            mirrors = append(mirrors, currentMirror)
            currentMirror = &Mirror{make([]string, 0), make([]string, 0)}
        } else if len(line) <= 0 {
            continue
        } else {
            currentMirror.rows = append(currentMirror.rows, line)
        }
    }

    for _, mirror := range mirrors {
        for i := 0; i < len(mirror.rows[0]); i++ {
            col := ""
            for j := 0; j < len(mirror.rows); j++ {
                col += string(mirror.rows[j][i])
            }
            mirror.cols = append(mirror.cols, col)
        }
    }

    res := 0
    for _, mirror := range mirrors {
        originalRows, err := findMirrorIndex(mirror.rows)
        originalCols, err := findMirrorIndex(mirror.cols)


        index, err := findMirrorIndexWithTolerance(mirror.rows, originalRows)
        if err == nil {
            res += index * 100
            continue
        }

        index, err = findMirrorIndexWithTolerance(mirror.cols, originalCols)
        if err == nil {
            res += index
            continue
        }

        panic("no mirror found")
    }

    fmt.Println("solution", res)
}

func findMirrorIndexWithTolerance(data []string, original int) (int, error) {
    for i := 0; i < len(data)-1; i++ {
        limit := 1

        left := i
        right := i+1
        for {
            if left < 0 || right >= len(data) {
                if i+1 != original {
                    return i+1, nil
                } else {
                    break
                }
            }

            diff := stringDifference(data[left], data[right])
            limit -= diff
            if limit < 0 {
                break
            }

            left--
            right++
        }
    }

    return -1, fmt.Errorf("no mirror found")
}

func stringDifference(a string, b string) int {
    res := 0

    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            res++
        }
    }

    return res
}

