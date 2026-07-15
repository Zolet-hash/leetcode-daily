package stack

import "slices"

/*
https://leetcode.com/problems/build-an-array-with-stack-operations/description/?envType=problem-list-v2&envId=dsa-linear-shoal-stack
*/

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]

}

func BuildArray(target []int, n int) []string {
	result := []string{}

	last := len(target) - 1 // last index
	for num := 1; num <= target[last]; num++ {
		if !slices.Contains(target, num) {
			result = append(result, "Push", "Pop")

		} else {
			result = append(result, "Push")
		}
	}
	return result

}
