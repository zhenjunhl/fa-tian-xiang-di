package f

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomList 生成随机数 list
func RandomList(maxNumber int64, minNumber int64, number int64) []int64 {
	// 确保 minNumber 小于等于 maxNumber
	if minNumber > maxNumber {
		minNumber, maxNumber = maxNumber, minNumber
	}
	// 计算范围内的数字总数
	rangeSize := maxNumber - minNumber + 1
	// 检查是否有足够的数字来生成指定数量的不重复随机数
	if rangeSize < number {
		fmt.Println("指定范围内没有足够的数字来生成所需数量的不重复随机数")
		return nil
	}

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 用于存储已生成的随机数
	used := make(map[int64]bool)
	// 存储最终结果的切片
	result := make([]int64, 0, number)

	for len(result) < int(number) {
		// 生成一个指定范围内的随机数
		randNum := rand.Int63n(rangeSize) + minNumber
		// 检查该随机数是否已经被使用过
		if !used[randNum] {
			// 如果未使用过，则添加到结果切片中
			result = append(result, randNum)
			// 标记该随机数已被使用
			used[randNum] = true
		}
	}

	return result
}
