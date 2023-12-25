package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
    "container/heap"
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

type Point struct {
    x int
    y int
    z int
}

type Brick struct {
    name int
    start *Point
    end *Point
}

type BrickHeap []*Brick

func (h BrickHeap) Len() int {
    return len(h)
}

func (h BrickHeap) Less(i, j int) bool {
    iMin := min(h[i].start.z, h[i].end.z)
    jMin := min(h[j].start.z, h[j].end.z)
    return iMin < jMin
}

func (h BrickHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *BrickHeap) Push(x interface{}) {
    *h = append(*h, x.(*Brick))
}

func (h *BrickHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

func linesToBricks(lines []string) ([]*Brick, int, int) {
    xMax := 0
    yMax := 0
    grid := []*Brick{}

    for i, line := range lines {
        if len(line) == 0 { continue }
        replacedLine := strings.ReplaceAll(line, "~", ",")
        splitLine := strings.Split(replacedLine, ",")
        xStart, _ := strconv.Atoi(splitLine[0])
        yStart, _ := strconv.Atoi(splitLine[1])
        zStart, _ := strconv.Atoi(splitLine[2])
        xEnd, _ := strconv.Atoi(splitLine[3])
        yEnd, _ := strconv.Atoi(splitLine[4])
        zEnd, _ := strconv.Atoi(splitLine[5])

        brick := &Brick{
            name: i,
            start: &Point{
                x: xStart,
                y: yStart,
                z: zStart,
            },
            end: &Point{
                x: xEnd,
                y: yEnd,
                z: zEnd,
            },
        }

        if xEnd > xMax { xMax = xEnd }
        if xStart > xMax { xMax = xStart }

        if yEnd > yMax { yMax = yEnd }
        if yStart > yMax { yMax = yStart }

        grid = append(grid, brick)
    }

    return grid, xMax, yMax
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)
    bricks, xMaxGrid, yMaxGrid := linesToBricks(lines)

    maxes := [][]int{}
    for y := 0; y <= yMaxGrid; y++ {
        row := []int{}
        for x := 0; x <= xMaxGrid; x++ {
            row = append(row, 0)
        }
        maxes = append(maxes, row)
    }

    grid := [][][]*Brick{}
    for z := 0; z <= 1000; z++ {
        layer := [][]*Brick{}
        for y := 0; y <= yMaxGrid; y++ {
            row := []*Brick{}
            for x := 0; x <= xMaxGrid; x++ {
                row = append(row, nil)
            }
            layer = append(layer, row)
        }
        grid = append(grid, layer)
    }

    brickHeap := &BrickHeap{}
    for _, brick := range bricks {
        heap.Push(brickHeap, brick)
    }

    supportingBricksToSupportedBricks := map[*Brick]map[*Brick]bool{}
    supportedBricksToSupportingBricks := map[*Brick]map[*Brick]bool{}
    for brickHeap.Len() > 0 {
        brick := heap.Pop(brickHeap).(*Brick)
        zMaxMaxes := 0

        xMin := min(brick.start.x, brick.end.x)
        xMax := max(brick.start.x, brick.end.x)
        yMin := min(brick.start.y, brick.end.y)
        yMax := max(brick.start.y, brick.end.y)
        zMin := min(brick.start.z, brick.end.z)
        zMax := max(brick.start.z, brick.end.z)
        for x := xMin; x <= xMax; x++ {
            for y := yMin; y <= yMax; y++ {
                zMaxMaxes = max(zMaxMaxes, maxes[y][x])
            }
        }

        zMinMaxes := zMaxMaxes + 1
        diff := min(brick.end.z, brick.start.z) - zMinMaxes
        for x := xMin; x <= xMax; x++ {
            for y := yMin; y <= yMax; y++ {
                for z := zMin; z <= zMax; z++ {
                    grid[z - diff][y][x] = brick
                    maxes[y][x] = max(maxes[y][x], z - diff)

                    if z - diff - 1 > 0 {
                        supportingBrick := grid[z - diff - 1][y][x]
                        if supportingBrick == nil { continue }
                        if supportingBrick == brick { continue }

                        if _, ok := supportingBricksToSupportedBricks[supportingBrick]; !ok {
                            supportingBricksToSupportedBricks[supportingBrick] = map[*Brick]bool{}
                        }
                        if _, ok := supportedBricksToSupportingBricks[brick]; !ok {
                            supportedBricksToSupportingBricks[brick] = map[*Brick]bool{}
                        }

                        supportingBricksToSupportedBricks[supportingBrick][brick] = true
                        supportedBricksToSupportingBricks[brick][supportingBrick] = true
                    }
                }
            }
        }
    }

    count := 0
    for _, brick := range bricks {
        supportedBricks := supportingBricksToSupportedBricks[brick]

        allSupported := true
        for supportedBrick := range supportedBricks {
            supportingBricks := supportedBricksToSupportingBricks[supportedBrick]

            if len(supportingBricks) <= 1 {
                allSupported = false
                break
            }
        }
        if allSupported {
            count++
        }
    }

    fmt.Println("solution", count)
}

