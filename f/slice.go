package f

// Uniq 返回一个新的片段，其中包含给定集合中的唯一元素。
/*
 * @param collection 要处理的集合。
 * @return 包含唯一元素的新片段。
 */
func Uniq[T comparable, Slice ~[]T](collection Slice) Slice {
	result := make(Slice, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for i := range collection {
		if _, ok := seen[collection[i]]; ok {
			continue
		}

		seen[collection[i]] = struct{}{}
		result = append(result, collection[i])
	}

	return result
}
func CalcPageSize(total int, page int) int {
	// 非法页数直接返回 0
	if page <= 0 {
		return 0
	}
	// 刚好整除
	if total%page == 0 {
		return total / page
	}
	// 不整除 → 向上取整（最后一页不满也算一页）
	return total/page + 1
}
