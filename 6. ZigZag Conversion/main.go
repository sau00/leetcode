package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]rune, numRows)

	currentRow := 0
	down := false
	for _, sym := range s {
		rows[currentRow] = append(rows[currentRow], sym)
		if currentRow == 0 || currentRow == (numRows-1) {
			down = !down
		}

		if down {
			currentRow++
		} else {
			currentRow--
		}
	}

	res := ""
	for _, r := range rows {
		res += string(r)
	}

	return res
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
