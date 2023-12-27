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
    part2("input")
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

func isIntersection2(hailstone1 Hailstone, hailstone2 Hailstone, xVel float64, yVel float64, zVel float64) (float64, float64, float64) {
    xPos1 := float64(hailstone1.position.x)
    xVel1 := float64(hailstone1.velocity.x)
    yPos1 := float64(hailstone1.position.y)
    yVel1 := float64(hailstone1.velocity.y)

    xPos2 := float64(hailstone2.position.x)
    xVel2 := float64(hailstone2.velocity.x)
    yPos2 := float64(hailstone2.position.y)
    yVel2 := float64(hailstone2.velocity.y)

    a1 := yVel1-yVel
    b1 := -xVel1+xVel
    c1 := xPos1 * (yVel1-yVel) - yPos1 * (xVel1-xVel)
    a2 := yVel2-yVel
    b2 := -xVel2+xVel
    c2 := xPos2 * (yVel2-yVel) - yPos2 * (xVel2-xVel)
    x := -(b1 * c2 - b2 * c1) / (a1 * b2 - a2 * b1)
    y := -(a2 * c1 - a1 * c2) / (a1 * b2 - a2 * b1)

    t := (x - xPos1) / (xVel1 - xVel)
    z := float64(hailstone1.position.z) + t*(float64(hailstone1.velocity.z)-zVel)

    return x, y, z
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)
    hailstones := linesToHailstones(lines)

    xMap := map[int]bool{}
    yMap := map[int]bool{}
    zMap := map[int]bool{}
    for i := 0; i < len(hailstones); i++ {
        for j := i+1; j < len(hailstones); j++ {
            xPos1 := hailstones[i].position.x
            yPos1 := hailstones[i].position.y
            zPos1 := hailstones[i].position.z
            xVel1 := hailstones[i].velocity.x
            yVel1 := hailstones[i].velocity.y
            zVel1 := hailstones[i].velocity.z
            xPos2 := hailstones[j].position.x
            yPos2 := hailstones[j].position.y
            zPos2 := hailstones[j].position.z
            xVel2 := hailstones[j].velocity.x
            yVel2 := hailstones[j].velocity.y
            zVel2 := hailstones[j].velocity.z

            if xVel1 == xVel2 && (xVel1 > 100 || xVel1 < -100) {
                xMapNew := map[int]bool{}
                diff := xPos2 - xPos1
                for i := -1000; i < 1000; i++ {
                    if i == xVel1 { continue }
                    if diff % (i - xVel1) == 0 {
                        xMapNew[i] = true
                    }
                }
                if len(xMap) == 0 {
                    xMap = xMapNew
                } else {
                    for key := range xMap {
                        if _, ok := xMapNew[key]; !ok {
                            delete(xMap, key)
                        }
                    }
                }
            }
            if yVel1 == yVel2 && (yVel1 > 100 || yVel1 < -100) {
                yMapNew := map[int]bool{}
                diff := yPos2 - yPos1
                for i := -1000; i < 1000; i++ {
                    if i == yVel1 { continue }
                    if diff % (i - yVel1) == 0 {
                        yMapNew[i] = true
                    }
                }
                if len(yMap) == 0 {
                    yMap = yMapNew
                } else {
                    for key := range yMap {
                        if _, ok := yMapNew[key]; !ok {
                            delete(yMap, key)
                        }
                    }
                }
            }
            if zVel1 == zVel2 && (zVel1 > 100 || zVel1 < -100) {
                zMapNew := map[int]bool{}
                diff := zPos2 - zPos1
                for i := -1000; i < 1000; i++ {
                    if i == zVel1 { continue }
                    if diff % (i - zVel1) == 0 {
                        zMapNew[i] = true
                    }
                }
                if len(zMap) == 0 {
                    zMap = zMapNew
                } else {
                    for key := range zMap {
                        if _, ok := zMapNew[key]; !ok {
                            delete(zMap, key)
                        }
                    }
                }
            }
        }
    }
    xVel := 0
    for key := range xMap {
        xVel = key
        break
    }
    yVel := 0
    for key := range yMap {
        yVel = key
        break
    }
    zVel := 0
    for key := range zMap {
        zVel = key
        break
    }

    x, y, z := isIntersection2(hailstones[0], hailstones[1], float64(xVel), float64(yVel), float64(zVel))

    fmt.Println("solution", int(x + y + z))
}

