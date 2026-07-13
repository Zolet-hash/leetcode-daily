package array1

/*Q1. Concatenation of Array
Solved
Easy
Topics
premium lock iconCompanies
Hint

Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

Specifically, ans is the concatenation of two nums arrays.

Return the array ans. */

func GetConcatenation(nums []int) []int {
	// nums = [1,2,1] becomes nums2 = [1,2,1,1,2,1]
	numLen := len(nums) * 2

	num2 := make([]int, 0, numLen)

	for _, r := range nums {
		num2 = append(num2, int(r))
	}

	for _, r := range nums {
		num2 = append(num2, int(r))
	}
	return num2
}
