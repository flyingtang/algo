package main

import (
	"fmt"
	"math"
)

// Goods 物品
type goods struct {
	weight uint // 商品重量
	value  uint // 商品价值
}

// goodlist 商品列表
type goodlist struct {
	gs    []goods // 商量列表
	total uint    // 商品列表总量

}

// backpack 背包
type backpack struct {
	currentWeight uint // 当前重量
	currentPrice  uint // 当前价值
	// t uint // 当前底基层
	capacity uint // 书包容量
	// sf       []uint //当前书包装的来自哪些楼层
}

// 初始化数据
func initData() (gl goodlist, bp backpack) {
	gl = goodlist{gs: []goods{{2, 6}, {5, 3}, {4, 5}, {2, 4}, {3, 10}}, total: 5} // 初始化商品列表
	bp = backpack{capacity: 10}                                                   // 初始化话背包
	return
}

// 版本1 //  最优解出来了 但是还不满足回溯的需求，每次都是一个背包的拷贝
func bag(gl goodlist, bp backpack, currentfloor uint) uint {
	// 已经到了顶楼
	nextFloor := currentfloor + 1
	// 已经最后一层了
	if nextFloor == gl.total {
		// 不选 就是不作为
		// backet(gl, bp, nextFloor)
		// 选
		if bp.capacity < bp.currentWeight+gl.gs[currentfloor].weight {
			// 超重不能选了
			return bp.currentPrice
		}
		bp.currentPrice += gl.gs[currentfloor].value
		bp.currentWeight += gl.gs[currentfloor].weight
		// bp.sf = append(bp.sf, currentfloor)
		fmt.Printf("当前楼层: %v, 当前重量: %v, 当前价值 %v \n", currentfloor, bp.currentWeight, bp.currentPrice)
		return bp.currentPrice
	}
	// 不选
	us := bag(gl, bp, nextFloor)
	// 选
	if bp.capacity < bp.currentWeight+gl.gs[currentfloor].weight {
		// 超重不能选了
		fmt.Printf("当前楼层: %v, 当前重量: %v, 当前价值 %v \n", currentfloor, bp.currentWeight, bp.currentPrice)
		return uint(math.Max(float64(us), float64(bp.currentPrice)))
	}
	bp.currentPrice += gl.gs[currentfloor].value
	bp.currentWeight += gl.gs[currentfloor].weight
	// bp.sf = append(bp.sf, currentfloor)
	s := bag(gl, bp, nextFloor)
	return uint(math.Max(float64(s), float64(us)))
}

func main() {
	// 背包最大容量
	gl, bp := initData()
	fmt.Println("背包最大价值:", bag(gl, bp, 0))
}
