package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    num_of_incr := 0
    i := 0
    var previous_num int

    // open file
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        current_num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        if i != 0 {
            if previous_num < current_num {
                num_of_incr++
            }
        }
        previous_num = current_num
        i++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println(num_of_incr)
}
