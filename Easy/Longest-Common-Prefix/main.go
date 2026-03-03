package main

import "fmt"

func main() {
    fmt.Println(longestCommonPrefix([]string{"flower","flower","flower", "flower"}))
}


func longestCommonPrefix(strs []string) string {
    
    if len(strs) == 1{
        return strs[0]
    }

    j := 1
    
    prefix := strs[0][:j]

    for i := range strs {

        if i >= len(strs) || j > len(strs[i]){
            return prefix
        }

        if prefix != strs[1][:j] && j == 1 {
            return ""
        }

        if prefix != strs[i][:j] {
            return prefix[:j-1]
        }

        j++

        prefix = strs[0][:j]

        if len(strs[i]) >= j{
            if len(strs[i+1]) < j{
                return prefix[:j-1]
            } 

            if prefix != strs[i+1][:j]{
                return prefix[:j-1]
            }

        }
        
        
    }

    return "nil"

}