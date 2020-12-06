package main

import (
    "bufio"
    "fmt"
    "os"
)

func countUniq(s string) int {
    uniq := make(map[rune]bool)
    count := 0
    for _, c := range s {
        if _, ok := uniq[c]; !ok {
            uniq[c] = true
            count++
        }
    }
    return count
}

func countAll(s string, p int) int {
    uniq := make(map[rune]int)
    count := 0
    for _, c := range s {
        if _, ok := uniq[c]; !ok {
            uniq[c] = 1
        } else {
            uniq[c]++
        }
    }
    for _, v := range uniq {
        if v == p {
            count++
        }
    }
    return count
}

func main() {
    file, _ := os.Open("/Users/dlokk/code/advent2020/day6/input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var res []int
    var res2 []int
    group := ""
    people := 0
    for scanner.Scan() {
        tmp := scanner.Text()
        if tmp == "" {
            res = append(res, countUniq(group))
            res2 = append(res2, countAll(group, people))
            group = ""
            people = 0
        } else {
            group += tmp
            people++
        }
    }
    sum := 0
    for _, r := range res {
        sum += r
    }
    fmt.Printf("sum1: %v\n", sum)
    sum2 := 0
    for _, r2 := range res2 {
        sum2 += r2
    }
    fmt.Printf("sum2: %v\n", sum2)

}
