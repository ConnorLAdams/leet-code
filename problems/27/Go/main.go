package main

import "fmt"

func main(){
    l := []int{3,2,2,3}
    v := 3
    removeElement(l, v)
}

func removeElement(nums []int, val int) int {
    c := 0
    k := len(nums)
    n := 0
    m := len(nums) - 1
    for n<=m {
        c++
        if nums[m] == val {
            k--
            m--
        } else {
            if nums[n] == val {
                nums[m], nums[n] = nums[n], nums[m]
                k--
                n++
                m--
            } else {
                n++
            }
        }
    }
    fmt.Println(nums)
    fmt.Println(k)
    fmt.Println(c)
    return k
}
