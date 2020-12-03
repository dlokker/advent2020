package main

import (
    "bufio"
    "fmt"
    "os"
)

type slope struct {
    rise, run int
}

func main() {
    file, _ := os.Open("/Users/dlokk/code/advent2020/day3/input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var grid []string
    for scanner.Scan() {
        grid = append(grid, scanner.Text())
    }
    
    var slopes = []slope { 
        slope { rise: 1, run: 1, },
        slope { rise: 1, run: 3, },
        slope { rise: 1, run: 5, },
        slope { rise: 1, run: 7, },
        slope { rise: 2, run: 1, },
    }

    result := 1
    for _, s := range slopes {
        var x, y, trees int
        for {
            x += s.run
            y += s.rise
            if y >= len(grid) {
                break
            }
            if x >= len(grid[y]) {
                x = x - len(grid[y])
            }
            if grid[y][x] == '#' {
                trees++
            }
        }
        fmt.Printf("Slope %d,%d hits %d trees\n",s.run, s.rise, trees)
        result = result * trees
    }
    fmt.Printf("Result: %d\n", result)
}