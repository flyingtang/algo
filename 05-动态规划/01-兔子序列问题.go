package dynamic

//        1  , n=1
//  fn    1  , n=2
//        f(n-1)+f(n-2), n>2

func fab(n int) int {
	if n < 1 {
		return -1
	} else if n <= 2 {
		return 1
	} else {
		return fab(n-1) + fab(n-2)
	}
}
