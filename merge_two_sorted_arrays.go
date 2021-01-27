package main

import (
	"fmt"
)

func main()  {
	var n1, n2 int
	var err error
	_, err = fmt.Scanf("%d", &n1)

	arr1 := make([]int, n1)
	for i:=0; i<n1; i++ {
		_, err = fmt.Scanf("%d", &arr1[i])
	}
	_, err = fmt.Scanf("%d", &n2)

	arr2 := make([]int, n2)
	for i:=0; i<n2; i++ {
		_, err = fmt.Scanf("%d", &arr2[i])
	}

	if err != nil { panic("") }

	res := make([]int, n1+n2)

	var i,j,k int

	for ; i < n1 && j < n2; k++{
		if arr1[i] < arr2[j] {
			res[k] = arr1[i]
			i++
		} else {
			res[k] = arr2[j]
			j++
		}
	}

	for ;i<n1; i++{
		res[k] = arr1[i]
		k++
	}

	for ;j<n2; j++{
		res[k] = arr2[j]
		k++
	}

	fmt.Println("Result - ", res)

}
