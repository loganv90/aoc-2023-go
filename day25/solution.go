package main

import (
	"fmt"
	"os"
	"strings"
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

func linesToGraph(lines []string) map[string]map[string]bool {
    graph := map[string]map[string]bool{}

    for _, line := range lines {
        if line == "" { continue }
        replacedLine := strings.ReplaceAll(line, ":", "")
        splitLine := strings.Split(replacedLine, " ")

        component := splitLine[0]
        connections := splitLine[1:]

        if _, ok := graph[component]; !ok {
            graph[component] = map[string]bool{}
        }

        for _, connection := range connections {
            if _, ok := graph[connection]; !ok {
                graph[connection] = map[string]bool{}
            }
        }

        for _, connection := range connections {
            graph[component][connection] = true
            graph[connection][component] = true
        }
    }

    return graph
}

func bfs(graph map[string]map[string]bool, start string, end string, visitedEdges map[string]map[string]bool) []string {
    visitedVertices := map[string]bool{}

    queue := make(chan []string, 100000)
    queue <- []string{start}

    for len(queue) > 0 {
        path := <-queue
        vertex := path[len(path) - 1]

        if _, ok := visitedVertices[vertex]; ok { continue }
        visitedVertices[vertex] = true

        if vertex == end {
            return path
        }

        for neighbor := range graph[vertex] {
            if _, ok := visitedEdges[vertex][neighbor]; ok { continue }
            if _, ok := visitedEdges[neighbor][vertex]; ok { continue }
            visitedEdges[vertex][neighbor] = true
            visitedEdges[neighbor][vertex] = true

            newPath := make([]string, len(path) + 1)
            copy(newPath, path)
            newPath[len(newPath) - 1] = neighbor
            queue <- newPath
        }
    }

    return []string{}
}

func countPaths(graph map[string]map[string]bool, start string, end string) int {
    visitedEdges := map[string]map[string]bool{}
    for vertex := range graph {
        visitedEdges[vertex] = map[string]bool{}
    }
    visitedEdges[start][end] = true
    visitedEdges[end][start] = true
    count := 0

    for {
        localVisitedEdges := map[string]map[string]bool{}
        for vertex := range graph {
            localVisitedEdges[vertex] = map[string]bool{}
        }
        for vertex1, value := range visitedEdges {
            for vertex2 := range value {
                localVisitedEdges[vertex1][vertex2] = true
            }
        }

        path := bfs(graph, start, end, localVisitedEdges)
        if len(path) == 0 {
            break
        }

        count++
        for i := 0; i < len(path) - 1; i++ {
            visitedEdges[path[i]][path[i + 1]] = true
            visitedEdges[path[i + 1]][path[i]] = true
        }
    }

    return count
}

func disconnectGraph(graph map[string]map[string]bool) (string, string) {
    edgesToRemove := map[string]map[string]bool{}
    for vertex1, value := range graph {
        for vertex2 := range value {
            count := countPaths(graph, vertex1, vertex2)
            if count < 3 {
                if _, ok := edgesToRemove[vertex1]; !ok {
                    edgesToRemove[vertex1] = map[string]bool{}
                }
                edgesToRemove[vertex1][vertex2] = true
            }
        }
    }

    var vertexSection1, vertexSection2 string
    for vertex1, value := range edgesToRemove {
        for vertex2 := range value {
            delete(graph[vertex1], vertex2)
            vertexSection1 = vertex1
            vertexSection2 = vertex2
        }
    }

    return vertexSection1, vertexSection2
}

func countVertices(graph map[string]map[string]bool, start string) int {
    visitedVertices := map[string]bool{}

    queue := make(chan string, 100000)
    queue <- start

    for len(queue) > 0 {
        vertex := <-queue

        if _, ok := visitedVertices[vertex]; ok { continue }
        visitedVertices[vertex] = true

        for neighbor := range graph[vertex] {
            queue <- neighbor
        }
    }

    return len(visitedVertices)
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)
    graph := linesToGraph(lines)

    vertexSection1, vertexSection2 := disconnectGraph(graph)
    countSection1 := countVertices(graph, vertexSection1)
    countSection2 := countVertices(graph, vertexSection2)

    fmt.Println("solution", countSection1 * countSection2)
}

