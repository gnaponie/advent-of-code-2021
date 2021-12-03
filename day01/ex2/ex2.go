package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func sum_array(array []int) int {
    result := 0
    for _, v := range array {
        result += v
    }
    return result
}

func main() {
    var num_of_incr, i, current_sum, previous_sum = 0, 0, 0, 0
    var three_nums []int

    // open file
    file, err := os.Open("../input.txt")
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
        if i < 3 {
            three_nums = append(three_nums, current_num)
            previous_sum += current_num
        } else {
            // remove first element and add the current
            three_nums = three_nums[1:]
            three_nums = append(three_nums, current_num)

            current_sum = sum_array(three_nums)

            if previous_sum < current_sum {
                num_of_incr++
            }
            previous_sum = current_sum
        }
        i++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println(num_of_incr)
}
