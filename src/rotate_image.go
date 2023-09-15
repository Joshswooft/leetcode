package main

// func reverseMatrixX(matrix [][]int) {
// 	numRows := len(matrix)
// 	if numRows == 0 {
// 		return // Return if the matrix is empty.
// 	}
// 	numCols := len(matrix[0])

// 	for i := 0; i < numRows/2; i++ {
// 		for j := 0; j < numCols; j++ {
// 			matrix[i][j], matrix[numRows-1-i][j] = matrix[numRows-1-i][j], matrix[i][j]
// 		}
// 	}
// }

func transpose(matrix [][]int) {
	numRows := len(matrix)
	if numRows == 0 {
		return // Return if the matrix is empty.
	}
	numCols := len(matrix[0])

	for i := 0; i < numRows; i++ {
		for j := i + 1; j < numCols; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverseMatrixY(matrix [][]int) {
	numRows := len(matrix)
	if numRows == 0 {
		return // Return if the matrix is empty.
	}

	for i := 0; i < numRows; i++ {
		numCols := len(matrix[i])
		for j := 0; j < numCols/2; j++ {
			matrix[i][j], matrix[i][numCols-1-j] = matrix[i][numCols-1-j], matrix[i][j]
		}
	}
}

func rotate(matrix [][]int) {
	transpose(matrix)
	// reverseMatrixX(matrix)
	reverseMatrixY(matrix)
}
