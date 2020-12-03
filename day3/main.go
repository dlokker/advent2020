package main

import (
    "bufio"
    "fmt"
    "os"
    "sync"
)

type slope struct {
    rise, run int
}

func hits(s slope, grid []string, wg *sync.WaitGroup, i int, output *[]int) {
    defer wg.Done()

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
    fmt.Printf("Slope %d,%d hits %d trees\n",s.rise, s.run, trees)
    (*output)[i] = trees
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

    var wg sync.WaitGroup
    output := make([]int, 5)
    for i, s := range slopes {
        wg.Add(1)
        go hits(s, grid, &wg, i, &output)
    }
    wg.Wait()

    res := 1
    for _, r := range output {
        res = res * r
    }
    fmt.Printf("Result: %v\n", res)
}
