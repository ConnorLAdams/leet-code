package main

import (
    "fmt"
)

func main() {
    l := []int{7,1,5,6,3,4}
    // l := []int{2,4,1}
    fmt.Println(maxProfit(l))
}

func maxProfit(prices []int) int {
    fmt.Println(prices)
    minP := prices[0]
    maxP := prices[0]
    maxD := maxP - minP
    for _, p := range prices[1:] {
        if p < minP {
            minP = p
            maxP = p
        } else if p > maxP {
            maxP = p
        }
        if maxP - minP > maxD {
            maxD = maxP - minP
        }
    }
    return max(maxD, maxP-minP)
}
