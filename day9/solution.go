package main

import (
    "strconv"
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

    res := 0
    for _, line := range lines {
        if len(line) == 0 { continue }
        stringHistory := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })

        history := make([]int, 0)
        for _, v := range stringHistory {
            h, err := strconv.Atoi(v)
            if err != nil { panic(err) }
            history = append(history, h)
        }

        res += findNext(history)
    }

    fmt.Println("solution", res)
}

func findNext(history []int) int {
    allZero := true
    for i := 0; i < len(history); i++ {
        if history[i] != 0 {
            allZero = false
            break
        }
    }
    if allZero { return 0 }

    nextIteration := make([]int, 0)
    for i := 0; i < len(history) - 1; i++ {
        nextIteration = append(nextIteration, history[i+1] - history[i])
    }

    return findNext(nextIteration) + history[len(history) - 1]
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    res := 0
    for _, line := range lines {
        if len(line) == 0 { continue }
        stringHistory := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })

        history := make([]int, 0)
        for _, v := range stringHistory {
            h, err := strconv.Atoi(v)
            if err != nil { panic(err) }
            history = append(history, h)
        }

        res += findPrev(history)
    }

    fmt.Println("solution", res)
}

func findPrev(history []int) int {
    allZero := true
    for i := 0; i < len(history); i++ {
        if history[i] != 0 {
            allZero = false
            break
        }
    }
    if allZero { return 0 }

    nextIteration := make([]int, 0)
    for i := 0; i < len(history) - 1; i++ {
        nextIteration = append(nextIteration, history[i+1] - history[i])
    }

    return history[0] - findPrev(nextIteration)
}

