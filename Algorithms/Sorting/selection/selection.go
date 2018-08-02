package selection

// Sort will sort the given int slice using selection sort
// assume 0 based indexing
// for i = 0 to A.Length - 2
// 	min = i
// 	for j = i + 1 to A.Length - 1
// 		if A[j] < min
//			min = j
//  	swap a[i] with a[min]
func Sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i

		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
