package main

import (
	"fmt"
	"leetcode/quests/array1"
	"leetcode/quests/stack"
)

func main() {
	nums := []int{1, 1}
	fmt.Println(array1.GetConcatenation(nums))

	// Q2 shuffle array
	//forw := []int{1, 2, 2, 4, 5}
	//back := []int{1, 2, 4, 4, 5}
	//dig := []int{4, 3, 2, 7, 8, 2, 3, 1, 12, 16}
	// digb := []int{2, 2}
	//larg := []int{4, 2, 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 3}
	// mid := []int{2, 1, 4, 4, 5}
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	fmt.Println(stack.EvalRPN(tokens))

}

/* test cases
1. missing 1 -> done
2. two digits -> done
3. middle after dup no missing [2,1,4,4,5] -> 4, 3
4. before dup [1,2,3,4,6,6,7] -> [6,5]
5. dup at the end [1,2,3,4,5,5] -> after
6. largest missing [6,2,3,4,5,6,7,8,1]
*/
