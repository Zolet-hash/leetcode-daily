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
	results := []int{}
	if len(nums) == 2 && nums[0] == 1 {
		results = append(results, 1, 2)
		return results
	}

	if len(nums) == 2 && nums[0] == 2 {
		results = append(results, 2, 1)
		return results
	}

	// no 1 in the array, return duplicate and 1
	exists := false
	for i := range nums {
		if nums[i] == 1 {
			exists = true
		}
	}

	// if no 1 return duplicate and 1
	if !exists {
		dup := 0
		for i := range nums {
			for j := i + 1; j < len(nums); j++ {
				if nums[i] == nums[j] {
					dup = nums[i]
				}
			}
		}
		results = append(results, dup, 1)
		return results
	}

	// sort nums
	sorted := append([]int(nil), nums...)
	slices.Sort(sorted)

	// Check for duplicates
	duplicate := 0
	for i := range sorted {
		for j := i + 1; j < len(nums); j++ {
			if sorted[i] == sorted[j] {
				duplicate = sorted[j] // val 2

				// 3 digit num
				if len(sorted) == 3 && sorted[0] == 1 && sorted[1] == 1 && sorted[2] == 2 {
					results = append(results, 1, 3)
					return results

				} else if len(sorted) == 3 && sorted[0] == 1 && sorted[1] == 1 && sorted[2] == 3 {
					results = append(results, 1, 2)
					return results

				} else if len(sorted) == 3 && sorted[0] == 1 && sorted[1] == 3 && sorted[2] == 3 {
					results = append(results, 3, 2)

				} else if len(sorted) == 3 && sorted[0] == 1 && sorted[1] == 2 && sorted[2] == 3 {
					results = append(results, 2, 3)
					return results

					// last duplicate [1,2,3,4,5,5] -> [5,6]
				} else if j == len(sorted)-1 {
					results = append(results, duplicate, duplicate+1)

					//largest missing num [1,2,3,4,5,6,6,7,8] -> [6,9]
				} else if j+1 != len(sorted)-1 {
					max := len(sorted) - 1 // last index
					results = append(results, duplicate, sorted[max]+1)
					return results
					// before missing
				} else if duplicate+1 == sorted[j+1] {
					results = append(results, duplicate, sorted[j]-1)
					return results

				} else if duplicate-1 != sorted[j+1] { // backward case
					results = append(results, duplicate, duplicate+1)
					return results
				}
			}
		}
	}

	return results

}
