package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
)

func main() {
    part1("input1", 7, 7, 27, 27)
    part1("input", 200000000000000, 200000000000000, 400000000000000, 400000000000000)
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

type Point struct {
    x int
    y int
    z int
}

type Hailstone struct {
    position Point
    velocity Point
}

func linesToHailstones(lines []string) []Hailstone {
    hailstones := []Hailstone{}

    for _, line := range lines {
        if len(line) == 0 { continue }

        replacedLine := strings.ReplaceAll(line, ",", "")
        splitLine := strings.Fields(replacedLine)

        xPos, _ := strconv.Atoi(splitLine[0])
        yPos, _ := strconv.Atoi(splitLine[1])
        zPos, _ := strconv.Atoi(splitLine[2])
        xVel, _ := strconv.Atoi(splitLine[4])
        yVel, _ := strconv.Atoi(splitLine[5])
        zVel, _ := strconv.Atoi(splitLine[6])

        hailstone := Hailstone{
            position: Point{x: xPos, y: yPos, z: zPos},
            velocity: Point{x: xVel, y: yVel, z: zVel},
        }

        hailstones = append(hailstones, hailstone)
    }

    return hailstones
}

func isIntersection(hailstone1 Hailstone, hailstone2 Hailstone, xMin int, yMin int, xMax int, yMax int) bool {
    xPos1 := float64(hailstone1.position.x)
    xVel1 := float64(hailstone1.velocity.x)
    yPos1 := float64(hailstone1.position.y)
    yVel1 := float64(hailstone1.velocity.y)

    xPos2 := float64(hailstone2.position.x)
    xVel2 := float64(hailstone2.velocity.x)
    yPos2 := float64(hailstone2.position.y)
    yVel2 := float64(hailstone2.velocity.y)

    a1 := yVel1
    b1 := -xVel1
    c1 := xPos1 * yVel1 - yPos1 * xVel1
    a2 := yVel2
    b2 := -xVel2
    c2 := xPos2 * yVel2 - yPos2 * xVel2
    x := -(b1 * c2 - b2 * c1) / (a1 * b2 - a2 * b1)
    y := -(a2 * c1 - a1 * c2) / (a1 * b2 - a2 * b1)

    if xVel1 > 0 && x < xPos1 { return false }
    if xVel1 < 0 && x > xPos1 { return false }
    if yVel1 > 0 && y < yPos1 { return false }
    if yVel1 < 0 && y > yPos1 { return false }
    if xVel2 > 0 && x < xPos2 { return false }
    if xVel2 < 0 && x > xPos2 { return false }
    if yVel2 > 0 && y < yPos2 { return false }
    if yVel2 < 0 && y > yPos2 { return false }

    if x < float64(xMin) || x > float64(xMax) { return false }
    if y < float64(yMin) || y > float64(yMax) { return false }
    return true
}

func countIntersections(hailstones []Hailstone, xMin int, yMin int, xMax int, yMax int) int {
    count := 0

    for i := 0; i < len(hailstones); i++ {
        for j := i+1; j < len(hailstones); j++ {
            if isIntersection(hailstones[i], hailstones[j], xMin, yMin, xMax, yMax) {
                count++
            }
        }
    }

    return count
}

func part1(inputFilename string, xMin int, yMin int, xMax int, yMax int) {
    lines := getLines(inputFilename)
    hailstones := linesToHailstones(lines)

    intersections := countIntersections(hailstones, xMin, yMin, xMax, yMax)

    fmt.Println("solution", intersections)
}

