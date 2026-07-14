package array2

import (
	"slices"
)

/*
You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.

Example 1:

Input: nums = [1,2,2,4]
Output: [2,3]

Example 2:

Input: nums = [1,1]
Output: [1,2]

Constraints:

    2 <= nums.length <= 104
    1 <= nums[i] <= 104


*/

func FindErrNums(nums []int) []int {
	sorted := append([]int(nil), nums...)
	slices.Sort(sorted)

	dup := 0
	for i := 1; i < len(nums); i++ {
		if sorted[i] == sorted[i-1] {
			dup = sorted[i]
			break
		}

	}

	for i := 1; i < len(sorted); i++ {
		if !slices.Contains(sorted, i) {
			return []int{dup, i}
		}
	}
	max := len(sorted) - 1
	return []int{dup, sorted[max] + 1}

}

// Time complexity is O(n2) -> very expensive

/* Q2. How Many Numbers Are Smaller Than the Current Number
Easy
Topics
premium lock iconCompanies
Hint

Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it. That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

Return the answer in an array.



Example 1:

Input: nums = [8,1,2,2,3]
Output: [4,0,1,1,3]
Explanation:
For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
For nums[1]=1 does not exist any smaller number than it.
For nums[2]=2 there exist one smaller number than it (1).
For nums[3]=2 there exist one smaller number than it (1).
For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).

Example 2:

Input: nums = [6,5,4,8]
Output: [2,1,0,3]

Example 3:

Input: nums = [7,7,7,7]
Output: [0,0,0,0] */

// time complexity O(n log n) space complexity O(n)
// sort and use map
func SmallerNumbersThanCurrent(nums []int) []int {
	results := make([]int, len(nums))
	sorted := append([]int(nil), nums...)
	slices.Sort(sorted)

	first := make(map[int]int) //first occurence

	for i, v := range sorted {
		if _, ok := first[v]; !ok {
			first[v] = i

		}

	}

	//Build results using original order
	for i, v := range nums {
		results[i] = first[v]
	}

	return results

}

// Time complexity: O(n2) -
// space compexity: O(1) - constant space
// Ideal for smaller inputs
func BSmallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))

	for i := range nums {
		count := 0
		for j := range nums {
			if nums[j] < nums[i] {
				count++
			}
		}
		result[i] = count
	}

	return result
}

// Brute force

func FindDisappearedNumbers(nums []int) []int {
	sorted := append([]int(nil), nums...)
	slices.Sort(sorted)
	results := []int{}

	for num := 1; num <= len(sorted); num++ {
		if !slices.Contains(sorted, num) {
			results = append(results, num)

		}
	}
	return results
}
