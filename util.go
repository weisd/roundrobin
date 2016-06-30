package roundrobin

// 两个数的最大公约数 欧几里德算法
func gcd(a, b int) int {

	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a

	}
	return gcd(b, a%b)
}

// n个数的最大公约数算法
// 说明:
// 把n个数保存为一个数组
// 参数为数组的指针和数组的大小(需要计算的数的个数)
// 然后先求出gcd(a[0],a[1]), 然后将所求的gcd与数组的下一个元素作为gcd的参数继续求gcd
// 这样就产生一个递归的求ngcd的算法
func ngcd(ints []int) int {

	n := len(ints)

	if n == 1 {
		return ints[0]
	}

	return gcd(ints[n-1], ngcd(ints[0:n-1]))
}
