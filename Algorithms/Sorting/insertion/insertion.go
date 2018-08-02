package insertion

// Sort will sort the given int slice using insertion sort
func Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}

	for j := 1; j < size; j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
}

// SortReverse will sort the given int slice using insertion sort in reverse order
func SortReverse(a []int) {
	size := len(a)
	if size < 2 {
		return
	}

	for j := 1; j < size; j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] < key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
}
