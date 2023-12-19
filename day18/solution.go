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

type Point struct {
    row int
    col int
}

type Instruction struct {
    direction *Point
    distance int
    color string
}

type MatrixPoint struct {
    dug bool
    color string
}

func convertInstructionsToMatrix(instructions []*Instruction) {
    maxRow := 0
    maxCol := 0
    minRow := 0
    minCol := 0
    matrixPointMap := map[Point]*MatrixPoint{}

    currentPoint := &Point{0, 0}
    for _, instruction := range instructions {
        for i := 0; i < instruction.distance; i++ {
            currentPoint.row += instruction.direction.row
            currentPoint.col += instruction.direction.col

            if currentPoint.row > maxRow { maxRow = currentPoint.row }
            if currentPoint.col > maxCol { maxCol = currentPoint.col }
            if currentPoint.row < minRow { minRow = currentPoint.row }
            if currentPoint.col < minCol { minCol = currentPoint.col }

            matrixPointMap[*currentPoint] = &MatrixPoint{true, instruction.color}
        }
    }

    matrix := [][]*MatrixPoint{}
    for row := minRow; row <= maxRow; row++ {
        matrixRow := []*MatrixPoint{}
        for col := minCol; col <= maxCol; col++ {
            if matrixPoint, ok := matrixPointMap[Point{row, col}]; ok {
                matrixRow = append(matrixRow, matrixPoint)
            } else {
                matrixRow = append(matrixRow, &MatrixPoint{false, ""})
            }
        }
        matrix = append(matrix, matrixRow)
    }

    area := calculateDugArea(matrix) + len(matrixPointMap)
    fmt.Println("solution", area)
}

func calculateDugArea(matrix [][]*MatrixPoint) int {
    area := 0

    for row := 0; row < len(matrix); row++ {
        insideDug := false

        for col := 0; col < len(matrix[row]); col++ {
            if !insideDug && !matrix[row][col].dug {
                continue
            } else if insideDug && !matrix[row][col].dug {
                area++
                continue
            }

            startDug := col
            endDug := col
            for i := col+1; i < len(matrix[row]); i++ {
                if matrix[row][i].dug {
                    endDug = i
                } else {
                    break
                }
            }
            col = endDug

            if startDug == endDug {
                insideDug = !insideDug
                continue
            }

            if row > 0 && matrix[row-1][startDug].dug && matrix[row-1][endDug].dug {
                continue
            }

            if row < len(matrix)-1 && matrix[row+1][startDug].dug && matrix[row+1][endDug].dug {
                continue
            }

            insideDug = !insideDug
        }
    }

    return area
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    instructions := make([]*Instruction, 0)
    for _, line := range lines {
        if len(line) == 0 { continue }

        lineSplit := strings.Split(line, " ")

        var direction *Point
        if lineSplit[0] == "R" {
            direction = &Point{0, 1}
        } else if lineSplit[0] == "L" {
            direction = &Point{0, -1}
        } else if lineSplit[0] == "U" {
            direction = &Point{-1, 0}
        } else if lineSplit[0] == "D" {
            direction = &Point{1, 0}
        } else {
            panic("unknown direction")
        }

        distance, err := strconv.Atoi(lineSplit[1])
        if err != nil { panic("can't parse distance") }

        color := lineSplit[2][1:len(lineSplit[2])-1]

        instructions = append(instructions, &Instruction{direction, distance, color})
    }

    convertInstructionsToMatrix(instructions)
}

func calculateAreaWithShoelaceFormula(instructions []*Instruction) {
    pointList := make([]Point, 0)
    perimeter := 2 // pree sure this accounts for the outer vs inner edges of the box

    currentPoint := &Point{0, 0}
    for _, instruction := range instructions {
        perimeter += instruction.distance

        currentPoint.row += instruction.direction.row * instruction.distance
        currentPoint.col += instruction.direction.col * instruction.distance

        pointList = append(pointList, *currentPoint)
    }

    area := 0
    for i := 0; i < len(pointList)-1; i++ {
        p1 := pointList[i]
        p2 := pointList[i+1]

        area += (p1.row * p2.col) - (p1.col * p2.row)
    }
    p1 := pointList[len(pointList)-1]
    p2 := pointList[0]
    area += (p1.row * p2.col) - (p1.col * p2.row)

    if area < 0 { area = -area }
    area += perimeter

    fmt.Println("solution", area/2)
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    instructions := make([]*Instruction, 0)
    for _, line := range lines {
        if len(line) == 0 { continue }

        lineSplit := strings.Split(line, " ")

        color := lineSplit[2][1:len(lineSplit[2])-1]

        distanceString := lineSplit[2][2:7]
        distance64, err := strconv.ParseInt(distanceString, 16, 64)
        if err != nil { panic("can't parse distance") }
        distance := int(distance64)
        if int64(distance) != distance64 { panic("can't parse distance") }

        var direction *Point
        directionString := lineSplit[2][7:8]
        if directionString == "0" {
            direction = &Point{0, 1}
        } else if directionString == "1" {
            direction = &Point{1, 0}
        } else if directionString == "2" {
            direction = &Point{0, -1}
        } else if directionString == "3" {
            direction = &Point{-1, 0}
        } else {
            panic("unknown direction")
        }

        instructions = append(instructions, &Instruction{direction, distance, color})
    }

    calculateAreaWithShoelaceFormula(instructions)
}

