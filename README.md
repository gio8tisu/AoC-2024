# Advent of Code in Go

This repository contains solutions for [Advent of Code](https://adventofcode.com/) implemented in Go.

Most of the solutions are quick-and-dirty implementations so expect shortcuts and don't expect optimized, clean or idiomatic code.

## Structure

- **`cmd/`**: Day-specific solutions.
- **`internal/utils/`**: Shared utility functions.
- **`input/`**: Puzzle inputs.

## Usage

1. Place the input for a specific day in the `input/` folder (e.g., `input/day01.txt`).
2. Run the solution for that day, e.g.:
   ```bash
   go run ./cmd/day01
   ```
   or
   ```bash
   just solve 01
   ```
