package main

import (
    "bufio"
    "fmt"
    "os"
    // "regexp"
    "strconv"
    "strings"
    "sync"
)

var acc int
var accs []int
var ins []inst

type inst struct {
    a, op string
    x int
}

func parseIn(s string) inst {
    buf := strings.Split(s, " ")
    a := buf[0]
    op := string(buf[1][:1])
    x, _ := strconv.Atoi(buf[1][1:])
    
    return inst{ a: a, op: op, x: x }
}

func evalIn(i int, j *int, ac *int) int {
    t := ins[i]
    a := t.a
    if j != nil && *j == i {
        if t.a == "nop" {
            a = "jmp"
        }
        if t.a == "jmp" {
            a = "nop"
        }
    }
    switch a {
    case "nop":
        return i+1
    case "acc":
        if t.op == "+" {
            *ac += t.x
        } else {
            *ac -= t.x
        }
        return i+1
    case "jmp":
        if t.op == "+" {
            i += t.x
        } else {
            i -= t.x
        }
        return i
    default:
        fmt.Printf("unknown instruction: %v\n", t)
    }
    return -1
}

func swapLoop(wg *sync.WaitGroup, j int, ac *int) {
    defer wg.Done()
    visited := make(map[int]bool)
    var i int
    for {
        if _, ok := visited[i]; ok {
            // fmt.Printf("Loop encountered on swap: %v\n", j)
            return
        }
        visited[i] = true
        i = evalIn(i, &j, ac)
        if i >= len(ins) || i < 0 {
            fmt.Printf("Part 2 Swap %v completed with acc = %v\n", j, *ac)
            return
        }
    }
    return
}

func main() {
    file, _ := os.Open("/Users/dlokk/code/advent2020/day8/input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var swaps []int
    for scanner.Scan() {
        ins = append(ins, parseIn(scanner.Text()))
    }
    for i, v := range ins {
        if v.a == "nop" || v.a == "jmp" {
            swaps = append(swaps, i)
        }
    }

    // Part 1
    visited := make(map[int]bool)
    var i int
    for {
        if _, ok := visited[i]; ok {
            fmt.Printf("Part 1 Accumulator: %v\n", acc)
            break
        }
        visited[i] = true
        i = evalIn(i, nil, &acc)
    }

    // Part 2
    accs = make([]int, len(swaps))
    var wg sync.WaitGroup
    for i, j := range swaps {
        wg.Add(1)
        go swapLoop(&wg, j, &accs[i])
    }
    wg.Wait()

}
