package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    // horizontal position
    var hor_pos, depth, aim = 0, 0, 0

    // open file
    file, err := os.Open("../input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        num, err := strconv.Atoi(line[1])
        if err != nil {
            log.Fatal(err)
        }
        if line[0] == "forward" {
            hor_pos += num 
            depth = depth + (aim * num)
        } else if line[0] == "down" {
            aim += num
        } else if line[0] == "up" {
            aim -= num
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println(hor_pos * depth)
}
