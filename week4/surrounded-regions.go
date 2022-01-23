func solve(board [][]byte) {

	row, col := len(board), len(board[0])

	//遍历每一行，针对每行首列以及最后一列元素进行深度优先搜索
	for x := 0; x < row; x++ {
		dfs(x, 0, board)
		dfs(x, col-1, board)
	}
	//遍历第二列到倒数第二列，针对第一行和最后一行数据数据进行深度优先搜索
	for y := 0; y < col; y++ {
		dfs(0, y, board)
		dfs(row-1, y, board)
	}
	//遍历每一个元素，如果A就更新其值为O，如果是O,则更新值为X
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if board[x][y] == 'A' {
				board[x][y] = 'O'
			} else if board[x][y] == 'O' {
				board[x][y] = 'X'
			}
		}
	}
}
func dfs(x, y int, board [][]byte) {
	//越界返回
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	//
	board[x][y] = 'A'
	//分别查找上下左右
	dfs(x-1, y, board)
	dfs(x+1, y, board)
	dfs(x, y-1, board)
	dfs(x, y+1, board)
}
    