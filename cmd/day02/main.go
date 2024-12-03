package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"

    "advent-of-code/internal/utils"
)

func main() {
    input, err := utils.ReadFile("input/day02.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    reports := make([][]int, len(input))
    for i, line := range input {
        levels := strings.Split(line, " ")
        report := make([]int, len(levels))
        for j, level := range levels {
            num, err := strconv.Atoi(level)
            if err != nil {
                fmt.Println("Error converting string to num")
            }
            report[j] = num
        }
        reports[i] = report
    }

    fmt.Println("Solution Part 1:", solvePart1(reports))
    fmt.Println("Solution Part 2:", solvePart2(reports))
}

func solvePart1(reports [][]int) int {
    var safeCount int
    for _, report := range reports {
        if isSafe(report) {
            safeCount++
        }
    }
    return safeCount
}

func solvePart2(reports [][]int) int {
    var safeCount int
    for _, report := range reports {
        if isSafe2(report) {
            safeCount++
        }
    }
    return safeCount
}

func isSafe(report []int) bool {
    var ascending bool
    switch {
    case report[1] - report[0] > 0:
        ascending = true
    case report[1] - report[0] < 0:
        ascending = false
    default:
        return false
    }
    for i := 1; i < len(report); i++ {
        switch {
        case report[i] - report[i - 1] > 0 && report[i] - report[i - 1] < 4: // Is ascending
            if !ascending {
                return false
            }
        case report[i] - report[i - 1] < 0 && report[i] - report[i - 1] > -4: // Is descending
            if ascending {
                return false
            }
        default:
            return false
        }
    }
    return true
}

func isSafe2(report []int) bool {
    if isSafe(report) {
        return true
    }
    var s = make([]int, len(report) - 1)
    for i := 0; i < len(report); i++ {
        copy(s[:i], report[:i])
        copy(s[i:], report[i+1:])
        safe := isSafe(s)
        if safe {
            return true
        }
    }
    return false
}
