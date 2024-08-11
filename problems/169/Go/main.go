package main

import "fmt"

func main() {
    // l := []int{2,2,1,1,1,2,2}
    l := []int{3,2,3}
    fmt.Println(majorityElement(l))
}

func majorityElement(nums []int) int {
    v := nums[0]
    cnt := 1
    r := len(nums)-1
    for i := 1; cnt <= r; i++ {
        if nums[i] == v {
            cnt++
        } else {
            if cnt == 1 {
                v = nums[i]
            } else {
                cnt--
            }
        }
        r--
    }

    return v
}
