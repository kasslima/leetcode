package main

import "fmt"

func main() {
	//Example
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums {
        m[v] = i
    }
    
    for j := range nums {
        if c, ok := m[target - nums[j]]; ok && c != j{
            return []int{j, c}
        }
    }
    
    return nil
}
