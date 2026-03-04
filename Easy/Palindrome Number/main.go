package main

import "fmt"

func main() {
	//Example
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
    if x < 0{
        return false
    }
    
    before := x
    var reverse int

    for x > 0 {
        digit := x % 10
        reverse = reverse*10 + digit
        x = x / 10
    }

    if before == reverse{
        return true
    }

    return false
}