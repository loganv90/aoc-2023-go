package main

import (
    "regexp"
	"fmt"
	"os"
	"strings"
)

func main() {
    part1("input1")
    part1("input2")
    part1("input")
    part2("input3")
    part2_2("input4")
    part2_2("input")
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

type Node struct {
    left string
    right string
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    instructions := lines[0]
    nodes := make(map[string]Node)
    r := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

    for _, line := range lines[1:] {
        if line == "" { continue }

        matches := r.FindStringSubmatch(line)
        if r.NumSubexp() != 3 { panic("bad regex") }

        nodes[matches[1]] = Node{matches[2], matches[3]}
    }

    res := 0
    cur := "AAA"
    for {
        for _, c := range instructions {
            if cur == "ZZZ" {
                fmt.Println("solution", res)
                return
            }

            if c == 'R' {
                cur = nodes[cur].right
                res++
            } else if c == 'L' {
                cur = nodes[cur].left
                res++
            }
        }
    }
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    instructions := lines[0]
    nodes := make(map[string]Node)
    r := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)
    cur := make([]string, 0)
    res := 0

    for _, line := range lines[1:] {
        if line == "" { continue }

        matches := r.FindStringSubmatch(line)
        if r.NumSubexp() != 3 { panic("bad regex") }

        nodes[matches[1]] = Node{matches[2], matches[3]}

        if matches[1][len(matches[1]) - 1] == 'A' {
            cur = append(cur, matches[1])
        }
    }

    for {
        for _, c := range instructions {
            end := true
            for _, cu := range cur {
                if cu[len(cu) - 1] != 'Z' {
                    end = false
                }
            }
            if end {
                fmt.Println("solution", res)
                return
            }

            for i, cu := range cur {
                if c == 'R' {
                    cur[i] = nodes[cu].right
                } else if c == 'L' {
                    cur[i] = nodes[cu].left
                }
            }
            res++
        }
    }
}

type Loop struct {
    start int
    length int
    ends []int
}

func GCD(a int, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}

func LCM(a int, b int) int {
    return a * b / GCD(a, b)
}

func reduce(loop1 Loop, loop2 Loop) Loop {
    if loop1.ends[0] > loop2.ends[0] {
        temp := loop1
        loop1 = loop2
        loop2 = temp
    }

    current := loop1.ends[0]
    for current % loop2.length != loop2.ends[0] {
        current += loop1.length
    }
    
    return Loop{0, LCM(loop1.length, loop2.length), []int{current}}
}

func part2_2(inputFilename string) { // cus part2 is too slow
    lines := getLines(inputFilename)

    instructions := lines[0]
    nodes := make(map[string]Node)
    r := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)
    cur := make([]string, 0)

    for _, line := range lines[1:] {
        if line == "" { continue }

        matches := r.FindStringSubmatch(line)
        if r.NumSubexp() != 3 { panic("bad regex") }

        nodes[matches[1]] = Node{matches[2], matches[3]}

        if matches[1][len(matches[1]) - 1] == 'A' {
            cur = append(cur, matches[1])
        }
    }

    loops := make([]Loop, 0)
    for _, cu := range cur {
        loop := createLoop(instructions, nodes, cu)
        loops = append(loops, loop)

        if loop.start != 0 || len(loop.ends) != 1 {
            panic("math time") // luckily, this never happens
        }
    }

    reduced := loops[0]
    for _, loop := range loops[1:] {
        reduced = reduce(reduced, loop)
    }

    fmt.Println("solution", reduced.ends[0])
}

func createLoop(instructions string, nodes map[string]Node, cur string) Loop {
    loop := Loop{0, 0, []int{}}
    visited := make(map[string]int)
    current := cur
    counter := 0

    for {
        for i, c := range instructions {
            if index, ok := visited[current]; ok && i == index {
                loop.start = index
                loop.length = counter - index
                return loop
            }
            if current[len(current) - 1] == 'Z' {
                loop.ends = append(loop.ends, counter)
            }

            visited[current] = i
            if c == 'R' {
                current = nodes[current].right
            } else if c == 'L' {
                current = nodes[current].left
            }
            counter++
        }
    }
}

