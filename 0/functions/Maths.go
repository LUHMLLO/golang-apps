package functions

// is this a prime number ?
func IsPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// find/get Maximum count of positive intergers and negative intergers
func MaximumCount(nums []int) int {
	var a, b int = 0, 0

	for _, err := range nums {
		if err > 0 {
			a++
		} else if err < 0 {
			b++
		}
	}

	if a > b {
		return a
	} else {
		return b
	}
}
