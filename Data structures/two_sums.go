func twoSum(nums []int, target int) []int {
    checkNums := map[int]int{}
    for key, val := range nums {
        aux := target - val
        if j, found := checkNums[aux]; found {
            return []int{key, j}
        }
        checkNums[val] = key
    }
    return []int{}
}