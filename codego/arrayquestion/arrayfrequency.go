package main

import (
	"fmt"
	"sort"
)

func main() {
	arr3 := []int{6, 2, 3, 3, 4, 4, 2, 2, 1}

	// Step 1: Count frequency
	freq := make(map[int]int)
	for _, v := range arr3 {
		freq[v]++
	}

	// Step 2: Convert map to slice of pairs
	type pair struct {
		val  int
		freq int
	}
	var pairs []pair
	for k, v := range freq {
		pairs = append(pairs, pair{k, v})
	}

	// Step 3: Sort by frequency, then by value
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].freq == pairs[j].freq {
			return pairs[i].val < pairs[j].val
		}
		return pairs[i].freq < pairs[j].freq
	})

	// Step 4: Print values
	for _, p := range pairs {
		fmt.Print(p.val, " ")
	}
}
