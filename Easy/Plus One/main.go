package main

import "fmt"

func main() {
	//Example
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{4,3,2,1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9,9}))
}

func plusOne(digits []int) []int {
    if len(digits) == 0{
        return digits
    }

    for i:= len(digits)-1; i >= 0; i-- {

        if digits[i] >= 9{
            if i == 0 && digits[i] >= 9{
                if digits[i] == 9 && len(digits) != 1{
                    return digits
                }
                digits[i] = 0
                digits = append([]int{1}, digits...)
                return digits
            }

            if digits[i] == 9 && len(digits)-1 != i{
                return digits
            }

            digits[i] = 0
            digits[i-1]++

        }else if i == len(digits)-1{
            digits[len(digits)-1]++
            return digits
        }

    }

    return digits
}