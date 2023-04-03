package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// line is of type  41-45,40-45 - save them into a-b,c-d
		// a-b is the first number range, c-d is the second number range
		// check if a-b contains c-d or vice versa
		// if yes, count++
		// if no, continue
		// fmt.Println(line)
		var a, b, c, d int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d)
		// fmt.Println(a, b, c, d)
		if range_contains_range(a, b, c, d) {
			count++
		}

	}
	fmt.Println(count)
}
func range_contains_range(a, b, c, d int) bool {
	return (a <= c && c <= b) || (a <= d && d <= b) || (c <= a && a <= d) || (c <= b && b <= d)
}
