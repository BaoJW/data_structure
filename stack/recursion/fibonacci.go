package recursion

// 斐波那契数列的实现
/*
 F(n) = {
			0, 当n=0
			1, 当n=1
			F(n-1)+F(n-2), 当n>1
		}
*/

// 斐波那契数列的迭代实现(非递归实现)
// 很显然他的时间复杂度是O(N)
func FibByIteration(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a := 0
	b := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = a + b
		a, b = b, result
	}

	return result
}

// 斐波那契数列的递归实现
// 他的时间复杂度是O(2^N)
func FibByRecursion(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return FibByRecursion(n-1) + FibByRecursion(n-2)

}
