package f

// SplitNumberByArr 将数字拆分为10个数字，总和为n,前9个数字相同，最后一位保证等于n
func SplitNumberByArr(n int64) []int64 {
	if n < 10 {
		return []int64{n}
	}
	base := n / 10
	rem := n % 10
	result := make([]int64, 10)
	if rem == 0 {
		for i := 0; i < 10; i++ {
			result[i] = base
		}
	} else if rem < 9 {
		for i := 0; i < 9; i++ {
			result[i] = base
		}
		result[9] = base + rem
	} else {
		for i := 0; i < 9; i++ {
			result[i] = base + 1
		}
		result[9] = base + (rem - 9)
	}
	return result
}
