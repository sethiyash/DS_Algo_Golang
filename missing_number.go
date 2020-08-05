// Find the missing number in sorted array having elements from 1 to n in log(n) time complexity
// Example - input = [1,2,3,4,5,6,8] output = 7
// using divide and conqur strategy (Binary search)

package main

import (
	"fmt"
)

func missingNumber(input []int, n int) int {
	start, end := 0, n-1
	var mid int
	for ; end - start > 1;{
		mid = (start + end)/2
		if input[start] - start != input[mid] - mid {
			end = mid
		}
		if input[end]-end != input[mid] - mid {
			start = mid
		}
	}
	return input[mid]+1
}

func main() {
	input := []int{1,2,3,4,5,6,8}
	fmt.Println(missingNumber(input,len(input)))
}
