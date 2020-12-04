package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "reflect"
    "regexp"
)

type passport struct {
    Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid string
}

func buildPass(p *passport, fields []string) {
    for _, field := range fields {
        fs := strings.Split(field, ":")
        f := strings.Title(fs[0])
        v := fs[1]
        r := reflect.ValueOf(p)
        reflect.Indirect(r).FieldByName(f).SetString(v)
    }
}

func checkValid(p *passport) bool {
    r := reflect.ValueOf(p)
    for _, f := range constFields {
        v := reflect.Indirect(r).FieldByName(f).String()
        if v == "" {
            return false
        }
        switch f {
        case "Byr":
            n, err := strconv.Atoi(v)
            if len(v) != 4 || n < 1920 || n > 2002 || err != nil {
                return false
            }
        case "Iyr":
            n, err := strconv.Atoi(v)
            if len(v) != 4 || n < 2010 || n > 2020 || err != nil {
                return false
            }
        case "Eyr":
            n, err := strconv.Atoi(v)
            if len(v) != 4 || n < 2020 || n > 2030 || err != nil {
                return false
            }
        case "Hgt":
            n, _ := strconv.Atoi(v[:len(v)-2])
            unit := v[len(v)-2:]
            if unit == "cm" {
                if n < 150 || n > 193 {
                    return false
                }
            } else if unit == "in" {
                if n < 59 || n > 76 {
                    return false
                }
            } else {
                return false
            }
        case "Hcl":
            if match, _ := regexp.MatchString("[#][a-f0-9]{6}", v) ; !match {
                return false
            }
        case "Ecl":
            a := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
            found := false
            for _, c := range a {
                if v == c {
                    found = true
                }
            }
            if !found {
                return false
            }
        case "Pid":
            _, err := strconv.Atoi(v)
            if len(v) != 9 || err != nil {
                return false
            }
        default:
            fmt.Printf("invalid field: %v\n", f)
        }
    }
    return true
}

var constFields [7]string

func main() {
    file, _ := os.Open("/Users/dlokk/code/advent2020/day4/input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    constFields = [...]string{"Byr", "Iyr", "Eyr", "Hgt", "Hcl", "Ecl", "Pid"}
    valid := 0
    p := passport{}
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        buildPass(&p, fields)
        if line == "" {
            if checkValid(&p) {
                valid++
            }
            p = passport{}
        }
    }
    fmt.Printf("Valid passports: %d\n", valid)
}
