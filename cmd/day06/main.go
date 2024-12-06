package main

import (
    "fmt"
    "log"
    // "strings"
    // "strconv"

    "advent-of-code/internal/utils"
)

type coordinate struct {
    x int
    y int
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
        return coordinate{x: c.x, y: c.y - 1}
    case down:
        return coordinate{x: c.x, y: c.y + 1}
    case right:
        return coordinate{x: c.x + 1, y: c.y}
    case left:
        return coordinate{x: c.x - 1, y: c.y}
    default:
        // Invalid
        return c
    }
}

func (d direction) turn() direction {
    switch d {
    case up:
        return right
    case down:
        return left
    case right:
        return down
    case left:
        return up
    default:
        // Invalid
        return d
    }
}
func (d direction) walk(position coordinate, map_ []string, p map[coordinate]bool) (coordinate, bool) {
    nextPosition := position.step(d)
    for nextPosition.x >= 0 && nextPosition.x <= len(map_) - 1 && nextPosition.y >= 0 && nextPosition.y <= len(map_[0]) - 1 && map_[nextPosition.y][nextPosition.x] != '#' {
        p[position] = true
        position = nextPosition
        nextPosition = position.step(d)
    }
    p[position] = true
    return position, nextPosition.x < 0 || nextPosition.x > len(map_) - 1 || nextPosition.y < 0 || nextPosition.y > len(map_[0]) - 1
}

func main() {
    input, err := utils.ReadFile("input/day06.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    fmt.Println("Solution Part 1:", solvePart1(input))
    // fmt.Println("Solution Part 2:", solvePart2(input))
}

func solvePart1(map_ []string) int {
    positions := map[coordinate]bool{}
    var position coordinate
    var dir direction
    out:
    for i, line := range map_ {
        for j, char := range line {
            if d := direction(char); d == up || d == down || d == right || d == left {
                fmt.Println("found start position", i, j)
                position = coordinate{x: j, y: i}
                dir = d
                break out
            }
        }
    }
    for end := false; !end; dir = dir.turn() {
        fmt.Println(position)
        position, end = dir.walk(position, map_, positions)
        
    }
    return len(positions)
}

func solvePart2(map_ []string) int {
    return 0
}
