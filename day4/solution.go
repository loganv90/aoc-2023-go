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
    res := 0

    readFile, err := os.Open(inputFilename)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        text := scanner.Text()

        cardScore := 0

        splitText := strings.FieldsFunc(text, func(r rune) bool { return r == ':' || r == '|' })
        winningNumbers := strings.FieldsFunc(strings.TrimSpace(splitText[1]), func(r rune) bool { return r == ' ' })
        playingNumbers := strings.FieldsFunc(strings.TrimSpace(splitText[2]), func(r rune) bool { return r == ' ' })

        var winningNumberMap = make(map[int]bool)

        for i := 0; i < len(winningNumbers); i++ {
            winningNumber, err := strconv.Atoi(winningNumbers[i])
            if err != nil {
                panic(err)
            }

            winningNumberMap[winningNumber] = true
        }

        for i := 0; i < len(playingNumbers); i++ {
            playingNumber, err := strconv.Atoi(playingNumbers[i])
            if err != nil {
                panic(err)
            }

            if _, ok := winningNumberMap[playingNumber]; ok {
                if cardScore == 0 {
                    cardScore = 1
                } else {
                    cardScore *= 2
                }
            }
        }

        res += cardScore
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

    cardScoreList := make([]int, 0)
    cardCountList := make([]int, 0)

    for scanner.Scan() {
        text := scanner.Text()

        cardScore := 0

        splitText := strings.FieldsFunc(text, func(r rune) bool { return r == ':' || r == '|' })
        winningNumbers := strings.FieldsFunc(strings.TrimSpace(splitText[1]), func(r rune) bool { return r == ' ' })
        playingNumbers := strings.FieldsFunc(strings.TrimSpace(splitText[2]), func(r rune) bool { return r == ' ' })

        var winningNumberMap = make(map[int]bool)

        for i := 0; i < len(winningNumbers); i++ {
            winningNumber, err := strconv.Atoi(winningNumbers[i])
            if err != nil {
                panic(err)
            }

            winningNumberMap[winningNumber] = true
        }

        for i := 0; i < len(playingNumbers); i++ {
            playingNumber, err := strconv.Atoi(playingNumbers[i])
            if err != nil {
                panic(err)
            }

            if _, ok := winningNumberMap[playingNumber]; ok {
                cardScore += 1
            }
        }

        cardScoreList = append(cardScoreList, cardScore)
        cardCountList = append(cardCountList, 1)
    }

    for i := 0; i < len(cardScoreList); i++ {
        res += cardCountList[i]

        for j := i + 1; j <= i + cardScoreList[i]; j++ {
            if j < len(cardCountList) {
                cardCountList[j] += cardCountList[i]
            }
        }
    }

    readFile.Close()

    fmt.Println("Solution", res)
}
