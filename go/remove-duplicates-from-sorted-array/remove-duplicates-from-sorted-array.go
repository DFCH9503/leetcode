package main

import(
	"fmt"
)

func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i <len(nums); i++{
		if nums[i] != nums[i-1]{
			nums [index] = nums[i]
			index++
		}
		fmt.Println(nums[index], index)
	}


	return index
}

func main(){
	array1 := []int {1 , 1 , 2, 2, 3, 4, 4}
	fmt.Println(array1, len(array1), cap(array1))
	fmt.Println("Response is",removeDuplicates(array1))
}