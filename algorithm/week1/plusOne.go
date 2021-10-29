package main

func plusOne(digits []int) []int {
	var jin = 1
	for i:=len(digits)-1;i>=0;i-- {
		if digits[i]+jin > 9 {
			digits[i] = 0
		} else {
			digits[i] += jin
			jin = 0
		}
	}
	if jin == 1 {
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	}

	return digits
}
