package main

import "fmt"

func main() {
	//Example
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"aac","acab","aa","abba","aa"}))
	fmt.Println(longestCommonPrefix([]string{"abca","aba","aaab"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}


func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	strs = normalizeToShortest(strs)

	prefix := strs[0]
	j:= 0

	var k int
	temp:= ""
	new:= ""

	for i := range strs{
		k = 0

		for j = 0; j <= len(strs[i]); j++{
			word := strs[i]

			if k == len(word) {
				prefix = word[:k]
				break
			}

			if word[:j+1] != prefix[:j+1] {
				if j == 0 {
					return ""
				}
				break
			}

			k++

		}

		new = prefix[:k]

		if temp == "" || len(new) < len(temp) {
			temp = new
		}
	}

	return temp

}

func normalizeToShortest(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}

	shortest := len(strs[0])
	for _, s := range strs {
		if len(s) < shortest {
			shortest = len(s)
		}
	}

	result := make([]string, len(strs))
	for i, s := range strs {
		result[i] = s[:shortest]
	}

	return result
}
