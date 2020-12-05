package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func calcId(s string) uint64 {
    row := strings.Replace(s[0:7], "F", "0", -1)
    row = strings.Replace(row, "B", "1", -1)
    col := strings.Replace(s[len(s)-3:], "L", "0", -1)
    col = strings.Replace(col, "R", "1", -1)

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
