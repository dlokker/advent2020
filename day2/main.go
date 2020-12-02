package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    arg := os.Args[1]

    file, _ := os.Open("/Users/dlokk/code/advent2020/day2/input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    valid := 0
    for scanner.Scan() {
        elements := strings.Split(scanner.Text(), " ")
        bounds   := strings.Split(elements[0], "-")
        lower, _ := strconv.Atoi(bounds[0])
        upper, _ := strconv.Atoi(bounds[1])
        char     := strings.Trim(elements[1], ":")
        pass     := elements[2]

        switch arg {
        case "1":
            n := strings.Count(pass, char)
            if n >= lower && n <= upper {
                fmt.Printf("Valid: %v\n", elements)
                valid++
            }
        case "2":
            a := string(pass[lower-1]) == char
            b := string(pass[upper-1]) == char
            if (a || b) && !(a && b) {
                fmt.Printf("Valid: %v\n", elements)
                valid++
            }
        default:
            fmt.Println("invalid input: ", arg)
            return
        }

    }
    fmt.Printf("Valid passwords: %d\n", valid)
}