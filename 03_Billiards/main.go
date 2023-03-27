package main

import (
	"fmt"
	"strings"
)

func main() {
	//var balls [][]int
	//var inner []int
	//inner = append(inner, 2)
	//inner = append(inner, 7)
	//balls = append(balls, inner)
	//solution(10,10,3,7,balls)
	id := "TN920_9911961058"
	//id2 := "TN920_9911961058_9911961058"
	idSplit := strings.SplitN(id, "_", 2)
	fmt.Println(idSplit[0])
	fmt.Println(idSplit[1])
}

func solution(m int, n int, startX int, startY int, balls [][]int) (result []int) {
	for _, target := range balls {
		var minValue int
		distance := 1000 * 1000
		distance2 := 1000 * 1000

		height := 0
		width := 0

		// 같은 x축
		if startX == target[0] {
			// 직선 쿠션 사용 가능
			if startY < target[1] {
				distance = startY + target[1]
			} else {
				distance = (n * 2) - startY - target[1]
			}

			// 대각 쿠션
			if startX <= m/2 {
				width = startX * 2
			} else {
				width = (m - startX) * 2
			}
			height = startY - target[1]
			distance2 = (height * height) + (width * width)

			if (distance * distance) < distance2 {
				minValue = distance * distance
			} else {
				minValue = distance2
			}

			result = append(result, minValue)

		} else if startY == target[1] { // 같은 y축
			// 직선 쿠션 사용 가능
			if startX < target[0] { // (3,7) -> (2,7)
				distance = startX + target[0]
			} else {
				distance = (m * 2) - startX - target[0]
			}

			// 대각 쿠션
			if startY <= n/2 {
				height = startY * 2
			} else {
				height = (n - startY) * 2
			}
			width = startX - target[0]
			distance2 = (height * height) + (width * width)

			if (distance * distance) < distance2 {
				minValue = distance * distance
			} else {
				minValue = distance2
			}

			result = append(result, minValue)

		} else { // 일반적 대각 쿠션
			result = append(result, findMinDistance(m, n, startX, startY, target[0], target[1]))
		}
	}

	return result
}

func findMinDistance(m int, n int, startX int, startY int, endX int, endY int) int {
	var min []int
	minValue := 1000 * 1000

	min = append(min, distanceTwoPoints(startX*(-1), startY, endX, endY))
	min = append(min, distanceTwoPoints(startX, startY*(-1), endX, endY))
	min = append(min, distanceTwoPoints(startX, startY, endX*(-1), endY))
	min = append(min, distanceTwoPoints(startX, startY, endX, endY*(-1)))
	min = append(min, distanceTwoPoints((2*m)-startX, startY, endX, endY))
	min = append(min, distanceTwoPoints(startX, startY, (2*m)-endX, endY))
	min = append(min, distanceTwoPoints(startX, (2*n)-startY, endX, endY))
	min = append(min, distanceTwoPoints(startX, startY, endX, (2*n)-endY))

	for _, target := range min {
		if minValue > target {
			minValue = target
		}
	}

	return minValue
}

func distanceTwoPoints(startX int, startY int, endX int, endY int) int {
	return ((startX - endX) * (startX - endX)) + ((startY - endY) * (startY - endY))
}
