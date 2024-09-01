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
    nums2 := make([]int, len(nums))
    copy(nums2, nums)
    for id, val := range nums2 {
        fmt.Println((id+k) % len(nums))
        nums[(id+k) % len(nums)] = val
    }
    fmt.Println(nums)
}
