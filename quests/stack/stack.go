package stack

import (
	"slices"
	"strconv"
	"strings"
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

func ExclusiveTime(n int, logs []string) []int {
	results := make([]int, n)
	stack := []int{}
	prevTime := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")

		id, _ := strconv.Atoi(parts[0])
		event := parts[1] // start/stop
		time, _ := strconv.Atoi(parts[2])
		if event == "start" {
			// the function currently on top has been running
			// from prev time until just this time
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				results[top] += time - prevTime
			}
			stack = append(stack, id)
			prevTime = time

		} else { // "end"
			// The current function finishes at END of time
			// so we include this timestamp (+1)
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			results[top] += time - prevTime + 1

			//Next execution after this timestamp
			prevTime = time + 1
		}

	}
	return results
}

// ["0:start:0","1:start:2","1:end:5","0:end:6"]
