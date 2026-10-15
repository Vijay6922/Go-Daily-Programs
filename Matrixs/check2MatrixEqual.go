package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	matrix2 := [][]int{{1, 2, 2}, {4, 5, 6}}

	equal := true
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix2[i][j] {
				equal = false
				break
			}
		}
		if !equal {
			break
		}
	}

	if equal == true {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}