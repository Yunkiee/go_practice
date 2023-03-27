package _2_Hotel_room

import (
	"strconv"
	"strings"
)

func solution(book_time [][]string) int {
	//book_time_int := timeStringToMinuteInt(book_time)
	//isBooked := false
	//roomCount := 0
	//for _, book := range book_time_int {
	//	if book
	//
	//}

	return 0
}

func timeStringToMinuteInt(book_time [][]string) (book_int [][]int) {
	// ["00:10","01:50"] -> [10,110]
	for index, target := range book_time {
		start_time := strings.Split(target[0], ":")
		start_hour, _ := strconv.Atoi(start_time[0])
		start_minute, _ := strconv.Atoi(start_time[1])

		end_time := strings.Split(target[1], ":")
		end_hour, _ := strconv.Atoi(end_time[0])
		end_minute, _ := strconv.Atoi(end_time[1])

		start_int := start_hour*60 + start_minute
		end_int := end_hour*60 + end_minute

		book_int[index][0] = start_int
		book_int[index][1] = end_int
	}

	return
}
