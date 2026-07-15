package stack

import (
	"slices"
	"strconv"
)

/*
https://leetcode.com/problems/build-an-array-with-stack-operations/description/?envType=problem-list-v2&envId=dsa-linear-shoal-stack
*/

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

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]

}
func operate() {

}
// https://leetcode.com/problems/evaluate-reverse-polish-notation/submissions/2069099778/?envType=problem-list-v2&envId=dsa-linear-shoal-stack
func isOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

func EvalRPN(tokens []string) int {
	items := []int{} // stack LIFO
	for _, tok := range tokens {
		if !isOperator(tok) {
			tokInt, _ := strconv.Atoi(tok)
			items = append(items, tokInt)
		} else {
			switch tok {
			case "/":
				pop1 := items[len(items)-2]
				pop2 := items[len(items)-1]
				items = items[:len(items)-1]
				items = items[:len(items)-1]
				op := pop1 / pop2
				items = append(items, op)
			case "-":
				pop1 := items[len(items)-2]
				pop2 := items[len(items)-1]
				items = items[:len(items)-1]
				items = items[:len(items)-1]
				op := pop1 - pop2
				items = append(items, op)
			case "+":
				pop1 := items[len(items)-2]
				pop2 := items[len(items)-1]
				items = items[:len(items)-1]
				items = items[:len(items)-1]
				op := pop1 + pop2
				items = append(items, op)
			default:
				pop1 := items[len(items)-2]
				pop2 := items[len(items)-1]
				items = items[:len(items)-1]
				items = items[:len(items)-1]
				op := pop1 * pop2
				items = append(items, op)
			}
		}
	}
	return int(items[0])

}
