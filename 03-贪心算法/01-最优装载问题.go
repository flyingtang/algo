package greedy

import (
	"fmt"
	"sort"
)

// 背景：加勒比海盗以打劫为生，有一天截获了一艘装满了各种金银珠宝的超级大船，每一件宝贝都价值连城，
// 万一打碎了，就失去了价值。但是珠宝太多，海盗们自己船装不下呢，怎么办？
// 贪心的海盗们就开始琢磨：
// 		既然我们船的载重是C，每一件的重量为wi,又不能拆分，我们怎样才能竟可能多的把宝贝装上船上呢？

// 问题分析： 既然每一件宝贝都价值连城，海盗船的载重只能为C, 那么每次都选择重量最小的放进船里,不就得到了最优解了么？
// 思路：
// 1. 排序： 从小到大排序
// 2. 选择： 挨个选择装上去，知道不能装

func pirateship(w []float64, c float64) int {
	fmt.Println(w)
	// 排序
	sort.Float64s(w)
	fmt.Println(w)
	// 选择
	length := len(w)
	tempw := 0.0
	for i := 0; i < length; i++ {
		tempw += w[i]
		if tempw > c {
			return i
		}
	}
	return length
}
