package main

import "fmt"

func distributeCandies(candies []int) int {
	m := make(map[int]interface{})
	v := 0
	for i := 0; v < len(candies)/2 && i < len(candies); i++ {
		if _, ok := m[candies[i]]; ok == false {   // mean what?
			v++
			m[candies[i]] = nil
		}
	}
	return v
}

func main()  {
	fmt.Println(distributeCandies([]int{1, 1, 1, 1, 2, 3}))
}
