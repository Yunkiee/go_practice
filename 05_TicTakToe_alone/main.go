package main

import "strings"

func main() {
	board1 := []string{"O.X", ".O.", "..X"}
	board2 := []string{"OOO", "...", "XXX"}
	board3 := []string{"...", ".X.", "..."}
	board4 := []string{"...", "...", "..."}
	print(solution(board1))
	print(solution(board2))
	print(solution(board3))
	print(solution(board4))
}

func solution(board []string) int {
	countO := 0
	countX := 0
	for _, column := range board {
		countO += strings.Count(column, "O")
		countX += strings.Count(column, "X")
	}

	if countO < countX {
		// 선공이 O 이므로 나올 수 없는 경우
		return 0
	} else if countO == countX {
		// 개수가 같을 경우
		// O가 게임을 이긴 경우가 있으면 안됨
		var firstRow, secondRow, thirdRow string
		for _, column := range board {
			firstRow += string(column[0])
			secondRow += string(column[1])
			thirdRow += string(column[2])
			if column == "OOO" {
				return 0
			}
		}
		if firstRow == "OOO" || secondRow == "OOO" || thirdRow == "OOO" {
			return 0
		}
		if string(board[0][0])+string(board[1][1])+string(board[2][2]) == "OOO" {
			return 0
		}
		if string(board[0][2])+string(board[1][1])+string(board[2][0]) == "OOO" {
			return 0
		}

	} else {
		if countO > countX+1 {
			// O의 갯수가 X의 개수 +1 보다 크면 안됨
			return 0
		} else {
			// countO = countX + 1 일때
			// X가 게임을 이긴 경우가 있으면 안됨
			var firstRow, secondRow, thirdRow string
			for _, column := range board {
				firstRow += string(column[0])
				secondRow += string(column[1])
				thirdRow += string(column[2])
				if column == "XXX" {
					return 0
				}
			}
			if firstRow == "XXX" || secondRow == "XXX" || thirdRow == "XXX" {
				return 0
			}
			if string(board[0][0])+string(board[1][1])+string(board[2][2]) == "XXX" {
				return 0
			}
			if string(board[0][2])+string(board[1][1])+string(board[2][0]) == "XXX" {
				return 0
			}
		}
	}
	return 1
}
