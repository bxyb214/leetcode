package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) == 0 || len(nums)*len(nums[0]) != r*c {
		return nums
	}
	mat := make([][]int, r)
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]int, c)
	}
	idx := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			mat[idx/c][idx%c] = nums[i][j]    //important
			idx++
		}
	}
	return mat
}

func main()  {
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 2, 4))

}