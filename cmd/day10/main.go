package main

import (
    "fmt"
    "log"

    "advent-of-code/internal/utils"
)

type topographicMap [][]int

type coordinate struct {
    row int
    col int
}

type direction int
const (
    up direction = 0
    down direction = 1
    right direction = 2
    left direction = 3
)

func (m topographicMap) height(position coordinate) int {
    return m[position.row][position.col]
}
func (m topographicMap) withinLimits(position coordinate) bool {
    return !(position.row < 0 || position.col < 0 || position.row >= len(m) || position.col >= len(m[0]))
}

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

func main() {
    input, err := utils.ReadFile("input/day10.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    var trailheads []coordinate
    map_ := make(topographicMap, len(input))
    for i, line := range input {
        row := make([]int, len(line))
        for j, c := range line {
            num := int(c - '0')
            if num == 0 {
                trailheads = append(trailheads, coordinate{row: i, col: j})
            }
            row[j] = num
        }
        map_[i] = row
    }

    // fmt.Println("Solution Part 1:", solvePart1(map_, trailheads))
    fmt.Println("Solution Part 2:", solvePart2(map_, trailheads))
}

func solvePart1(map_ topographicMap, trailheads []coordinate) int {
    sum := 0
    for _, trailhead := range trailheads {
        trails := make(map[coordinate]bool)
        findTrails(map_, trailhead, trails)
        score := len(trails)
        sum += score
    }
    return sum
}

func solvePart2(map_ topographicMap, trailheads []coordinate) int {
    sum := 0
    for _, trailhead := range trailheads {
        rating := 0
        findRating(map_, trailhead, &rating)
        sum += rating
    }
    return sum
}

func findTrails(map_ topographicMap, position coordinate, trails map[coordinate]bool) {
    height := map_.height(position)
    if height == 9 {
        trails[position] = true
        return
    }
    for _, d := range []direction{up, down, right, left} {
        nextPosition := position.step(d)
        if map_.withinLimits(nextPosition) {
            nextHeight := map_.height(nextPosition)
            if nextHeight == height + 1 {
                findTrails(map_, nextPosition, trails)
            }
        }
    }
}

func findRating(map_ topographicMap, position coordinate, score *int) {
    height := map_.height(position)
    if height == 9 {
        *score++
        return
    }
    for _, d := range []direction{up, down, right, left} {
        nextPosition := position.step(d)
        if map_.withinLimits(nextPosition) {
            nextHeight := map_.height(nextPosition)
            if nextHeight == height + 1 {
                findRating(map_, nextPosition, score)
            }
        }
    }
}
