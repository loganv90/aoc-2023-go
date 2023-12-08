package main

import (
    "sort"
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

func createHand(cards string, bid int) Hand {
    counts := make(map[rune]int)

    for _, c := range cards {
        counts[c] += 1
    }

    if len(counts) == 1 {
        return Hand{cards, bid, 6} // five of a kind
    } else if len(counts) == 2 {
        for _, v := range counts {
            if v == 4 || v == 1 {
                return Hand{cards, bid, 5} // four of a kind
            } else {
                return Hand{cards, bid, 4} // full house
            }
        }
    } else if len(counts) == 3 {
        for _, v := range counts {
            if v == 3 {
                return Hand{cards, bid, 3} // three of a kind
            }
        }
        return Hand{cards, bid, 2} // two pairs
    } else if len(counts) == 4 {
        return Hand{cards, bid, 1} // one pair
    }

    return Hand{cards, bid, 0} // high card
}

type Hand struct {
    Cards string
    Bid int
    Type int
}

func (h Hand) LessThan(other Hand) bool {
    if h.Type != other.Type {
        return h.Type < other.Type
    }

    for i := 0; i < len(h.Cards); i++ {
        if h.Cards[i] != other.Cards[i] {
            return LessThan(rune(h.Cards[i]), rune(other.Cards[i]))
        }
    }

    return false
}

func LessThan(rune1 rune, rune2 rune) bool {
    if rune1 == rune2 {
        return false
    }

    if rune1 == 'A' || rune2 == 'A' {
        return rune2 == 'A'
    }

    if rune1 == 'K' || rune2 == 'K' {
        return rune2 == 'K'
    }

    if rune1 == 'Q' || rune2 == 'Q' {
        return rune2 == 'Q'
    }

    if rune1 == 'J' || rune2 == 'J' {
        return rune2 == 'J'
    }

    if rune1 == 'T' || rune2 == 'T' {
        return rune2 == 'T'
    }

    return rune1 < rune2
}

func part1(inputFilename string) {
    lines := getLines(inputFilename)

    hands := make([]Hand, 0)

    for _, line := range lines {
        if line == "" { continue }
        lineSplit := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
        bid, err := strconv.Atoi(lineSplit[1])
        if err != nil { panic(err) }
        cards := lineSplit[0]
        hands = append(hands, createHand(cards, bid))
    }

    sort.Slice(hands, func(i, j int) bool {
        return hands[i].LessThan(hands[j])
    })

    res := 0
    for i, hand := range hands {
        res += hand.Bid * (i + 1)
    }

    fmt.Println("solution", res)
}

func createHand2(cards string, bid int) Hand {
    counts := make(map[rune]int)
    jokers := 0

    for _, c := range cards {
        if c == 'J' {
            jokers += 1
            continue
        }
        counts[c] += 1
    }

    if len(counts) == 1 || jokers > 3 {
        return Hand{cards, bid, 6} // five of a kind
    } else if len(counts) == 2 {
        if jokers > 1 {
            return Hand{cards, bid, 5} // four of a kind
        } else if jokers == 1 {
            for _, v := range counts {
                if v == 3 || v == 1 {
                    return Hand{cards, bid, 5} // four of a kind
                }
            }
            return Hand{cards, bid, 4} // full house
        }

        for _, v := range counts {
            if v == 4 || v == 1 {
                return Hand{cards, bid, 5} // four of a kind
            } else {
                return Hand{cards, bid, 4} // full house
            }
        }
    } else if len(counts) == 3 {
        if jokers > 0 {
            return Hand{cards, bid, 3} // three of a kind
        }

        for _, v := range counts {
            if v == 3 {
                return Hand{cards, bid, 3} // three of a kind
            }
        }
        return Hand{cards, bid, 2} // two pairs
    } else if len(counts) == 4 {
        return Hand{cards, bid, 1} // one pair
    }

    return Hand{cards, bid, 0} // high card
}

func (h Hand) LessThan2(other Hand) bool {
    if h.Type != other.Type {
        return h.Type < other.Type
    }

    for i := 0; i < len(h.Cards); i++ {
        if h.Cards[i] != other.Cards[i] {
            return LessThan2(rune(h.Cards[i]), rune(other.Cards[i]))
        }
    }

    return false
}

func LessThan2(rune1 rune, rune2 rune) bool {
    if rune1 == rune2 {
        return false
    }

    if rune1 == 'J' || rune2 == 'J' {
        return rune1 == 'J'
    }

    if rune1 == 'A' || rune2 == 'A' {
        return rune2 == 'A'
    }

    if rune1 == 'K' || rune2 == 'K' {
        return rune2 == 'K'
    }

    if rune1 == 'Q' || rune2 == 'Q' {
        return rune2 == 'Q'
    }

    if rune1 == 'T' || rune2 == 'T' {
        return rune2 == 'T'
    }

    return rune1 < rune2
}

func part2(inputFilename string) {
    lines := getLines(inputFilename)

    hands := make([]Hand, 0)

    for _, line := range lines {
        if line == "" { continue }
        lineSplit := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
        bid, err := strconv.Atoi(lineSplit[1])
        if err != nil { panic(err) }
        cards := lineSplit[0]
        hands = append(hands, createHand2(cards, bid))
    }

    sort.Slice(hands, func(i, j int) bool {
        return hands[i].LessThan2(hands[j])
    })

    res := 0
    for i, hand := range hands {
        res += hand.Bid * (i + 1)
    }

    fmt.Println("solution", res)
}

