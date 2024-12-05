package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"

    "advent-of-code/internal/utils"
)

type rule struct {
    before int
    after int
}
type ruleSet []rule
type update []int

func (r ruleSet) breaksRule(page1, page2 int) bool {
    for _, rule := range r {
        if page1 == rule.after && page2 == rule.before {
            return true
        }
    }
    return false
}

func (u update) isCorrectOrder(rules ruleSet) bool {
    for i := 0; i < len(u) - 1; i++ {
        for j := i + 1; j < len(u); j++ {
            if rules.breaksRule(u[i], u[j]) {
                return false
            }
        }
    }
    return true
}

func (u update) fix(rules ruleSet) update {
    correctUpdate := make(update, len(u))
    if u.isCorrectOrder(rules) {
        copy(correctUpdate, u)
        return correctUpdate
    }
    for i := 0; i < len(u) - 1; i++ {
        for j := i + 1; j < len(u); j++ {
            if rules.breaksRule(u[i], u[j]) {
                copy(correctUpdate[:i], u[:i])
                correctUpdate[i] = u[j]
                copy(correctUpdate[i+1:j+1], u[i:j])
                copy(correctUpdate[j+1:], u[j+1:])
                return correctUpdate.fix(rules)
            }
        }
    }
    return correctUpdate
}

func main() {
    input, err := utils.ReadFile("input/day05.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    var rules ruleSet
    var updates []update
    isRules := true
    for _, line := range input {
        if line == "" {
            isRules = false
            continue
        }
        if isRules {
            nums := strings.Split(line, "|")
            before, err := strconv.Atoi(nums[0])
            if err != nil {
                fmt.Println("Error converting string to num")
                return
            }
            after, err := strconv.Atoi(nums[1])
            if err != nil {
                fmt.Println("Error converting string to num")
                return
            }
            rules = append(rules, rule{before: before, after: after})
        } else {
            nums := strings.Split(line, ",")
            var updt update
            for _, num := range nums {
                page, err := strconv.Atoi(num)
                if err != nil {
                    fmt.Println("Error converting string to num")
                    return
                }
                updt = append(updt, page)
            }
            updates = append(updates, updt)
        }
    }

    fmt.Println("Solution Part 1:", solvePart1(rules, updates))
    fmt.Println("Solution Part 2:", solvePart2(rules, updates))
}

func solvePart1(rules ruleSet, updates []update) int {
    sum := 0
    a := correctOrder(rules, updates)
    for _, r := range a {
        sum += r[len(r) / 2]
    }
    return sum
}

func solvePart2(rules ruleSet, updates []update) int {
    sum := 0
    a := incorrectOrder(rules, updates)
    for _, r := range a {
        sum += r[len(r) / 2]
    }
    return sum
}

func correctOrder(rules ruleSet, updates []update) []update {
    var correctUpdates []update
    for _, update := range updates {
        if update.isCorrectOrder(rules) {
            correctUpdates = append(correctUpdates, update)
        }
    }
    return correctUpdates
}

func incorrectOrder(rules ruleSet, updates []update) []update {
    var correctUpdates []update
    for _, update := range updates {
        if !update.isCorrectOrder(rules) {
            fixedUpdate := update.fix(rules)
            correctUpdates = append(correctUpdates, fixedUpdate)
        }
    }
    return correctUpdates
}
