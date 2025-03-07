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
