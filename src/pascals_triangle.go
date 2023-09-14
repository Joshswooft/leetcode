package main

// generates a pascal triangle for the given amount of rows e.g. generate(5) = [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
func generate(numRows int) [][]int {

	elems := [][]int{}

	// loop through rows
	for i := 0; i < numRows; i++ {

		rowElems := []int{}

		// loop through elements in that row
		for j := 0; j <= i; j++ {

			// 1st or last element in the row is 1
			if j == 0 || j == i {
				rowElems = append(rowElems, 1)
			} else {
				previousRow := elems[len(elems)-1]
				nextElem := 1
				if len(previousRow) > j+1 {
					nextElem = previousRow[j]
				}
				el := previousRow[j-1] + nextElem
				rowElems = append(rowElems, el)
			}
		}

		elems = append(elems, rowElems)

	}
	return elems
}
