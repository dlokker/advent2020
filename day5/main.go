package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func calcId(s string) uint64 {
    row := ""
    col := ""
    for _, c := range s {
        switch string(c) {
        case "F":
            row += "0"
        case "B":
            row += "1"
        case "R":
            col += "1"
        case "L":
            col += "0"
        default:
            fmt.Printf("invalid input: %v\n", c)
        }
    }
    ri, _ := strconv.ParseUint(row, 2, 32)
    ci, _ := strconv.ParseUint(col, 2, 32)

    if _, ok := seats[ri]; !ok {
        seats[ri] = make([]bool, 8)
    }
    seats[ri][ci] = true
    
    res := (ri * 8) + ci
    return res
}

var seats = make(map[uint64][]bool)

func main() {
    file, _ := os.Open("/Users/dlokk/code/advent2020/day5/input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    ids := make(map[uint64]bool)
    var max uint64
    for scanner.Scan() {
        line := scanner.Text()
        id := calcId(line)
        ids[id] = true
        if id > max {
            max = id
        }
    }
    fmt.Printf("Max ID: %v\n", max)

    for k, v := range seats{
        for i, s := range v {
            if s == false {
                // fmt.Printf("Seat %v,%v is empty\n",k,i)
                id := (k*8)+uint64(i)
                if _, ok := ids[id+1]; ok {
                    if _, okk := ids[id-1]; okk {
                        fmt.Printf("ID %v is probably your seat\n", id)
                        return
                    }
                }
            }
        }
    }
}
