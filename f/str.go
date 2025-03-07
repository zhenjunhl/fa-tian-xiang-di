package f

import (
	"github.com/bytedance/sonic"
)

// ToString 转换为 String
/*
 * @param obj 要转换的对象
 * @return string 转换后的字符串
 */
func ToString(obj interface{}) string {
	str, _ := sonic.MarshalString(obj)
	return str
}

// MergeAndDeduplicateStrings 合并两个字符串数组并去除重复元素
/*
 * @param arr1 第一个字符串数组
 * @param arr2 第二个字符串数组
 * @return []string 合并后的字符串数组
 */
func MergeAndDeduplicateStrings(arr1, arr2 []string) []string {
	// 创建一个 map 用于存储已经出现过的元素
	uniqueMap := make(map[string]bool)
	var result []string

	// 遍历第一个数组
	for _, str := range arr1 {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			result = append(result, str)
		}
	}

	// 遍历第二个数组
	for _, str := range arr2 {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			result = append(result, str)
		}
	}

	return result
}
