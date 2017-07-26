package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))

}

func hammingDistance(x int, y int) int {
	res := x ^ y
	cnt := 0
	for res > 0 {
		res &= res - 1
		cnt++
	}
	return cnt
}