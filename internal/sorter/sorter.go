package sorter

func SortPhone(arrPhone []int) []int {
	if len(arrPhone) <= 1 {
		return arrPhone // Базовый случай: массив из 0 или 1 элемента уже отсортирован
	}
	// Делим массив на две части
	mid := len(arrPhone) / 2
	left := SortPhone(arrPhone[:mid])
	right := SortPhone(arrPhone[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var sortedList []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sortedList = append(sortedList, left[i])
			i++
		} else {
			sortedList = append(sortedList, right[j])
			j++
		}
	}
	sortedList = append(sortedList, left[i:]...)
	sortedList = append(sortedList, right[j:]...)
	return sortedList
}
