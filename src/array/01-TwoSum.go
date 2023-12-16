
func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	ans := make([]int, 2, 2)
	for i, num := range nums {
		if myMap[target-num] != 0 {
			ans[0] = myMap[target-num] - 1
			ans[1] = i
			break
		} else {
			myMap[num] = i + 1 // storing intentionally i+1 because, we are storing index as value in map
		}
	}
	return ans
}