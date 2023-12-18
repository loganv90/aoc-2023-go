package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
    "container/heap"
    "math"
)

func main() {
    part1("input1")
    part1("input")
    part2("input1")
    part2("input2")
    part2("input")
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

var directions = []*Point{
    {row: -1, col: 0},
    {row: 1, col: 0},
    {row: 0, col: -1},
    {row: 0, col: 1},
}

type HeadHeap []*Head

func (h HeadHeap) Len() int {
    return len(h)
}

func (h HeadHeap) Less(i, j int) bool {
    return h[i].distance < h[j].distance
}

func (h HeadHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *HeadHeap) Push(x interface{}) {
    *h = append(*h, x.(*Head))
}

func (h *HeadHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

type Space struct {
    number int
    distanceMap map[MapKey]int
}

type MapKey struct {
    row int
    col int
    straight int
    directionRow int
    directionCol int
}

type Point struct {
    row int
    col int
}

func (p *Point) Equal(other *Point) bool {
    return p.row == other.row && p.col == other.col
}

func (p *Point) Opposite(other *Point) bool {
    return p.row == -other.row && p.col == -other.col
}

type Head struct {
    distance int
    location *Point
    direction *Point
    straight int
}

func (h *Head) Move(direction *Point) (*Head, error) {
    if h.direction.Opposite(direction) {
        return nil, fmt.Errorf("cannot move in opposite direction")
    }

    var straight int
    if h.direction.Equal(direction) {
        straight = h.straight + 1
    } else {
        straight = 1
    }

    if straight > 3 {
        return nil, fmt.Errorf("cannot move straight more than 3 times")
    }

    return &Head{
        distance: h.distance,
        location: &Point{row: h.location.row + direction.row, col: h.location.col + direction.col},
        direction: direction,
        straight: straight,
    }, nil
}

func (h *Head) Move2(direction *Point) (*Head, error) {
    if h.direction.Opposite(direction) {
        return nil, fmt.Errorf("cannot move in opposite direction")
    }

    var straight int
    var dRow int
    var dCol int
    if h.direction.Equal(direction) {
        straight = h.straight + 1
        dRow = direction.row
        dCol = direction.col
    } else {
        straight = 4
        dRow = direction.row * 4
        dCol = direction.col * 4
    }

    if straight > 10 {
        return nil, fmt.Errorf("cannot move straight more than 3 times")
    }

    return &Head{
        distance: h.distance,
        location: &Point{row: h.location.row + dRow, col: h.location.col + dCol},
        direction: direction,
        straight: straight,
    }, nil
}

func dijkstra(spaces [][]*Space) {
    queue := &HeadHeap{
        &Head{
            distance: 0,
            location: &Point{row: 0, col: 0},
            direction: &Point{row: 0, col: 0},
            straight: 0,
        },
    }

    for queue.Len() > 0 {
        currentHead := heap.Pop(queue).(*Head)

        for _, direction := range directions {
            movedHead, err := currentHead.Move(direction)

            if err != nil { continue }
            if movedHead.location.row < 0 || movedHead.location.row >= len(spaces) { continue }
            if movedHead.location.col < 0 || movedHead.location.col >= len(spaces[0]) { continue }

            movedHead.distance += spaces[movedHead.location.row][movedHead.location.col].number

            mapKey := MapKey{
                row: movedHead.location.row,
                col: movedHead.location.col,
                straight: movedHead.straight,
                directionRow: movedHead.direction.row,
                directionCol: movedHead.direction.col,
            }

            spaceDistance, ok := spaces[movedHead.location.row][movedHead.location.col].distanceMap[mapKey]
            if ok && movedHead.distance >= spaceDistance { continue }

            spaces[movedHead.location.row][movedHead.location.col].distanceMap[mapKey] = movedHead.distance

            heap.Push(queue, movedHead)
        }
    }
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    spaces := [][]*Space{}
    for _, line := range lines {
        if len(line) == 0 { continue }
        row := []*Space{}
        for _, char := range line {
            number, err := strconv.Atoi(string(char))
            if err != nil { panic(err) }

            space := &Space{
                number: number,
                distanceMap: map[MapKey]int{},
            }
            row = append(row, space)
        }
        spaces = append(spaces, row)
    }

    dijkstra(spaces)

    res := math.MaxInt
    for _, value := range spaces[len(spaces) - 1][len(spaces[0]) - 1].distanceMap {
        if value < res {
            res = value
        }
    }

    fmt.Println("solution", res)
}

func dijkstra2(spaces [][]*Space) {
    queue := &HeadHeap{
        &Head{
            distance: 0,
            location: &Point{row: 0, col: 0},
            direction: &Point{row: 0, col: 0},
            straight: 0,
        },
    }

    for queue.Len() > 0 {
        currentHead := heap.Pop(queue).(*Head)

        for _, direction := range directions {
            movedHead, err := currentHead.Move2(direction)

            if err != nil { continue }
            if movedHead.location.row < 0 || movedHead.location.row >= len(spaces) { continue }
            if movedHead.location.col < 0 || movedHead.location.col >= len(spaces[0]) { continue }

            if currentHead.location.row > movedHead.location.row {
                for r := currentHead.location.row - 1; r >= movedHead.location.row; r-- {
                    movedHead.distance += spaces[r][movedHead.location.col].number
                }
            } else if currentHead.location.row < movedHead.location.row {
                for r := currentHead.location.row + 1; r <= movedHead.location.row; r++ {
                    movedHead.distance += spaces[r][movedHead.location.col].number
                }
            } else if currentHead.location.col > movedHead.location.col {
                for c := currentHead.location.col - 1; c >= movedHead.location.col; c-- {
                    movedHead.distance += spaces[movedHead.location.row][c].number
                }
            } else if currentHead.location.col < movedHead.location.col {
                for c := currentHead.location.col + 1; c <= movedHead.location.col; c++ {
                    movedHead.distance += spaces[movedHead.location.row][c].number
                }
            } else {
                panic("same location")
            }

            mapKey := MapKey{
                row: movedHead.location.row,
                col: movedHead.location.col,
                straight: movedHead.straight,
                directionRow: movedHead.direction.row,
                directionCol: movedHead.direction.col,
            }

            spaceDistance, ok := spaces[movedHead.location.row][movedHead.location.col].distanceMap[mapKey]
            if ok && movedHead.distance >= spaceDistance { continue }

            spaces[movedHead.location.row][movedHead.location.col].distanceMap[mapKey] = movedHead.distance

            heap.Push(queue, movedHead)
        }
    }
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    spaces := [][]*Space{}
    for _, line := range lines {
        if len(line) == 0 { continue }
        row := []*Space{}
        for _, char := range line {
            number, err := strconv.Atoi(string(char))
            if err != nil { panic(err) }

            space := &Space{
                number: number,
                distanceMap: map[MapKey]int{},
            }
            row = append(row, space)
        }
        spaces = append(spaces, row)
    }

    dijkstra2(spaces)

    res := math.MaxInt
    for _, value := range spaces[len(spaces) - 1][len(spaces[0]) - 1].distanceMap {
        if value < res {
            res = value
        }
    }

    fmt.Println("solution", res)
}

