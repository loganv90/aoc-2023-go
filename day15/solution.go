package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
)

func main() {
    part1("input1")
    part1("input2")
    part1("input")
    part2("input2")
    part2("input")
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    lines[0] = strings.ReplaceAll(lines[0], "\n", "")
    commands := strings.Split(lines[0], ",")

    res := 0
    for _, command := range commands {
        res += hashAlgorithm(command)
    }

    fmt.Println("solution", res)
}

func hashAlgorithm(input string) int {
    current := 0

    for _, c := range input {
        value := int(c)
        current = (current + value) * 17
        current %= 256
    }

    return current
}

type Box struct {
    lense int
    position int
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    lines[0] = strings.ReplaceAll(lines[0], "\n", "")
    commands := strings.Split(lines[0], ",")

    boxes := make([]map[string]*Box, 256)
    for i := range boxes {
        boxes[i] = map[string]*Box{}
    }

    for _, command := range commands {
        label, add, addNumber := parseCommand(command)
        hash := hashAlgorithm(label)

        if add {
            boxMap := boxes[hash]
            box, ok := boxMap[label]

            if ok {
                box.lense = addNumber
                continue
            }

            boxMapLen := len(boxMap)
            box = &Box{addNumber, boxMapLen}
            boxMap[label] = box
        } else {
            boxMap := boxes[hash]
            box, ok := boxMap[label]

            if !ok {
                continue
            }

            position := box.position

            for _, value := range boxMap {
                if value.position > position {
                    value.position--
                }
            }

            delete(boxMap, label)
        }
    }

    res := 0
    for i, boxMap := range boxes {
        for _, box := range boxMap {
            res += box.lense * (i + 1) * (box.position + 1)
        }
    }

    fmt.Println("solution", res)
}

func parseCommand(input string) (string, bool, int) {
    index := strings.Index(input, "=")
    if index == -1 {
        index = strings.Index(input, "-")
        if index == -1 {
            panic("can't parse command")
        }
    }
    add := input[index] == '='
    label := input[:index]
    addNumber := -1

    if add {
        number, err := strconv.Atoi(input[index+1:])
        if err != nil {
            panic("can't parse number")
        }
        addNumber = number
    }

    return label, add, addNumber
}

