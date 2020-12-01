package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("/Users/dlokk/code/advent2020/day1/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    entries := make([]int, 0)
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        n, _ := strconv.Atoi(scanner.Text())
        entries = append(entries, n)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for i, m1 := range entries {
        for j := i+1; j < len(entries); j++ {
            m2 := entries[j]
            if (m1 + m2) == 2020 {
                fmt.Printf("match: %d + %d = 2020\n", m1, m2)
                fmt.Printf("product: %d\n", m1*m2)
            }
            for k := j+1; k < len(entries); k++ {
                m3 := entries[k]
                if (m1 + m2 + m3) == 2020 {
                    fmt.Printf("match: %d + %d + %d = 2020\n", m1, m2, m3)
                    fmt.Printf("product: %d\n", m1*m2*m3)
                    return
                }
            }
        }
    }
}