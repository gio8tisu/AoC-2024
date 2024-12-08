package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"

    "advent-of-code/internal/utils"
)

type operator rune
func (o operator) do(a, b int) int {
    if o == '+' {
        return a + b
    }
    if o == '*' {
        return a * b
    }
    if o == '|' {
	    str1 := strconv.Itoa(a)
	    str2 := strconv.Itoa(b)
	    concatenatedStr := str1 + str2
	    concatenatedInt, _ := strconv.Atoi(concatenatedStr)
	    return concatenatedInt
    }
    return 0
}

type equation struct {
    result int
    operands []int
}

func main() {
    input, err := utils.ReadFile("input/day07.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    equations := make([]equation, len(input))
    for i, line := range input {
        nums := strings.Split(line, " ")
        resultStr := strings.TrimSuffix(nums[0], ":")
        operands := make([]int, len(nums) - 1)
        for j, numStr := range nums[1:] {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                fmt.Println("Error converting string to num")
            }
            operands[j] = num
        }
        result, err := strconv.Atoi(resultStr)
        if err != nil {
            fmt.Println("Error converting string to num")
        }
        equations[i] = equation{result: result, operands: operands}
    }

    fmt.Println("Solution Part 1:", solvePart1(equations))
    fmt.Println("Solution Part 2:", solvePart2(equations))
}

func solvePart1(equations []equation) int {
    var operators = []operator{'+', '*'}
    sum := 0
    for _, equation := range equations {
        if isSolvable(&equation, operators) {
            sum += equation.result
        }
    }
    return sum
}

func solvePart2(equations []equation) int {
    var operators = []operator{'+', '*', '|'}
    sum := 0
    for _, equation := range equations {
        if isSolvable(&equation, operators) {
            sum += equation.result
        }
    }
    return sum
}

func isSolvable(e *equation, operators []operator) bool {
    if e.operands[0] > e.result {
        return false
    }
    if len(e.operands) == 1 {
        if e.operands[0] == e.result {
            return true
        } else {
            return false
        }
    }
    for _, o := range operators {
        firstOperand := o.do(e.operands[0], e.operands[1])
        operands := []int{firstOperand}
        operands = append(operands, e.operands[2:]...)
        newEq := equation{result: e.result, operands: operands}
        if isSolvable(&newEq, operators) {
            return true
        }
    }
    return false
}
