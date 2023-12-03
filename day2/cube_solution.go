package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    inputFilename := "cube_input"
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

    for scanner.Scan() {
        text := scanner.Text()

        splitText := strings.FieldsFunc(text, func(r rune) bool { return r == ':' || r == ';' })

        gameInfo := strings.FieldsFunc(strings.TrimSpace(splitText[0]), func(r rune) bool { return r == ' ' })
        gameNumber, err := strconv.Atoi(gameInfo[1])
        if err != nil {
            panic(err)
        }

        for i := 1; i < len(splitText); i++ {
            colors := strings.FieldsFunc(splitText[i], func(r rune) bool { return r == ',' })
            requiredRed := 0
            requiredGreen := 0
            requiredBlue := 0

            for j := 0; j < len(colors); j++ {
                color := strings.FieldsFunc(strings.TrimSpace(colors[j]), func(r rune) bool { return r == ' ' })
                colorNumber, err := strconv.Atoi(color[0])
                colorName := color[1]
                if err != nil {
                    panic(err)
                }

                if colorName == "red" {
                    requiredRed = max(requiredRed, colorNumber)
                } else if colorName == "green" {
                    requiredGreen = max(requiredGreen, colorNumber)
                } else if colorName == "blue" {
                    requiredBlue = max(requiredBlue, colorNumber)
                }
            }

            if requiredRed > 12 || requiredGreen > 13 || requiredBlue > 14 {
                gameNumber = 0
                break
            }
        }

        res += gameNumber
    }

    readFile.Close()

    fmt.Println("Solution", res)
}

func part2(inputFilename string) {
    res := 0

    readFile, err := os.Open(inputFilename)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        text := scanner.Text()

        splitText := strings.FieldsFunc(text, func(r rune) bool { return r == ':' || r == ';' })

        requiredRed := 0
        requiredGreen := 0
        requiredBlue := 0

        for i := 1; i < len(splitText); i++ {
            colors := strings.FieldsFunc(splitText[i], func(r rune) bool { return r == ',' })

            for j := 0; j < len(colors); j++ {
                color := strings.FieldsFunc(strings.TrimSpace(colors[j]), func(r rune) bool { return r == ' ' })
                colorNumber, err := strconv.Atoi(color[0])
                colorName := color[1]
                if err != nil {
                    panic(err)
                }

                if colorName == "red" {
                    requiredRed = max(requiredRed, colorNumber)
                } else if colorName == "green" {
                    requiredGreen = max(requiredGreen, colorNumber)
                } else if colorName == "blue" {
                    requiredBlue = max(requiredBlue, colorNumber)
                }
            }
        }

        res += requiredRed * requiredGreen * requiredBlue
    }

    readFile.Close()

    fmt.Println("Solution", res)
}
