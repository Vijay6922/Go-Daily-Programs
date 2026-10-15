package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}

	rowSum := 0
	colSum := 0

	for i := 0; i < len(matrix); i++ {
		rowSum=0
		for j := 0; j <len(matrix[0]); j++ {
			rowSum += matrix[i][j]
		}
		fmt.Println("sum of rows :",rowSum)
	}

	for i := 0; i < len(matrix[0]); i++ {
		colSum=0
		for j := 0; j <len(matrix); j++ {
			colSum += matrix[j][i]
		}
		fmt.Println("sum of cols :",colSum)
	}

}