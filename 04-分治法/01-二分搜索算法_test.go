package binary

import "testing"

func TestBinarySearch(t *testing.T) {
	var testdata = []struct {
		sortlist []int
		x        int
		r        int
	}{{
		sortlist: []int{1, 4, 5, 6, 9},
		x:        1,
		r:        0,
	}, {
		sortlist: []int{1, 4, 5, 6, 9},
		x:        9,
		r:        4,
	}, {
		sortlist: []int{1, 4, 5, 6, 9},
		x:        1,
		r:        0,
	}}
	for i := 0; i < len(testdata); i++ {
		d := testdata[i]
		res := binarySearch(d.sortlist, d.x)
		if res == d.r {
			t.Log("测试通过")
		} else {
			t.Log("测试失败")
		}
	}
}
