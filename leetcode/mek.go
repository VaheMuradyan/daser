package main

import "fmt"

func main() {
	fmt.Println(12 / 10)
	fmt.Println(12 % 10)
}

func twoSum(nums []int, target int) []int {
	var answer []int
	firstTime := true

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				if firstTime {
					answer = append(answer, i)
					answer = append(answer, j)
					firstTime = false
				} else {
					if !exist(answer, i) {
						answer = append(answer, i)
					}
					if !exist(answer, j) {
						answer = append(answer, j)
					}
				}
			}
		}
	}

	return answer
}

func exist(values []int, number int) bool {
	for _, val := range values {
		if val == number {
			return true
		}
	}

	return false
}
