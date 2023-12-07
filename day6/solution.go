package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
    "unicode"
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

    time_strings := strings.FieldsFunc(lines[0], func(r rune) bool { return r == ' ' })[1:]
    distance_strings := strings.FieldsFunc(lines[1], func(r rune) bool { return r == ' ' })[1:]

    times := make([]int, len(time_strings))
    distances := make([]int, len(time_strings))
    ways := make([]int, len(time_strings))
    for i := 0; i < len(time_strings); i++ {
        time, err := strconv.Atoi(time_strings[i])
        if err != nil { panic(err) }
        times[i] = time

        distance, err := strconv.Atoi(distance_strings[i])
        if err != nil { panic(err) }
        distances[i] = distance

        ways[i] = 0
    }

    // oh ya ya brute force
    for i := 0; i < len(times); i++ {
        for j := 0; j <= times[i]; j++ {
            runTime := times[i] - j
            distance := j * runTime

            if distance > distances[i] {
                ways[i] += 1
            }
        }
    }

    res := 1
    for _, way := range ways {
        res *= way
    }

    fmt.Println("solution", res)
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    time_string := ""
    distance_string := ""

    for _, c := range lines[0] {
        if unicode.IsNumber(c) {
            time_string += string(c)
        }
    }

    for _, c := range lines[1] {
        if unicode.IsNumber(c) {
            distance_string += string(c)
        }
    }

    time, err := strconv.Atoi(time_string)
    if err != nil { panic(err) }
    distance, err := strconv.Atoi(distance_string)
    if err != nil { panic(err) }

    // brute force works but could do binary search here
    res := 0
    for i := 0; i <= time; i++ {
        myTime := time - i
        myDistance := i * myTime

        if myDistance > distance {
            res += 1
        }
    }

    fmt.Println("solution", res)
}

