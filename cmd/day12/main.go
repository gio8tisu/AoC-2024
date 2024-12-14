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

type direction rune
const (
    up direction = '^'
    down direction = 'v'
    right direction = '>'
    left direction = '<'
)

func (c coordinate) step(d direction) coordinate {
    switch d {
    case up:
        return coordinate{row: c.row, col: c.col - 1}
    case down:
        return coordinate{row: c.row, col: c.col + 1}
    case right:
        return coordinate{row: c.row + 1, col: c.col}
    case left:
        return coordinate{row: c.row - 1, col: c.col}
    default:
        // Invalid
        return c
    }
}

type region struct {
    positions map[coordinate]bool
    perimeter int
} 

func (r *region) add(c coordinate) {
    r.positions[c] = true
}

func (r *region) price() int {
    return r.perimeter * len(r.positions)
}

func (r *region) findRegion(label rune, c1 coordinate, input []string, seen map[coordinate]bool) {
    if (c1.row < 0 || c1.col < 0 || c1.row >= len(input) || c1.col >= len(input[0])) {
        r.perimeter++
        return
    }
    if []rune(input[c1.row])[c1.col] != label {
        r.perimeter++
        return
    }
    if seen[c1] {
        return
    }
    seen[c1] = true
    r.add(c1)
    for _, d := range []direction{up, right, down, left} {
        c2 := c1.step(d)
        r.findRegion(label, c2, input, seen)
    }
}

func main() {
    input, err := utils.ReadFile("input/day12.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    var map_ []region
    seen := make(map[coordinate]bool)
    for i, line := range input {
        for j, run := range line {
            c := coordinate{row: i, col: j}
            if seen[c] {
                continue
            }
            r := region{positions: make(map[coordinate]bool)}
            r.findRegion(run, c, input, seen)
            map_ = append(map_, r)
        }
    }

    fmt.Println("Solution Part 1:", solvePart1(map_))
    // fmt.Println("Solution Part 2:", solvePart2(map_))
}

func solvePart1(map_ []region) int {
    totalPrice := 0
    for _, region := range map_ {
        totalPrice += region.price()
    }
    return totalPrice
}

func solvePart2(map_ []string) int {
    return 0
}
