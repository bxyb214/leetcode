package main

import "fmt"

func findWords(words []string) []string {
	str := []string{"QWERTYUIOPqwertyuiop", "ASDFGHJKLasdfghjkl", "ZXCVBNMzxcvbnm"}
	dic := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			dic[str[i][j]] = i
		}
	}
	ret := make([]string, 0, len(words))
	for _, word := range words {
		flag := true
		line := dic[word[0]]
		for i := 0; flag && i < len(word); i++ {
			if line != dic[word[i]] {
				flag = false
			}
		}
		if flag {
			ret = append(ret, word)
		}
	}
	return ret
}

func main()  {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))}