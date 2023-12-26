package main

import (
	"fmt"
	"os"
	"strings"
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

func linesToGrid(lines []string) [][]string {
    grid := [][]string{}

    for _, line := range lines {
        if len(line) == 0 { continue }
        row := []string{}
        for _, char := range line {
            row = append(row, string(char))
        }
        grid = append(grid, row)
    }

    return grid
}

type Point struct {
    row int
    col int
}

type Direction struct {
    row int
    col int
    slope string
}

var directions = []Direction{
    {1, 0, "v"},
    {-1, 0, "^"},
    {0, 1, ">"},
    {0, -1, "<"},
}

func getStartPoint(grid [][]string) Point {
    startPoint := Point{-1, -1}

    for i, char := range grid[0] {
        if char == "." {
            startPoint = Point{0, i}
            break
        }
    }
    if startPoint.row == -1 || startPoint.col == -1 {
        panic("no start point found")
    }

    return startPoint
}

func getEndPoint(grid [][]string) Point {
    endPoint := Point{-1, -1}

    for i, char := range grid[len(grid) - 1] {
        if char == "." {
            endPoint = Point{len(grid) - 1, i}
            break
        }
    }
    if endPoint.row == -1 || endPoint.col == -1 {
        panic("no start point found")
    }

    return endPoint
}

func dfs(grid [][]string, currentPoint Point, endPoint Point, visited map[Point]bool) int {
    count := 0

    visited[currentPoint] = true

    for _, direction := range directions {
        nextRow := currentPoint.row + direction.row
        nextCol := currentPoint.col + direction.col
        if nextRow < 0 || nextRow >= len(grid) {
            continue
        }
        if nextCol < 0 || nextCol >= len(grid[nextRow]) {
            continue
        }

        nextPoint := Point{ row: nextRow, col: nextCol }
        if _, ok := visited[nextPoint]; ok {
            continue
        }

        nextChar := grid[nextPoint.row][nextPoint.col]
        if nextChar == "." || nextChar == direction.slope {
            count = max(count, dfs(grid, nextPoint, endPoint, visited))
        }
    }

    delete(visited, currentPoint)

    if count > 0 {
        return count + 1
    } else if currentPoint.row == endPoint.row && currentPoint.col == endPoint.col {
        return 1
    } else {
        return 0
    }
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)
    grid := linesToGrid(lines)

    startPoint := getStartPoint(grid)
    endPoint := getEndPoint(grid)

    longestPath := dfs(grid, startPoint, endPoint, map[Point]bool{}) - 1

    fmt.Println("solution", longestPath)
}

func dfsGraph(curr Point, end Point, visited map[Point]bool, graph map[Point]map[Point]int) int {
    visited[curr] = true

    count := 0 

    for nextPoint, weight := range graph[curr] {
        if _, ok := visited[nextPoint]; ok {
            continue
        }

        res := dfsGraph(nextPoint, end, visited, graph)
        if res > 0 {
            count = max(count, res + weight)
        }
    }

    delete(visited, curr)

    if count > 0 {
        return count
    } else if curr.row == end.row && curr.col == end.col {
        return 1
    } else {
        return 0
    }
}

func gridToGraph(grid [][]string, current Point, visited map[Point]bool, graph map[Point]map[Point]int) {
    visited[current] = true
    nextPoints := []Point{}

    for _, direction := range directions {
        nextRow := current.row + direction.row
        nextCol := current.col + direction.col
        if nextRow < 0 || nextRow >= len(grid) {
            continue
        }
        if nextCol < 0 || nextCol >= len(grid[nextRow]) {
            continue
        }

        nextPoint := Point{ row: nextRow, col: nextCol }
        if _, ok := visited[nextPoint]; ok {
            continue
        }

        nextChar := grid[nextPoint.row][nextPoint.col]
        if nextChar == "#" {
            continue
        }

        if _, ok := graph[current]; !ok {
            graph[current] = map[Point]int{}
        }
        graph[current][nextPoint] = 1

        if _, ok := graph[nextPoint]; !ok {
            graph[nextPoint] = map[Point]int{}
        }
        graph[nextPoint][current] = 1

        nextPoints = append(nextPoints, nextPoint)
    }

    for _, nextPoint := range nextPoints {
        gridToGraph(grid, nextPoint, visited, graph)
    }
}

func removeVertexWithTwoEdges(graph map[Point]map[Point]int) {
    vertexToRemove := []Point{}

    for vertex, edges := range graph {
        if len(edges) == 2 {
            vertexToRemove = append(vertexToRemove, vertex)
        }
    }

    for _, vertex := range vertexToRemove {
        edgeMap := graph[vertex]
        edges := []Point{}
        for edge := range edgeMap {
            edges = append(edges, edge)
        }

        edge0 := graph[edges[0]]
        cost0 := edge0[vertex]
        edge1 := graph[edges[1]]
        cost1 := edge1[vertex]

        edge0[edges[1]] = cost0 + cost1
        edge1[edges[0]] = cost0 + cost1

        delete(edge0, vertex)
        delete(edge1, vertex)
        delete(graph, vertex)
    }
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)
    grid := linesToGrid(lines)

    startPoint := getStartPoint(grid)
    endPoint := getEndPoint(grid)

    visited := map[Point]bool{}
    graph := map[Point]map[Point]int{}

    gridToGraph(grid, startPoint, visited, graph)
    removeVertexWithTwoEdges(graph)

    longestPath := dfsGraph(startPoint, endPoint, map[Point]bool{}, graph) - 1
    fmt.Println("solution", longestPath)
}

