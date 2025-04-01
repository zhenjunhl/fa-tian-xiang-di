package f

// SplitNumberByArr 将数字拆分为10个数字，总和为n,前9个数字相同，最后一位保证等于n
func SplitNumberByArr(n int64, s int64) []int64 {
	if n < s {
		return []int64{n}
	}
	base := n / s
	rem := n % s
	result := make([]int64, s)
	if rem == 0 {
		for i := 0; i < int(s); i++ {
			result[i] = base
		}
	} else if rem < s-1 {
		for i := 0; i < int(s-1); i++ {
			result[i] = base
		}
		result[s-1] = base + rem
	} else {
		for i := 0; i < int(s-1); i++ {
			result[i] = base + 1
		}
		result[s-1] = base + (rem - s - 1)
	}
	return result
}
