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

/* Q2. Shuffle the Array

Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

Return the array in the form [x1,y1,x2,y2,...,xn,yn].

Example 1:

Input: nums = [2,5,1,3,4,7], n = 3
Output: [2,3,5,4,1,7]
Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].

Example 2:

Input: nums = [1,2,3,4,4,3,2,1], n = 4
Output: [1,4,2,3,3,2,4,1] */

func Shuffle(nums []int, n int) []int {
	shuffled := make([]int, 0, len(nums))

	for i := range n {
		j := i + n

		shuffled = append(shuffled, nums[i], nums[j])
	}
	return shuffled
}

// space complexity is O(n) -> space grows as input grows
// time comlexity is O(n log n)

/* Q3. Max Consecutive Ones
Solved
Easy

Given a binary array nums, return the maximum number of consecutive 1's in the array.



Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2



Constraints:

    1 <= nums.length <= 105
    nums[i] is either 0 or 1.
*/

func findMaxConsecutiveOnes(nums []int) int {
	var count, reset, max int
	for _, num := range nums {
		if num == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = reset
		}
	}
	return max
}
