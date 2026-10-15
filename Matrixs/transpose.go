package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	for i:=0;i<len(matrix);i++ {
		fmt.Println(matrix[i])
	}

	// 2. Allocate the outer rows for matrix2 (it needs 'numCols' rows)
	matrix2:=make([][]int, len(matrix[0]))

	// 3. Allocate the inner columns for each row (each row needs 'numRows' columns)
	for j := 0; j < len(matrix[0]); j++ {
		matrix2[j] = make([]int, len(matrix[0]))
	}

	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			matrix2[j][i]=matrix[i][j]
		}
	}

	fmt.Println()
	for i:=0;i<len(matrix2);i++ {
		fmt.Println(matrix2[i])
	}
}