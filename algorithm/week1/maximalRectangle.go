package main

func maximalRectangle(matrix [][]byte) (max int) {
	if len(matrix) == 0 {
		return 0
	}
	h, w := len(matrix), len(matrix[0])
	left := make([][]int, h+1)
	left[0] = make([]int, w)
	for i, row := range matrix {
		left[i+1] = make([]int, w)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i+1][j] = 1
			} else {
				left[i+1][j] = left[i+1][j-1] + 1
			}
		}
	}

	for i:=0;i<w;i++ {
		st := make([]int, 0)
		for j:=h;j>=0;j-- {
			var l int
			for len(st) > 0 && st[len(st)-2] >= left[j][i] {
				l += st[len(st)-1]
				if l*st[len(st)-2] > max {
					max = l*st[len(st)-2]
				}
				st = st[:len(st)-2]
			}
			st = append(st, left[j][i], l+1)
		}
	}

	return
}
