package main

import (
	"fmt"
	"math"
)

func SqrtSelf(x, guess float64) (float64, int) {
	z := guess
	// 无论输入是什么，对 z 的一个恰当的猜测为 1
	min := 0.00000000001
	// 限度
	num := 0
	// 迭代次数

	result := z*z - x
	if result < 0 {
		result = -result
	}
	// 初值

	for result > min {
		result = z*z - x
		if result < 0 {
			result = -result
		}
		z -= (z*z - x) / (2 * z)
		num++
	}
	return z, num
}

func main() {
	var SqrtResult float64 = 0.0
	var Num int = 0
	var Guess float64 = 0.0

	for i := 1; i <= 10; i++ {
		Guess = float64(i) // 初始猜测值可更改
		SqrtResult, Num = SqrtSelf(float64(i), Guess)
		fmt.Printf("对 %d 进行计算\n", i)
		fmt.Printf("标准库 math.Sqrt 计算得 %g\n", math.Sqrt(float64(i)))
		fmt.Printf("SqrtSelf 计算得 %g, 初始猜测为 %g, 迭代次数为 %d\n", SqrtResult, Guess, Num)
		fmt.Printf("差为 %g\n\n", math.Sqrt(float64(i))-SqrtResult)
	}
}
