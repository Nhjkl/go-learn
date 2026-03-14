package main

func Sum(arr [5]int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
