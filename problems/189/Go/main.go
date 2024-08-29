package main

import (
    "fmt"
)

func main() {
    slice := []int{1,2,3,4,5,6,7}
    r := 3
    rotate(slice, r)
}

func rotate(nums []int, k int) {
    fmt.Println(nums, k)
}
