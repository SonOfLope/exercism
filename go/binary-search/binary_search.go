package binarysearch

func SearchInts(slice []int, key int) int {
	l, r := 0, len(slice)-1
	for l <= r {
		m := (l + r) / 2
		if slice[m] < key {
			l = m + 1
		} else if slice[m] > key {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
