package greedy

import "testing"

func TestPirateship(t *testing.T) {
	var datas = []struct {
		w     []float64
		c     float64
		count int
	}{
		{[]float64{4, 10, 7, 11, 3, 5, 14, 2}, 30, 5},
	}
	for i := 0; i < len(datas); i++ {
		data := datas[i]
		count := pirateship(data.w, data.c)
		t.Log(count)
		if data.count == count {
			t.Log("pass")
		} else {
			t.Fatal("TestPirateship failed")
		}
	}
}

// go test -v -timeout 30s github.com/txg5214/algo/03-贪心算法 -run ^TestPirateship$
