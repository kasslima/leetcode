package main

import "fmt"

func main() {
	//Example
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) + len(nums2) == 0{
        return 0.0
    }

    var j, k int
    merged := []int{}

    for j < len(nums1) && k < len(nums2) {
        if nums1[j] > nums2[k] {
            merged = append(merged, nums2[k])
            k++

        } else if nums1[j] < nums2[k] {
            merged = append(merged, nums1[j])
            j++

        } else {
            merged = append(merged, nums1[j])
            merged = append(merged, nums2[k])
            j++
            k++
        }
    }

    for j < len(nums1) {
        merged = append(merged, nums1[j])
        j++
    }

    for k < len(nums2) {
        merged = append(merged, nums2[k])
        k++
    }

    size:= len(merged)

    if size == 0{
        return 0.0
    }

    if size == 1{
        return float64(merged[0])
    }

    if size % 2 == 0{
        size:= size/2
        return float64(merged[size-1]+merged[size]) / 2.0
    }

    return float64(merged[((size + 1)/2.0)-1])
    
}