package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
    "regexp"
)

func main() {
    part1("input1")
    part1("input")
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

type Filter struct {
    name string
    value int
    less bool
    success string
}

func inputToFiltersAndParts(lines []string) (map[string][]*Filter, []map[string]int) {
    filterListMap := map[string][]*Filter{}
    partMapList := []map[string]int{}

    lineIndex := 0
    r := regexp.MustCompile(`(.*)\{(.*)\}`)
    for ; lineIndex < len(lines); lineIndex++ {
        currentLine := lines[lineIndex]
        if currentLine == "" {
            lineIndex++
            break
        }

        matches := r.FindStringSubmatch(currentLine)
        name := matches[1]
        filterSplit := strings.Split(matches[2], ",")
        filterList := []*Filter{}
        for _, filterString := range filterSplit {
            filterStringSplit := strings.Split(filterString, ":")

            if len(filterStringSplit) < 2 {
                filter := &Filter{
                    name: "",
                    value: 0,
                    less: false,
                    success: filterString,
                }
                filterList = append(filterList, filter)
                continue
            }

            less := filterStringSplit[0][1] == '<'
            name := string(filterStringSplit[0][0])
            value, err := strconv.Atoi(filterStringSplit[0][2:])
            if err != nil { panic(err) }
            success := filterStringSplit[1]

            filter := &Filter{
                name: name,
                value: value,
                less: less,
                success: success,
            }
            filterList = append(filterList, filter)
        }

        filterListMap[name] = filterList
    }

    for ; lineIndex < len(lines); lineIndex++ {
        currentLine := lines[lineIndex]
        if currentLine == "" {
            lineIndex++
            break
        }
        currentLine = currentLine[1:len(lines[lineIndex])-1]

        part := map[string]int{}
        splitLine := strings.Split(currentLine, ",")
        for _, valueString := range splitLine {
            splitValueString := strings.Split(valueString, "=")

            name := splitValueString[0]
            value, err := strconv.Atoi(splitValueString[1])
            if err != nil { panic(err) }

            part[name] = value
        }

        partMapList = append(partMapList, part)
    }

    return filterListMap, partMapList
}

func findEndString(filterListMap map[string][]*Filter, partMap map[string]int) string {
    currentName := "in"

    for {
        filterList, ok := filterListMap[currentName]
        if !ok {
            return currentName
        }

        for _, filter := range filterList {
            partValue, ok := partMap[filter.name]
            if !ok {
                currentName = filter.success
                break
            }

            if filter.less {
                if partValue < filter.value {
                    currentName = filter.success
                    break
                }
            } else {
                if partValue > filter.value {
                    currentName = filter.success
                    break
                }
            }
        }
    }
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)
    filterListMap, partMapList := inputToFiltersAndParts(lines)

    res := 0
    for _, partMap := range partMapList {
        endString := findEndString(filterListMap, partMap)
        if endString == "A" {
            for _, value := range partMap {
                res += value
            }
        } else if endString == "R" {
            // do nothing
        } else {
            panic("unknown end string")
        }
    }

    fmt.Println("solution", res)
}

