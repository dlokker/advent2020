package main

import (
    "bufio"
    "fmt"
    "os"
    // "regexp"
    "strconv"
    // "strings"
    // "sync"
)

const Preamble = 25
var lookback = make(map[int]bool)
var xmas []int

func findBounds(l int, h int) (int, int) {
    min := xmas[l]
    max := xmas[l]
    for i := l; i < h; i++ {
        if xmas[i] < min {
            min = xmas[i]
        }
        if xmas[i] > max {
            max = xmas[i]
        }
    }
    return min, max
}

func findSum(n int) bool {
    for k, _ := range lookback {
        diff := n - k
        if _, ok := lookback[diff]; ok {
            return true
        }
    }
    return false
}

func main() {
    file, _ := os.Open("/Users/dlokk/code/advent2020/day9/input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        i, _ := strconv.Atoi(scanner.Text())
        xmas = append(xmas, i)
    }

    var p1 int
    for i := 0; i < len(xmas); i++ {
        lookback[xmas[i]] = true
        if i < Preamble {
            continue
        }
        n := xmas[i]
        if !findSum(n) {
            p1 = n
            fmt.Printf("Part 1: %v\n", p1)
            break
        }
        delete(lookback, xmas[i-Preamble])
    }

    for i := 0; i < len(xmas); i++ {
        rsum := xmas[i]
        for j := i+1; j < len(xmas); j++ {
            rsum += xmas[j]
            if rsum == p1 {
                low, high := findBounds(i, j)
                res := low + high
                fmt.Printf("Part 2: %v\n", res)
                return
            }
            if rsum > p1 {
                break
            }
        }
    }

}
