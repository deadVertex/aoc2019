package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "math"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func calculateFuelRecursion(mass int) int {
    fuel := int(math.Floor(float64(mass) / 3.0) - 2)
    if fuel > 0 {
        fuel += calculateFuelRecursion(fuel)
    } else {
        fuel = 0
    }

    return fuel
}

func main() {
    dat, err := ioutil.ReadFile("input.txt")
    check(err)
    input := string(dat)
    lines := strings.Split(input, "\n")

    // Part 1
    total := 0
    for _, line := range lines {

        // Handle blank lines
        if len(line) > 0 {
            mass, err := strconv.Atoi(line)
            check(err)
            fuel := calculateFuelRecursion(mass)//int(math.Floor(float64(mass) / 3.0) - 2)
            total += fuel
        }
    }
    fmt.Println(total)

    // Part 2
    total += calculateFuelRecursion(total)
    fmt.Println(total)
}
