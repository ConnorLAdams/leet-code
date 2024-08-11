package main

import "fmt"

func main() {
    // l := []int{1,1,1,2,2,3}
    // l := []int{0,0,1,1,1,1,2,3,3}
    // l := []int{1}
    // l := []int{1,1}
    // l := []int{1,2}
    // l := []int{1,2,2}
    l := []int{1,1,1,1}
    fmt.Println(l)
    fmt.Println(removeDuplicates(l))
}

func removeDuplicates(nums []int) int {
    i := 0
    ni_occr := 1
    j := 2
    n := len(nums)

    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    } else {
        for j <= n {
            if nums[i] == nums[i+1] {
                if j != n {
                    if ni_occr >= 2 {
                        nums[i+1], nums[j] = nums[j], nums[i+1]
                        j++
                    } else {
                        i++
                        ni_occr++
                    }
                } else {
                    if ni_occr < 2 {
                        i++
                    } else {
                        i++
                        j++
                    }
                }
            } else if nums[i] > nums[i+1] {
                if j!=n {
                    nums[i+1], nums[j] = nums[j], nums[i+1]
                    j++
                } else {
                    i++
                    j++
                }
            } else {
                i++
                ni_occr = 1
            }
    
            if j-i < 2 {
                j++
            }
            if (i == n-1) && (nums[i-1] <= nums[i]) && (ni_occr < 2) {
                i++
            }
            fmt.Println("nums:", nums)
            fmt.Println("i:", i)
            fmt.Println("j:", j)
            fmt.Println("ni_occr:", ni_occr)
            fmt.Println("-------------------------------------------")
        }
    }
    fmt.Println(nums)
    return i
}
