//two sum
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, found := numMap[complement]; found {
            return []int{j, i}
        }
        numMap[num] = i
    }
    return nil // No solution found
}