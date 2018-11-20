package dynamic

import "testing"

func TestFab(t *testing.T) {
	var testDatas = []struct {
		n   int
		res int
	}{
		{-1, -1},
		{0, -1},
		{1, 1},
		{2, 1},
		{3, 2},
		{8, 21},
	}
	for i := 0; i < len(testDatas); i++ {
		data := testDatas[i]
		res := fab(data.n)
		if res != data.res {
			t.Errorf("%v is failed, actual is %v", data, res)
		}
	}
}
