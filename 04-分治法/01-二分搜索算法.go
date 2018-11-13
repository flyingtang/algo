package binary

// 以前我们玩过一种猜数字游戏，给定10以内的任何一个数，然人猜大了还是小了，最终猜到给定的数
//binarySearch  arrs 有序数组 x 要查找的数字
func binarySearch(arrs []int, x int) (pos int) {
	low, hight := 0, len(arrs)-1
	for low <= hight {
		middle := (hight-low)/2 + low
		if arrs[middle] == x {
			return middle
		} else if arrs[middle] > x {
			hight = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
