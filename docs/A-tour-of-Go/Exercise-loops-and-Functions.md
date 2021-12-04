# 练习：循环与函数

## 题目

为了练习函数与循环，我们来实现一个平方根函数：用牛顿法实现平方根函数。

计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：

```go
z -= (z*z - x) / (2*z)
```

重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。

在提供的 `func Sqrt` 中实现它。无论输入是什么，对 z 的一个恰当的猜测为 1。 要开始，请重复计算 10 次并随之打印每次的 z 值。观察对于不同的值 x（1、2、3 ...）， 你得到的答案是如何逼近结果的，猜测提升的速度有多快。

提示：用类型转换或浮点数语法来声明并初始化一个浮点数值：

```go
z := 1.0
z := float64(1)
```

然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数大于还是小于 10。 尝试改变 z 的初始猜测，如 x 或 x/2。你的函数结果与标准库中的 [math.Sqrt](https://go-zh.org/pkg/math/#Sqrt) 接近吗？

*注：* 如果你对该算法的细节感兴趣，上面的 z² − x 是 z² 到它所要到达的值（即 x）的距离， 除以的 2z 为 z² 的导数，我们通过 z² 的变化速度来改变 z 的调整量。 这种通用方法叫做[牛顿法](https://zh.wikipedia.org/wiki/牛顿法)。 它对很多函数，特别是平方根而言非常有效。

## 解答

逼近方法的区别有初始猜测和迭代(循环)终止条件。

初始猜测可以为 1，x，x/2。

迭代(循环)终止条件可以是 10 次，可以是这种逼近方法趋于收敛的某个限度。

代码 1：

迭代终止条件是 10 次，测试了对于不同的值 x（1 ~ 10）的逼近的过程。

```go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// 初始猜测

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
```

输出结果(节选)：

```shell
本次 z 的值为 2
本次 z 的值为 1.75
本次 z 的值为 1.7321428571428572
本次 z 的值为 1.7320508100147276
本次 z 的值为 1.7320508075688772
本次 z 的值为 1.7320508075688774
本次 z 的值为 1.7320508075688772
本次 z 的值为 1.7320508075688774
本次 z 的值为 1.7320508075688772
本次 z 的值为 1.7320508075688774
```

代码 2：

迭代终止条件是这种逼近方法趋于收敛的某个限度(0.00000000001，已经是很小的值了)，

测试了对于不同的值 x（1 ~ 10）的逼近的过程，并用标准库与之对比。

```go
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
```

输出结果(节选)：

```shell
对 10 进行计算
标准库 math.Sqrt 计算得 3.1622776601683795
SqrtSelf 计算得 3.162277660168379, 初始猜测为 10, 迭代次数为 7
差为 4.440892098500626e-16
```

<!-- 网址或引用 -->
