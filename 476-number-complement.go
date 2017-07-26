package main

import "fmt"

func findComplement(num int) int {
	ret := 0
	t := 1
	for num > 0 {
		if (num & 1) == 0 {
			ret += t
		}
		num >>= 1
		t <<= 1
	}
	return ret
}

func main()  {
	fmt.Println(findComplement(5))
}