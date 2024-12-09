package main

import (
    "fmt"
    "log"

    "advent-of-code/internal/utils"
)

type block struct {
    free bool
    fileId int
}
type disk []block

func (d disk) String() string {
    a := ""
    for _, b := range d {
        if b.free {
            a += fmt.Sprintf(".")
        } else {
            a += fmt.Sprintf("%v", b.fileId)
        }
    }
    return a
}

func (d disk) checksum() int {
    a := 0
    for i, b := range d {
        if !b.free {
            a += i * b.fileId
        }
    }
    return a
}

func main() {
    input, err := utils.ReadFile("input/day09.txt")
    if err != nil {
        log.Fatalf("Failed to read input file: %v", err)
    }

    var diskMap disk
    fileId := 0
    for i, c := range input[0] {
        isFile := i % 2 == 0
        size := int(c - '0')
        var b block
        if isFile {
            b = block{free: !isFile, fileId: fileId}
        } else {
            b = block{free: !isFile, fileId: -1}
        }
        for j := 0; j < size; j++ {
            diskMap = append(diskMap, b)
        }
        if isFile {
            fileId++
        }
    }

    fmt.Println("Solution Part 1:", solvePart1(diskMap))
    // fmt.Println("Solution Part 2:", solvePart2(reports))
}

func solvePart1(diskMap disk) int {
    i := 0
    j := len(diskMap) - 1
    // Find first free block
    for ; !diskMap[i].free; i++ {}
    // Find last non-free block
    for ; diskMap[j].free; j-- {}
    for i < j {
        diskMap[i] = diskMap[j]
        diskMap[j] = block{free: true, fileId: -1}
        // Find first free block
        for ; !diskMap[i].free; i++ {}
        // Find last non-free block
        for ; diskMap[j].free; j-- {}
    }
    return diskMap.checksum()
}

func solvePart2(diskMap disk) int {
    return 0
}
