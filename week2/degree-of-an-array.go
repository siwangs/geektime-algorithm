func findShortestSubArray(nums []int) int {
	// startArrayMap 记录每个数字的第一个下标位置
	// maxDegreeMap 每个数字的度，如果发现最大的度，用MaxDegree进行替换
	var startArrayMap = make(map[int]int)
	var maxDegreeMap = make(map[int]int)
	var maxDegree, result int

	for k, v := range nums {
		maxDegreeMap[v] += 1
		if _, ok := startArrayMap[v]; !ok {
			startArrayMap[v] = k
		}
		//更新最大度，并计算最相同度的子数组
		if maxDegreeMap[v] > maxDegree {
			maxDegree = maxDegreeMap[v]
			result = k - startArrayMap[v] + 1
		} else if maxDegreeMap[v] == maxDegree {
			result = min(k-startArrayMap[v]+1, result)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}