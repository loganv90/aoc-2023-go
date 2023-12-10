package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    part1("input1")
    part1("input")
    part2("input1")
    part2("input")
}

func part1(inputFilename string) {
    readFile, err := os.Open(inputFilename)
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    scanner.Scan()
    seedsText := scanner.Text()
    seedsSplitText := strings.FieldsFunc(seedsText, func(r rune) bool { return r == ' ' })
    seeds := make([]int, 0)
    tmpSeeds := make([]int, 0)

    for i := 1; i < len(seedsSplitText); i++ {
        seed, err := strconv.Atoi(seedsSplitText[i])
        if err != nil {
            panic(err)
        }
        seeds = append(seeds, seed)
        tmpSeeds = append(tmpSeeds, seed)
    }

    for scanner.Scan() {
        text := scanner.Text()

        if strings.Contains(text, ":") {
            continue
        }

        if strings.TrimSpace(text) == "" {
            for i := 0; i < len(seeds); i++ {
                seeds[i] = tmpSeeds[i]
            }
            continue
        }

        textSplit := strings.FieldsFunc(text, func(r rune) bool { return r == ' ' })
        destinationStart, _ := strconv.Atoi(textSplit[0])
        sourceStart, _ := strconv.Atoi(textSplit[1])
        rangeLength, _ := strconv.Atoi(textSplit[2])

        for i := 0; i < len(seeds); i++ {
            if seeds[i] >= sourceStart && seeds[i] < sourceStart + rangeLength {
                tmpSeeds[i] = seeds[i] - sourceStart + destinationStart
            }
        }
    }

    res := tmpSeeds[0]
    for i := 1; i < len(tmpSeeds); i++ {
        if tmpSeeds[i] < res {
            res = tmpSeeds[i]
        }
    }

    readFile.Close()
    fmt.Println("solution", res)
}

type Range struct {
    start int
    length int
}

type Mapping struct {
    destinationStart int
    sourceStart int
    length int
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

func getSeeds(line string) []Range {
    seeds := make([]Range, 0)

    seedsSplitText := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
    for i := 1; i < len(seedsSplitText) - 1; i += 2 {
        start, err := strconv.Atoi(seedsSplitText[i])
        if err != nil { panic(err) }
        length, err := strconv.Atoi(seedsSplitText[i + 1])
        if err != nil { panic(err) }
        
        seeds = append(seeds, Range{start, length})
    }

    return seeds
}

func getLevels(lines []string) [][]Mapping {
    levels := make([][]Mapping, 0)

    for i := 0; i < len(lines); i++ {
        if strings.Contains(lines[i], ":") {
            levels = append(levels, make([]Mapping, 0))
            continue
        }

        if strings.TrimSpace(lines[i]) == "" {
            continue
        }

        mappingSplitText := strings.FieldsFunc(lines[i], func(r rune) bool { return r == ' ' })
        destinationStart, err := strconv.Atoi(mappingSplitText[0])
        if err != nil { panic(err) }
        sourceStart, err := strconv.Atoi(mappingSplitText[1])
        if err != nil { panic(err) }
        rangeLength, err := strconv.Atoi(mappingSplitText[2])
        if err != nil { panic(err) }

        levels[len(levels) - 1] = append(levels[len(levels) - 1], Mapping{destinationStart, sourceStart, rangeLength})
    }

    return levels
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)
    seeds := getSeeds(lines[0])
    levels := getLevels(lines[1:])

    i := 0
    for {
        current := i

        for j := len(levels) - 1; j >= 0; j-- {
            for k := 0; k < len(levels[j]); k++ {
                if levels[j][k].destinationStart <= current && levels[j][k].destinationStart + levels[j][k].length > current {
                    current = levels[j][k].sourceStart + current - levels[j][k].destinationStart
                    break
                }
            }
        }

        for j := 0; j < len(seeds); j++ {
            if seeds[j].start <= current && seeds[j].start + seeds[j].length > current {
                fmt.Println("solution", i)
                return
            }
        }

        i++
    }
}

