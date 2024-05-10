package utils

// RawSlice 切片 中删除 delSlice 中的所有元素
// delList 不应比 RawSlice 长, 否则会 Panic!
// 所有的切片元素都不应有重复, 否则可能得不到想要的结果
func SliceRemove[T comparable](RawSlice []T, removeSlice []T) []T {
	lastIdx := len(RawSlice)
	for delIdx := 0; delIdx < len(removeSlice); delIdx++ {
		for findingIdx := 0; findingIdx < lastIdx; findingIdx++ {
			if removeSlice[delIdx] == RawSlice[findingIdx] {
				lastIdx -= 1
				RawSlice[findingIdx] = RawSlice[lastIdx]
				break
			}
		}
	}
	return RawSlice[:lastIdx]
}

// 切片中的指定元素计数
func SliceCounter[T comparable](RawSlice []T, element T) int {
	if len(RawSlice) == 0 {
		return 0
	}
	count := 0
	for _, eachElement := range RawSlice {
		if eachElement == element {
			count++
		}
	}
	return count
}

// 切片去除重复元素
func SliceDeduplicate[T anyInt](RawSlice []T) []T {
	result := make([]T, 0, len(RawSlice))
	tempMap := map[T]struct{}{}
	for _, item := range RawSlice {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

