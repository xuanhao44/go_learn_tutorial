# 练习：切片

## 题目

实现 `Pic`。它应当返回一个长度为 `dy` 的切片，其中每个元素是一个长度为 `dx`，元素类型为 `uint8` 的切片。

当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 `(x+y)/2`, `x*y`, `x^y`, `x*log(y)` 和 `x%(y+1)`。

（提示：需要使用循环来分配 `[][]uint8` 中的每个 `[]uint8`；请使用 `uint8(intValue)` 在类型之间转换；你可能会用到 `math` 包中的函数。）

## 解答

1. 用 `make` 创建二维切片 `[][]uint8`

   ```go
   // 创建二维切片
   xy := make([][]uint8, dy)
   for i := range xy {
   	xy[i] = make([]uint8, dx)
   }
   ```

2. 用循环来分配 `[][]uint8` 中的 `[]uint8`

   ```go
   for i := range xy {
   	for j := range xy[i] {
   		xy[i][j] = uint8(1)
   	}
   }
   ```

3. 用随机数来为 `[]uint8` 赋值

   ```go
   seedNum := time.Now().UnixNano()
   // 选择使用系统时间作为随机数种子，采用系统时间的毫秒数作为种子值
   rand.Seed(seedNum)
   // 创建随机数种子，种子的值决定了随机数的值
   rand.Intn(n)
   // 获取一个小于 n 的随机数
   ```

代码1：用随机数为其赋值

```go
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
```

输出结果：很长很长的 base64 码

<!-- 网址或引用 -->