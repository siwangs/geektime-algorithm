func findOrder(numCourses int, prerequisites [][]int) []int {
	// 画图并记录入度
	inDeg, toCourse := make([]int, numCourses), make([][]int, numCourses)
	for _, v := range prerequisites {
		inDeg[v[0]]++
		toCourse[v[1]] = append(toCourse[v[1]], v[0])
	}
	var studyQueue []int
	// 将入度为0加入可学列队
	for k, v := range inDeg {
		if v == 0 {
			studyQueue = append(studyQueue, k)
		}
	}
	// 开始学习
	var result []int
	for len(studyQueue) != 0 {
		head := studyQueue[0]
		studyQueue = studyQueue[1:]
		result = append(result, head)
		// 把下门课的入度-1， 如果入度为0，则可以加入列队进行学习
		for _, v := range toCourse[head] {
			inDeg[v]--
			if inDeg[v] == 0 {
				studyQueue = append(studyQueue, v)
			}
		}
	}
	if len(result) == numCourses {
		return result
	}
	return []int{}
}