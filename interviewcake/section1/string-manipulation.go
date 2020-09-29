package main

import (
	"fmt"
	"sort"
)

func main() {
	meetings := []Meeting{{1, 2}, {3, 6}, {2, 5}, {7, 9}}

	meetingsResult := mergeRanges(meetings)
	fmt.Println(meetingsResult)

	fmt.Println(stringReverse("cash"))
}

type Meeting struct {
	startTime int
	endTime   int
}

func mergeRanges(meetings []Meeting) []Meeting {
	// first sort meetings by their start time
	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i].startTime < meetings[j].startTime
	})

	for i := 1; i < len(meetings); i++ {
		a := meetings[i-1]
		b := meetings[i]

		if b.endTime <= a.endTime {
			meetings = append(meetings[:i], meetings[i+1:]...)
			i--
		} else if b.startTime <= a.endTime {
			meetings[i-1] = Meeting{a.startTime, b.endTime}
			meetings = append(meetings[:i], meetings[i+1:]...)
			i--
		}
	}

	return meetings
}

func stringReverse(input string) string {
	inputBytes := []byte(input)
	for i := 1; i <= len(inputBytes)/2; i++ {
		stash := inputBytes[i-1]
		inputBytes[i-1] = inputBytes[len(inputBytes)-i]
		inputBytes[len(inputBytes)-i] = stash
	}

	return string(inputBytes)
}