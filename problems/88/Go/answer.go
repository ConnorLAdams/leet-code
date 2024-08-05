package main

func main() {
	l1 := []int{1,2,3,0,0,0}
	lm := 3
	l2 := []int{2,5,6}
	ln := 3
	merge(l1, lm, l2, ln)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
