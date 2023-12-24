package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    part1("input1")
    part1("input2")
    part1("input")
    part2("input3")
    part2("input")
}

func getLines(inputFilename string) []string {
    text, err := os.ReadFile(inputFilename)
    if err != nil { panic(err) }

    return strings.Split(string(text), "\n")
}

type Signal struct {
    receiver string
    sender string
    high bool
}

type Module interface {
    receive(signal Signal)
    addInput(input string)
    getOutputs() []string
}

type FlipFlop struct {
    queue chan Signal
    outputs []string
    high bool
}

func (f *FlipFlop) receive(signal Signal) {
    if signal.high {
        return
    }

    f.high = !f.high

    for _, output := range f.outputs {
        f.queue <- Signal{output, signal.receiver, f.high}
    }
}

func (f *FlipFlop) addInput(input string) {
    return
}

func (f *FlipFlop) getOutputs() []string {
    return f.outputs
}

type Conjunction struct {
    queue chan Signal
    outputs []string
    highs map[string]bool
}

func (c *Conjunction) receive(signal Signal) {
    c.highs[signal.sender] = signal.high

    allHigh := true
    for _, state := range c.highs {
        if !state {
            allHigh = false
            break
        }
    }

    for _, output := range c.outputs {
        c.queue <- Signal{output, signal.receiver, !allHigh}
    }
}

func (c *Conjunction) addInput(input string) {
    c.highs[input] = false
}

func (c *Conjunction) getOutputs() []string {
    return c.outputs
}

type Broadcaster struct {
    queue chan Signal
    outputs []string
}

func (b *Broadcaster) receive(signal Signal) {
    for _, output := range b.outputs {
        b.queue <- Signal{output, signal.receiver, signal.high}
    }
}

func (b *Broadcaster) addInput(input string) {
    return
}

func (b *Broadcaster) getOutputs() []string {
    return b.outputs
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    modules := make(map[string]Module)
    queue := make(chan Signal, 100)
    createModules(lines, modules, queue)

    lowCount := 0
    highCount := 0
    for i := 0; i < 1000; i++ {
        l, h := pressButton(queue, modules)
        lowCount += l
        highCount += h
    }
    fmt.Println("solution", lowCount * highCount)
}

func pressButton(queue chan Signal, modules map[string]Module) (int, int) {
    queue <- Signal{"broadcaster", "", false}
    endLoop := false
    lowCount := 0
    highCount := 0
    for !endLoop {
        select {
        case x, ok := <-queue:
            if !ok {
                panic("channel closed")
            }
            if x.high {
                highCount++
            } else {
                lowCount++
            }
            if _, ok := modules[x.receiver]; ok {
                modules[x.receiver].receive(x)
            }
        default:
            endLoop = true
        }
    }

    return lowCount, highCount
}

func createModules(lines []string, modules map[string]Module, queue chan Signal) {
    for _, line := range lines {
        if len(line) <= 0 {
            continue
        }

        replacedLine := strings.ReplaceAll(line, ",", "")
        lineSplit := strings.Split(replacedLine, " ")

        input := lineSplit[0]
        outputs := lineSplit[2:]
        inputType := "b"
        if input != "broadcaster" {
            inputType = string(input[0])
            input = input[1:]
        }

        var module Module
        if inputType == "b" {
            module = &Broadcaster{queue, outputs}
        } else if inputType == "&" {
            module = &Conjunction{queue, outputs, make(map[string]bool)}
        } else if inputType == "%" {
            module = &FlipFlop{queue, outputs, false}
        } else {
            panic("unknown input type")
        }

        modules[input] = module
    }
    for name, module := range modules {
        for _, output := range module.getOutputs() {
            if _, ok := modules[output]; ok {
                modules[output].addInput(name)
            }
        }
    }
}

func printModules(modules map[string]Module) {
    for name, module := range modules {
        fmt.Println("name", name)
        fmt.Println("outputs", module.getOutputs())
        if flipFlop, ok := module.(*FlipFlop); ok {
            fmt.Println("flip flop high", flipFlop.high)
        } else if conjunction, ok := module.(*Conjunction); ok {
            fmt.Println("conjunction highs", conjunction.highs)
        } else if _, ok := module.(*Broadcaster); ok {
            fmt.Println("broadcaster")
        } else {
            panic("unknown module type")
        }
        fmt.Println()
    }
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    modules := make(map[string]Module)
    queue := make(chan Signal, 100)
    createModules(lines, modules, queue)

    watches := map[string][]int{}
    rxInput := ""
    for name, module := range modules {
        for _, output := range module.getOutputs() {
            if output == "rx" {
                _, ok := module.(*Conjunction)
                if !ok {
                    panic("rxInput")
                }
                rxInput = name
            }
        }
    }
    for name, module := range modules {
        for _, output := range module.getOutputs() {
            if output == rxInput {
                _, ok := module.(*Conjunction)
                if !ok {
                    panic("watches")
                }
                watches[name] = []int{}
            }
        }
    }

    for i := 0; i < 10000; i++ {
        pressButton2(queue, modules, watches, i)
    }

    loops := []Loop{}
    for _, list := range watches {
        start := list[0]
        length := list[1] - list[0]
        loops = append(loops, Loop{start, length})
    }

    reduced := loops[0]
    for _, loop := range loops[1:] {
        reduced = reduce(reduced, loop)
    }

    fmt.Println("solution", reduced.start)
}

type Loop struct {
    start int
    length int
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
    return Loop{LCM(loop1.start, loop2.start), LCM(loop1.length, loop2.length)}
}

func pressButton2(queue chan Signal, modules map[string]Module, watches map[string][]int, i int) (int, int) {
    queue <- Signal{"broadcaster", "", false}
    endLoop := false
    lowCount := 0
    highCount := 0
    for !endLoop {
        select {
        case x, ok := <-queue:
            if !ok {
                panic("channel closed")
            }
            if x.high {
                highCount++
            } else {
                lowCount++
            }
            if x.high {
                for name, list := range watches {
                    if name == x.sender {
                        list = append(list, i+1)
                        watches[name] = list
                    }
                }
            }
            if _, ok := modules[x.receiver]; ok {
                modules[x.receiver].receive(x)
            }
        default:
            endLoop = true
        }
    }

    return lowCount, highCount
}

