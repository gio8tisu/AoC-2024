package main

import (
    "fmt"
    "log"

    "advent-of-code/internal/utils"
)

func main() {
    input, err := utils.ReadFile("input/day04.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    // fmt.Println("Solution Part 1:", solvePart1(input))
    fmt.Println("Solution Part 2:", solvePart2(input))
}

func solvePart1(wordSearch []string) int {
    count := 0
    for i := 0; i < len(wordSearch); i++ {
        for j := 0; j < len(wordSearch[i]); j++ {
            // Left-to-right
            if j <= len(wordSearch[i]) - 4 && wordSearch[i][j] == 'X' {
                if wordSearch[i][j + 1] == 'M' {
                    if wordSearch[i][j + 2] == 'A' {
                        if wordSearch[i][j + 3] == 'S' {
                            count++
                        }
                    }
                }
            }
            // Right-to-left
            if j >= 3 && wordSearch[i][j] == 'X' {
                if wordSearch[i][j - 1] == 'M' {
                    if wordSearch[i][j - 2] == 'A' {
                        if wordSearch[i][j - 3] == 'S' {
                            count++
                        }
                    }
                }
            }
            // Top-to-bottom
            if i <= len(wordSearch) - 4 && wordSearch[i][j] == 'X' {
                if wordSearch[i + 1][j] == 'M' {
                    if wordSearch[i + 2][j] == 'A' {
                        if wordSearch[i + 3][j] == 'S' {
                            count++
                        }
                    }
                }
            }
            // Bottom-to-top
            if i >= 3 && wordSearch[i][j] == 'X' {
                if wordSearch[i - 1][j] == 'M' {
                    if wordSearch[i - 2][j] == 'A' {
                        if wordSearch[i - 3][j] == 'S' {
                            count++
                        }
                    }
                }
            }
            // NE diagonal
            if j <= len(wordSearch[i]) - 4 && i >= 3 && wordSearch[i][j] == 'X' {
                if wordSearch[i - 1][j + 1] == 'M' {
                    if wordSearch[i - 2][j + 2] == 'A' {
                        if wordSearch[i - 3][j + 3] == 'S' {
                            count++
                        }
                    }
                }
            }
            // SE diagonal
            if j <= len(wordSearch[i]) - 4 && i <= len(wordSearch) - 4 && wordSearch[i][j] == 'X' {
                if wordSearch[i + 1][j + 1] == 'M' {
                    if wordSearch[i + 2][j + 2] == 'A' {
                        if wordSearch[i + 3][j + 3] == 'S' {
                            count++
                        }
                    }
                }
            }
            // NW diagonal
            if j >= 3 && i >= 3 && wordSearch[i][j] == 'X' {
                if wordSearch[i - 1][j - 1] == 'M' {
                    if wordSearch[i - 2][j - 2] == 'A' {
                        if wordSearch[i - 3][j - 3] == 'S' {
                            count++
                        }
                    }
                }
            }
            // SW diagonal
            if j >= 3 && i <= len(wordSearch) - 4 && wordSearch[i][j] == 'X' {
                if wordSearch[i + 1][j - 1] == 'M' {
                    if wordSearch[i + 2][j - 2] == 'A' {
                        if wordSearch[i + 3][j - 3] == 'S' {
                            count++
                        }
                    }
                }
            }
        }
    }
    return count
}

func solvePart2(wordSearch []string) int {
    count := 0
    for i := 0; i < len(wordSearch); i++ {
        for j := 0; j < len(wordSearch[i]); j++ {
            oneDiag := false
            if i >= 1 && j >= 1 && i <= len(wordSearch) - 2 && j <= len(wordSearch[i]) - 2 && wordSearch[i][j] == 'A' {
                if wordSearch[i - 1][j - 1] == 'M' {
                    if wordSearch[i + 1][j + 1] == 'S' {
                        oneDiag = true
                    }
                } else if wordSearch[i - 1][j - 1] == 'S' {
                    if wordSearch[i + 1][j + 1] == 'M' {
                        oneDiag = true
                    }
                }
                if !oneDiag {
                    continue
                }
                if wordSearch[i - 1][j + 1] == 'M' {
                    if wordSearch[i + 1][j - 1] == 'S' {
                        count++
                    }
                } else if wordSearch[i - 1][j + 1] == 'S' {
                    if wordSearch[i + 1][j - 1] == 'M' {
                        count++
                    }
                }
            }
        }
    }
    return count
}
