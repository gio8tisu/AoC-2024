package main

import (
    "fmt"
    "log"
    "regexp"
    "strconv"

    "advent-of-code/internal/utils"
)

const (
    buttonACost = 3
    buttonBCost = 1
    budget = 2
)

type solution struct {
    buttonAPresses int
    buttonBPresses int
}

func (s solution) cost() int {
    if s.buttonAPresses == -1 || s.buttonBPresses == -1 {
        return 0
    }
    return s.buttonAPresses * buttonACost + s.buttonBPresses * buttonBCost
}

type coordinate struct {
    row int
    col int
}

type button struct {
    rowStep int
    colStep int
}

type machine struct {
    buttonA button
    buttonB button
    prize coordinate
}

func (m *machine) findSolution() solution {
	bNumerator := m.prize.row*m.buttonA.colStep - m.prize.col*m.buttonA.rowStep
	bDenominator := m.buttonB.rowStep*m.buttonA.colStep - m.buttonB.colStep*m.buttonA.rowStep

	if bDenominator != 0 && bNumerator%bDenominator == 0 {
		b := bNumerator / bDenominator
		aNumerator := m.prize.col - b*m.buttonB.colStep
		if m.buttonA.colStep != 0 && aNumerator%m.buttonA.colStep == 0 {
			a := aNumerator / m.buttonA.colStep
			return solution{buttonAPresses: a, buttonBPresses: b}
		}
	}
	return solution{buttonAPresses: -1, buttonBPresses: -1}
}

func main() {
    input, err := utils.ReadFile("input/day13.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

	reButton := regexp.MustCompile(`Button [AB]: X\+(\d+), Y\+(\d+)`)
	rePrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
    var machines []machine
    var m machine = machine{}
    i := 0
    for _, line := range input {
        if line == "" {
            m = machine{}
            i = 0
            continue
        }
        switch i {
        case 0:
	        matches := reButton.FindAllStringSubmatch(line, -1)
            colStep, err := strconv.Atoi(matches[0][1])
            rowStep, err := strconv.Atoi(matches[0][2])
            if err != nil {
                return
            }
            m.buttonA = button{rowStep: rowStep, colStep: colStep}
        case 1:
	        matches := reButton.FindAllStringSubmatch(line, -1)
            colStep, err := strconv.Atoi(matches[0][1])
            rowStep, err := strconv.Atoi(matches[0][2])
            if err != nil {
                return
            }
            m.buttonB = button{rowStep: rowStep, colStep: colStep}
        case 2:
	        matches := rePrize.FindAllStringSubmatch(line, -1)
            col, err := strconv.Atoi(matches[0][1])
            row, err := strconv.Atoi(matches[0][2])
            if err != nil {
                return
            }
            m.prize = coordinate{row: row, col: col}
            machines = append(machines, m)
        }
        i++
    }

    fmt.Println("Solution Part 1:", solvePart1(machines))
    fmt.Println("Solution Part 2:", solvePart2(machines))
}

func solvePart1(machines []machine) int {
    cost := 0
    for _, m := range machines {
        solution := m.findSolution()
        fmt.Println(solution)
        cost += solution.cost()
    }
    return cost
}

func solvePart2(machines []machine) int {
    cost := 0
    for _, m := range machines {
        m.prize.col += 10000000000000
        m.prize.row += 10000000000000
        solution := m.findSolution()
        fmt.Println(solution)
        cost += solution.cost()
    }
    return cost
}
