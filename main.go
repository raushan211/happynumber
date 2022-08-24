package main

import "fmt"

func main() {
	n := 19

	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {
	result := check(n)
	if result == 1 || result == 7 {
		return true
	}
	return false

}
func check(n int) int {
	var r int
	sum := 0

	for n > 0 {
		r = n % 10
		sum = sum + (r * r)
		n = n / 10
	}
	for sum > 9 {
		sum = check(sum)
	}
	return sum
}
