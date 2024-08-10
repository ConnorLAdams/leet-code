package main

import "fmt"

func main() {
    l := []int{1,2,3}
    fmt.Println(l)
    fmt.Println(removeDuplicates(l))
}

func removeDuplicates(nums []int) int {
    i := 0
    j := 2
    n := len(nums)

    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else if n == 2 {
        if nums[0] == nums[1] {
            return 1
        } else {
            return 2
        }
    } else {
        for j <= n {
            if nums[i] >= nums[i+1]{
                if j != n {
                    nums[i+1], nums[j] = nums[j], nums[i+1]
                    j++
                } else {
                    i++
                    j++
                }
            } else {
                if j-i == 2 {
                    i++
                    j++
                } else {
                    i++
                }
            }
        }
        if (i == n-1) && (nums[i-1] < nums[i]) {
            i++
        }
    }
    fmt.Println(nums)
    return i
}
