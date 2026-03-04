package main

import "fmt"

func main() {
	//fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	//fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}


func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	strs = equalizarPeloMenor(strs)

	prefix := strs[0]


	j:= 0

	var k int

	temp:= ""

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
				return word[:k]
			}
			k++
			
		}

		if temp != prefix[:k] {
			temp = prefix[:k]
		}
		
	}

	return prefix[:k]

}

func equalizarPeloMenor(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}

	menor := len(strs[0])
	for _, s := range strs {
		if len(s) < menor {
			menor = len(s)
		}
	}

	resultado := make([]string, len(strs))
	for i, s := range strs {
		resultado[i] = s[:menor]
	}

	return resultado
}
