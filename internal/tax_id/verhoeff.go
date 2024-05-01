package tax_id

import "fmt"

func verhoeffCheckSum(fiscalID string, serialNumberNormalized string, dateUTF8 string) string {
	verhoeffTableD := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
		{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
		{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
		{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
		{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
		{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
		{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
		{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	}
	verhoeffTableP := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
		{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
		{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
		{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
		{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
		{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
		{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
	}
	verhoeffTableInv := []int{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}

	utf8 := alphabetToOrd(fiscalID) + dateUTF8 + serialNumberNormalized
	c := 0
	for i := len(utf8) - 1; i >= 0; i-- {
		item := utf8[i]
		digit := int(item - '0')
		c = verhoeffTableD[c][verhoeffTableP[(len(utf8)-i)%8][digit]]
	}
	return string(rune(verhoeffTableInv[c] + '0'))
}

// convert alphabet to number
func alphabetToOrd(item string) string {
	var result []rune
	for _, char := range item {
		if !('0' <= char && char <= '9') {
			result = append(result, []rune(fmt.Sprintf("%d", int(char)))...)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
