package main

import "fmt"

func maxHappyGroups(batchSize int, groups []int) int {
	length := len(groups)
	bitMaskLen := 1 << length
	dp := make([]int, bitMaskLen)
	for i := 0; i < bitMaskLen; i++ {
		var mod = 0
		for n := 0; n < length; n++ {
			if (i & (1 << n)) != 0 {
				mod = (mod + groups[n]) % batchSize
			}
		}
		var happyCount = 0
		if mod%batchSize == 0 {
			happyCount = 1
		}
		for n := 0; n < length; n++ {
			if (i & (1 << n)) == 0 {
				mask := i | (1 << n)
				dp[mask] = max(dp[mask], dp[i]+happyCount)
			}
		}
	}
	return dp[bitMaskLen-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	type arg struct {
		BatchSize int
		Groups    []int
	}
	args := []arg{
		{
			BatchSize: 3,
			Groups:    []int{1, 2, 3, 4, 5, 6},
		},
		{
			BatchSize: 4,
			Groups:    []int{1, 3, 2, 5, 2, 2, 1, 6},
		},
	}
	for _, a := range args {
		fmt.Println(maxHappyGroups(a.BatchSize, a.Groups))
	}
}
