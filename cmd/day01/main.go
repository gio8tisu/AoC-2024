package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"
    "slices"

    "advent-of-code/internal/utils"
)

func main() {
    input, err := utils.ReadFile("input/day01.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    var num1, num2 uint64
    var nums []string
    var list1 = make([]uint64, len(input))
    var list2 = make([]uint64, len(input))
    for i, line := range input {
        nums = strings.Split(line, "   ")
        num1, err = strconv.ParseUint(nums[0], 10, 64)
        if err != nil {
            fmt.Println("Error converting string to num")
        }
        num2, err = strconv.ParseUint(nums[1], 10, 64)
        if err != nil {
            fmt.Println("Error converting string to num")
        }
        list1[i] = num1
        list2[i] = num2
    }
    slices.Sort(list1)
    slices.Sort(list2)

    fmt.Println("Solution Part 1:", solvePart1(list1, list2))
    fmt.Println("Solution Part 2:", solvePart2(list1, list2))
}

func solvePart1(list1, list2 []uint64) uint64 {
    // If there were repeated numbers, I'm not sure this pairing would work.
    var distance uint64
    for i := 0; i < len(list1); i++ {
        distance += utils.AbsDiffUint(list1[i], list2[i])
    }
    return distance
}

func solvePart2(list1, list2 []uint64) uint64 {
    var j, score uint64
    for i := 0; i < len(list1); i++ {
        count := uint64(0)
        for list2[j] <= list1[i] {
            if list2[j] == list1[i] {
                count++
            }
            j++
        }
        score += list1[i] * count
    }
    return score
}
