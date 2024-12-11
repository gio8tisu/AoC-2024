package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"

    "advent-of-code/internal/utils"
)

type stone struct {
    num int
    link *stone
}

func (s *stone) numDigits() int {
    if s.num == 0 {
        return 1
    }
    count := 0
    num := s.num
    for num != 0 {
        num /= 10
        count++
    }
    return count
}

func (s *stone) split() *stone {
    n := s.numDigits()
    j := 1
    num := s.num
    for i := 0; i < n / 2; i++ {
        j *= 10
    }
    s.num = num / j
    s2 := &stone{num: num % j, link: s.link}
    s.link = s2
    return s2
}

func (s1 *stone) length() int {
    i := 1
    s2 := s1
    for s2.link != nil {
        i++
        s2 = s2.link
    }
    return i
}

func (s1 *stone) blink() {
    s2 := s1
    for s2.link != nil {
        if s2.num == 0 {
            s2.num = 1
            s2 = s2.link
        } else if s2.numDigits() % 2 == 0 {
            s2 = s2.split()
            s2 = s2.link
        } else {
            s2.num = s2.num * 2024
            s2 = s2.link
        }
    }
    if s2.num == 0 {
        s2.num = 1
    } else if s2.numDigits() % 2 == 0 {
        s2 = s2.split()
    } else {
        s2.num = s2.num * 2024
    }
}

func main() {
    input, err := utils.ReadFile("input/day11.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    nums := strings.Split(input[0], " ")
    num, err := strconv.Atoi(nums[len(nums) - 1])
    if err != nil {
        fmt.Println("Error converting string to num")
        return
    }
    s := &stone{num: num}
    for i := len(nums) - 2; i >= 0; i-- {
        num, err = strconv.Atoi(nums[i])
        if err != nil {
            fmt.Println("Error converting string to num")
            return
        }
        s = &stone{num: num, link: s}
    }
    fmt.Println("Solution Part 1:", solvePart1(s))
    // fmt.Println("Solution Part 2:", solvePart2(s))
}

func solvePart1(s *stone) int {
    for i := 0; i < 25; i++ {
        s.blink()
    }
    return s.length()
}

func solvePart2(s *stone) int {
    return 0
}
