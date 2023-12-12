package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
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
        if len(line) == 0 {
            continue
        }
        splitLine := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' } )
        splitNumberStrings := strings.FieldsFunc(splitLine[1], func(r rune) bool { return r == ',' } )
        splitNumbers := make([]int, len(splitNumberStrings))
        for i, s := range splitNumberStrings {
            num, err := strconv.Atoi(s)
            if err != nil { panic(err) }
            splitNumbers[i] = num
        }
        res += countArrangements(splitLine[0], splitNumbers)
    }

    fmt.Println("solution", res)
}

func countArrangements(data string, numbers []int) int {
    if len(data) == 0 && len(numbers) == 0 {
        return 1
    }

    res := 0
    if len(data) > 0 {
        if data[0] == '.' {
            res += countArrangements(data[1:], numbers)
        } else if data[0] == '#' {
            if len(numbers) > 0 && isValid(data[1:], numbers[0]-1) {
                dataIndex := min(len(data), 1+numbers[0])
                res += countArrangements(data[dataIndex:], numbers[1:])
            }
        } else if data[0] == '?' {
            if len(numbers) > 0 && isValid(data[1:], numbers[0]-1) {
                dataIndex := min(len(data), 1+numbers[0])
                res += countArrangements(data[dataIndex:], numbers[1:]) // case where ? is #
            }
            res += countArrangements(data[1:], numbers) // case where ? is .
        }
    }

    return res
}

func isValid(data string, number int) bool {
    if len(data) < number {
        return false
    }

    for i := 0; i < number; i++ {
        if data[i] == '.' {
            return false
        }
    }

    if len(data) == number {
        return true
    }

    if data[number] == '.' || data[number] == '?' {
        return true
    }

    return false
}

type MemoKey struct {
    dataLen int
    numbersLen int
}

var memo = make(map[MemoKey]int)

func part2(inputFilename string) { // dynamic programming on the length of the array
    lines := getLines(inputFilename)

    res := 0
    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        splitLine := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' } )
        splitNumberStrings := strings.FieldsFunc(splitLine[1], func(r rune) bool { return r == ',' } )
        splitNumbers := make([]int, len(splitNumberStrings))
        for i, s := range splitNumberStrings {
            num, err := strconv.Atoi(s)
            if err != nil { panic(err) }
            splitNumbers[i] = num
        }

        unfoldedSplitLine := ""
        unfoldedSplitNumbers := make([]int, 0)
        for i := 0; i < 5; i++ {
            unfoldedSplitLine += "?"
            unfoldedSplitLine += splitLine[0]
            unfoldedSplitNumbers = append(unfoldedSplitNumbers, splitNumbers...)
        }

        memo = make(map[MemoKey]int)
        res += countArrangements2(unfoldedSplitLine[1:], unfoldedSplitNumbers)
    }

    fmt.Println("solution", res)
}

func countArrangements2(data string, numbers []int) int {
    dataLen := len(data)
    numbersLen := len(numbers)
    
    if dataLen == 0 && numbersLen == 0 {
        return 1
    }

    memoKey := MemoKey{dataLen, numbersLen}
    if val, ok := memo[memoKey]; ok {
        return val
    }

    res := 0
    if len(data) > 0 {
        if data[0] == '.' {
            res += countArrangements2(data[1:], numbers)
        } else if data[0] == '#' {
            if len(numbers) > 0 && isValid(data[1:], numbers[0]-1) {
                dataIndex := min(len(data), 1+numbers[0])
                res += countArrangements2(data[dataIndex:], numbers[1:])
            }
        } else if data[0] == '?' {
            if len(numbers) > 0 && isValid(data[1:], numbers[0]-1) {
                dataIndex := min(len(data), 1+numbers[0])
                res += countArrangements2(data[dataIndex:], numbers[1:]) // case where ? is #
            }
            res += countArrangements2(data[1:], numbers) // case where ? is .
        }
    }

    memo[memoKey] = res
    return res
}

