package recursion

func Fun1(x, y int) int {
	if x == 0 {
		return y
	}
	return Fun1(x-1, x+y)
}
