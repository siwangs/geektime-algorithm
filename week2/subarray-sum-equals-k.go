func subarraySum(nums []int, k int) int {
	//preSum用来记录前缀和
	result, preSum := 0, 0
	preSumMap := make(map[int]int)
	preSumMap[0] = 1
	for _, value := range nums {
		//默认记录当前前缀和
		preSum += value
		//如果如果前缀和存在，记录之前的结果
		if _, ok := preSumMap[preSum-k]; ok {
			result += preSumMap[preSum-k]
		}
		preSumMap[preSum] += 1
	}
	return result
}
    
    