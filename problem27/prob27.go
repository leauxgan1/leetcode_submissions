package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	count := length
	for i := 0; i < count; i++ {
		if nums[i] == val {
			count--

			for foundSpot := false; !foundSpot; {
				if i == count || count < 0 {
					return count
				} else if nums[count] == val {
					count--
				} else {
					foundSpot = true
				}
			}
			nums[i], nums[count] = nums[count], nums[i]
		}
	}
	return count
}

func main() {
	myArr := [5]int{1, 3, 3, 2, 3}
	fmt.Println("Initial state of Array: ", myArr)
	val := 3
	numRemaining := removeElement(myArr[0:], val)
	fmt.Printf("Final state of Array after removing %d: %v\n", val, myArr[0:numRemaining])
	fmt.Println("Number of values that were not matched: ", numRemaining)

}
