package warwuck

import (
	"warwuck/champions"
	"warwuck/extracthangul"

	"github.com/agnivade/levenshtein"
)

func Similarity(_str string) []int {
	str := extracthangul.ExtractHangul(_str)
	arr := []int{}
	for _, r := range champions.ExtractedChampions {
		arr = append(arr, levenshtein.ComputeDistance(str, r))
	}
	return arr
}

func FindChampion(_str string) string {
	arr := Similarity(_str)
	min := 0
	for i, v := range arr {
		if v < arr[min] {
			min = i
		}
	}
	return champions.Champions[min]
}
