package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// 无论输入是什么，对 z 的一个恰当的猜测为 1

	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("本次 z 的值为 %g\n", z)
	}
	return z
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("x = %d, %g\n", i, Sqrt(float64(i)))
	}
}
