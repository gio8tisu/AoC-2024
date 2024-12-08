package main

import (
    "fmt"
    "log"

    "advent-of-code/internal/utils"
)

type coordinate struct {
    row int
    col int
}

func (c1 *coordinate) diff(c2 *coordinate) coordinate {
    return coordinate{row: c1.row - c2.row, col: c1.col - c2.col}
}
func (c1 *coordinate) dist(c2 *coordinate) coordinate {
    return coordinate{row: c2.row - c1.row, col: c2.col - c1.col}
}

func main() {
    input, err := utils.ReadFile("input/day08.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    var maxCol int
    antennas := make(map[rune][]coordinate)
    for i, line := range input {
        for j, c := range line {
            maxCol = j
            if c == '.' {
                continue
            }
            coord := coordinate{row: i, col: j}
            val, ok := antennas[c]
            if !ok {
                val = []coordinate{coord}
            } else {
                val = append(val, coord)
            }
            antennas[c] = val
        }
    }

    // fmt.Println("Solution Part 1:", solvePart1(antennas, len(input) - 1, maxCol))
    fmt.Println("Solution Part 2:", solvePart2(antennas, len(input) - 1, maxCol))
}

func solvePart1(antennas map[rune][]coordinate, maxRow, maxCol int) int {
    uniqueAntiNodes := make(map[coordinate]bool)
    for _, coordinates := range antennas {
        for _, pair := range utils.AllPairs(coordinates) {
            dist := pair[1].diff(&pair[0])
            antiNode := pair[0].diff(&dist)
            if antiNode.row >= 0 && antiNode.row <= maxRow && antiNode.col >= 0 && antiNode.col <= maxCol {
                uniqueAntiNodes[antiNode] = true
            }
        }
    }
    return len(uniqueAntiNodes)
}

func solvePart2(antennas map[rune][]coordinate, maxRow, maxCol int) int {
    uniqueAntiNodes := make(map[coordinate]bool)
    for _, coordinates := range antennas {
        for _, pair := range utils.AllPairs(coordinates) {
            dist := pair[1].diff(&pair[0])
            antiNode := pair[0]
            for antiNode.row >= 0 && antiNode.row <= maxRow && antiNode.col >= 0 && antiNode.col <= maxCol {
                uniqueAntiNodes[antiNode] = true
                antiNode = antiNode.diff(&dist)
            }
        }
    }
    return len(uniqueAntiNodes)
}
