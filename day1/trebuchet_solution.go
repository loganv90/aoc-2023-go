package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    inputFilename := "trebuchet_input"
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

        left := 0
        for i := 0; i < len(text); i++ {
            if left, err = strconv.Atoi(string(text[i])); err == nil {
                break
            }
        }

        right := 0
        for i := len(text) - 1; i >= 0; i-- {
            if right, err = strconv.Atoi(string(text[i])); err == nil {
                break
            }
        }

        res += left*10 + right
    }

    readFile.Close()

    fmt.Println("Solution", res)
}

func part2(inputFilename string) {
    numberWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    reverseWords := []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}
    progressWords := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

    incrementProgress := func(words []string, letter byte) int {
        for i := 0; i < len(words); i++ {
            textLetter := string(letter)
            wordLetter := string(words[i][progressWords[i]])

            if textLetter == wordLetter {
                progressWords[i]++
                if progressWords[i] >= len(words[i]) {
                    return i + 1
                }
            }
        }

        return 0
    }

    res := 0

    readFile, err := os.Open(inputFilename)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        text := scanner.Text()

        progressWords = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
        left := 0
        for i := 0; i < len(text); i++ {
            if left, err = strconv.Atoi(string(text[i])); err == nil {
                break
            }

            if left = incrementProgress(numberWords, text[i]); left > 0 {
                break
            }
        }

        progressWords = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
        right := 0
        for i := len(text) - 1; i >= 0; i-- {
            if right, err = strconv.Atoi(string(text[i])); err == nil {
                break
            }

            if right = incrementProgress(reverseWords, text[i]); right > 0 {
                break
            }
        }

        res += left*10 + right
    }

    readFile.Close()

    fmt.Println("Solution", res)
}

