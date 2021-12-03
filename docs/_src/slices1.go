package main

import (
	"math/rand"
	"time"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// 创建二维切片
	xy := make([][]uint8, dy)
	for i := range xy {
		xy[i] = make([]uint8, dx)
	}
	// 选择使用系统时间作为随机数种子，采用系统时间的毫秒数作为种子值
	seedNum := time.Now().UnixNano()
	// 创建随机数种子，种子的值决定了随机数的值
	rand.Seed(seedNum)
	// 获取一个小于 n 的随机数
	for i := range xy {
		for j := range xy[i] {
			xy[i][j] = uint8(rand.Intn(100))
		}
	}
	return xy
}

func main() {
	pic.Show(Pic)
}
