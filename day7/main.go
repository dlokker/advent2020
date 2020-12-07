package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func extractInt(s string) int {
    re := regexp.MustCompile("[0-9]+")
    cs := re.FindAllString(s, -1)
    i, _ := strconv.Atoi(cs[0])
    return i
}

func extractCount(rule string, bag string) int {
    r := fmt.Sprintf("(\\d+)\\s*%s",bag)
    re := regexp.MustCompile(r)
    match := re.FindStringSubmatch(rule)
    return extractInt(match[0])
}

var golds int

func getGolds(bag string, count int) {
    rule := rules[bag]
    golds += count
    for b, _ := range rules {
        if strings.Contains(rule, b){
            i := extractCount(rule, b)
            for j := 0; j < i; j++ {
                getGolds(b, 1)
            }
        }
    }
}

var rules = make(map[string]string)

func main() {
    file, _ := os.Open("/Users/dlokk/code/advent2020/day7/input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, "bags contain")
        bag := fields[0]
        contents := fields[1]
        rules[bag] = contents
    }

    var queue []string
    goldies := 0
    for b, c := range rules {
        if strings.Contains(c, "shiny gold") {
            goldies++
            queue = append(queue, b)
        }
    }
    visited := make(map[string]bool)
    for len(queue) > 0 {
        bag := queue[0]
        queue = queue[1:]
        for b, c := range rules {
            if strings.Contains(c, bag) {
                if _, ok := visited[b]; !ok {
                    visited[b] = true
                } else {
                    continue
                }
                goldies++
                queue = append(queue, b)
            }
        }
    }
    fmt.Printf("Goldies: %v\n", goldies)

    getGolds("shiny gold ", 0)
    fmt.Printf("Gold holds: %v\n", golds)
}
