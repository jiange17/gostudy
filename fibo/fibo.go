package fibo

//GetFiboNum be sed to get the nth fibonacci num.
func GetFiboNum(n int) (res int) {
	if n < 0 {
		return -1
	}
	fib := []int{1, 1}
	for i := 2; i <= n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib[n]
}
