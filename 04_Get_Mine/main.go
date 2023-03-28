package main

import (
	"fmt"
	"sort"
)

type fiveMineralScoreSum struct {
	scoreSum     int
	mineralCount int
}

func main() {
	//picks := []int{0, 1, 0}
	//minerals := []string{"stone", "stone", "stone", "stone", "stone"}
	picks := []int{1, 3, 2}
	minerals := []string{"diamond", "diamond", "diamond", "iron", "iron", "diamond", "iron", "stone"}

	fmt.Println(solution(picks, minerals))
}

func solution(picks []int, minerals []string) (tired int) {
	pickCount := 0
	for _, pick := range picks {
		pickCount += pick
	}
	// 곡괭이가 부족할때 남은 광물 버리기
	if pickCount*5 < len(minerals) {
		minerals = minerals[:pickCount*5]
	}

	tired = 0
	fiveMineralScoreSumList := make([]fiveMineralScoreSum, 0)
	var fiveMinerals []string
	for len(minerals) > 5 {
		fiveMinerals = minerals[:5] // 처음 5개
		minerals = minerals[5:]     // 처음 5개 제외
		mydata := fiveMineralScoreSum{getScoreOfFiveMinerals(fiveMinerals), 5}
		fiveMineralScoreSumList = append(fiveMineralScoreSumList, mydata)
	}
	if len(minerals) != 0 {
		mydata := fiveMineralScoreSum{getScoreOfFiveMinerals(minerals), len(minerals)}
		fiveMineralScoreSumList = append(fiveMineralScoreSumList, mydata)
	}

	sort.Slice(fiveMineralScoreSumList, func(i, j int) bool {
		return fiveMineralScoreSumList[i].scoreSum > fiveMineralScoreSumList[j].scoreSum
	})

	for len(fiveMineralScoreSumList) > 0 {
		if picks[0] > 0 {
			if fiveMineralScoreSumList[0].scoreSum == 25 && fiveMineralScoreSumList[0].mineralCount == 5 {
				// 예외처리
				tired += fiveMineralScoreSumList[0].scoreSum / 5
				picks[0] -= 1
				fiveMineralScoreSumList = fiveMineralScoreSumList[1:]
			} else {
				tired += fiveMineralScoreSumList[0].scoreSum / 25
				fiveMineralScoreSumList[0].scoreSum = fiveMineralScoreSumList[0].scoreSum % 25
				tired += fiveMineralScoreSumList[0].scoreSum / 5
				fiveMineralScoreSumList[0].scoreSum = fiveMineralScoreSumList[0].scoreSum % 5
				tired += fiveMineralScoreSumList[0].scoreSum
				picks[0] -= 1
				fiveMineralScoreSumList = fiveMineralScoreSumList[1:]
			}
		} else if picks[1] > 0 {
			if fiveMineralScoreSumList[0].scoreSum == 5 && fiveMineralScoreSumList[0].mineralCount == 5 {
				// 예외처리
				tired += fiveMineralScoreSumList[0].scoreSum
				picks[1] -= 1
				fiveMineralScoreSumList = fiveMineralScoreSumList[1:]
			} else {
				tired += (fiveMineralScoreSumList[0].scoreSum / 25) * 5
				fiveMineralScoreSumList[0].scoreSum = fiveMineralScoreSumList[0].scoreSum % 25
				tired += fiveMineralScoreSumList[0].scoreSum / 5
				fiveMineralScoreSumList[0].scoreSum = fiveMineralScoreSumList[0].scoreSum % 5
				tired += fiveMineralScoreSumList[0].scoreSum
				picks[1] -= 1
				fiveMineralScoreSumList = fiveMineralScoreSumList[1:]
			}
		} else if picks[2] > 0 {
			tired += fiveMineralScoreSumList[0].scoreSum
			picks[2] -= 1
			fiveMineralScoreSumList = fiveMineralScoreSumList[1:]
		} else {
			return
		}
	}

	return
}

func getScoreOfFiveMinerals(minerals []string) int {
	score := 0
	for _, mineral := range minerals {
		if mineral == "diamond" {
			score += 25
		} else if mineral == "iron" {
			score += 5
		} else {
			score += 1
		}
	}
	return score
}
