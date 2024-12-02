package utils

import "strconv"

// ToInt converts a string to an integer, returning 0 if there's an error.
func ToInt(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        return 0
    }
    return i
}
