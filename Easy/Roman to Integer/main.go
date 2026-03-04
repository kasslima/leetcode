package main

import "fmt"

func main() {
	//Example
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
    roman := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    var result int
    
    for i:= 0; i< len(s); i++{
        j := string(s[i])
        
            result += roman[j]
        

        if i > 0 {
            before := string(s[i-1])
            if roman[j] > roman[before]{
                result -= (roman[before]*2)
            }
        }
    }

    return result
}