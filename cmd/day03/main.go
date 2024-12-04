package main


import (
    "fmt"
    "log"
    "strconv"
	"regexp"

    "advent-of-code/internal/utils"
)

func main() {
    input, err := utils.ReadFile("input/day03.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    fmt.Println("Solution Part 1:", solvePart1(input))
    fmt.Println("Solution Part 2:", solvePart2(input))
}

func solvePart1(lines []string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	result := 0
	for _, line := range lines {
	    matches := re.FindAllStringSubmatch(line, -1)
	    for _, match := range matches {
            result += mul(match[1], match[2])
	    }
	}
    return result
}

func solvePart2(lines []string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|(don't\(\))|(do\(\))`)
	result := 0
	do := true
	for _, line := range lines {
	    matches := re.FindAllStringSubmatch(line, -1)
	    for _, match := range matches {
	        if match[3] != "" {
	            do = false
	            continue
	        }
	        if match[4] != "" {
	            do = true
	            continue
	        }
	        if do {
                result += mul(match[1], match[2])
	        }
	    }
	}
    return result
}

func mul(numStr1, numStr2 string) int {
    num1, err := strconv.Atoi(numStr1)
    if err != nil {
        fmt.Println("Error converting string to num")
    }
    num2, err := strconv.Atoi(numStr2)
    if err != nil {
        fmt.Println("Error converting string to num")
    }
    return num1 * num2
}
