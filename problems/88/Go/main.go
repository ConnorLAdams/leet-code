package main

import (
	"fmt"
	"sort"
	"reflect"
)

func main() {
	l1 := []int{1,2,3,0,0,0}
	ln := 3
	l2 := []int{2,5,6}
	lm := 3
	t := []int{1,32,3,6,2}
	sort.Ints(t)
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(t)
	fmt.Println(l1, l2, ln, lm)
	merge(l1, l2, ln, lm)
}

func merge(nums1 []int, nums2 []int, n int, m int) {
	for i := n; i<n+m; i++ {
		nums1[i] = nums2[i-n]
	}
	sort.Ints(nums1)
	id := 0
	for id2, ele := range nums2 {
		fmt.Println("Nums2 element: ", ele)
		fmt.Println("Nums2[ID2]: ", nums2[id2])
		fmt.Println("ID: ", id)
		fmt.Println("Nums1[ID]: ", nums1[id])
		fmt.Println("N: ", n)
		for id <= n {
			if ele < nums1[id] {
				fmt.Println(id, nums1[id])
				for j := id; j <= n; j++ {
					if n < len(nums1) {
						nums1[j], nums2[id2] = nums2[id2], nums1[j]
					}
				}
				fmt.Println("Nums1: ", nums1)
				fmt.Println("Nums2: ", nums2)
				if n < len(nums1) {
					n += 1
				}
				id += 1
				break
			} else if nums1[id] == 0 {
				nums1[id], nums2[id2] = nums2[id2], nums1[id]
				if n < len(nums1) {
					n += 1
				}
				id += 1
				break
			} else {
				id += 1
			}
		}
	}
	fmt.Println(nums1, m)
}
