# 练习：切片

## 题目

实现 `Pic`。它应当返回一个长度为 `dy` 的切片，其中每个元素是一个长度为 `dx`，元素类型为 `uint8` 的切片。

当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 `(x+y)/2`, `x*y`, `x^y`, `x*log(y)` 和 `x%(y+1)`。

（提示：需要使用循环来分配 `[][]uint8` 中的每个 `[]uint8`；请使用 `uint8(intValue)` 在类型之间转换；你可能会用到 `math` 包中的函数。）

## 解答

前提：[练习：导入Go语言远程包](Go-details/Import-remote-pakcages.md)

将任务分解为几个步骤：

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

代码1：

用随机数为其赋值

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

输出结果：

![Exercise-slices-pic1][Exercise-slices-pic1]

代码 2：

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 创建二维切片
	xy := make([][]uint8, dy)
	for i := range xy {
		xy[i] = make([]uint8, dx)
	}
	for i := range xy {
		for j := range xy[i] {
			xy[i][j] = uint8((i+j)/2) // 可修改
		}
	}
	return xy
}

func main() {
	pic.Show(Pic)
}
```

输出结果：

![Exercise-slices-pic2][Exercise-slices-pic2]

按照切片元素的赋值给出图像：

注：log 需要导入 math 包，赋值语句应为：`uint8(float64(i) * math.Log(float64(j)))`

<!-- tabs:start --> 

#### **English**

Hello!

#### **French**

Bonjour!

#### **Italian**

Ciao! 

<!-- tabs:end -->

|   赋值   |                     图像                      |
| :------: | :-------------------------------------------: |
|   随机   | ![Exercise-slices-pic1][Exercise-slices-pic1] |
| (x+y)/2  | ![Exercise-slices-pic2][Exercise-slices-pic2] |
|   x*y    | ![Exercise-slices-pic3][Exercise-slices-pic3] |
|   x^y    | ![Exercise-slices-pic4][Exercise-slices-pic4] |
| x*log(y) | ![Exercise-slices-pic5][Exercise-slices-pic5] |
| x%(y+1)  | ![Exercise-slices-pic6][Exercise-slices-pic6] |

还有很多花样可以玩。

<!-- 图片 -->

[Exercise-slices-pic1]:../_images/Exercise-slices-pic1.png
[Exercise-slices-pic2]:../_images/Exercise-slices-pic2.png
[Exercise-slices-pic3]:../_images/Exercise-slices-pic3.png

[Exercise-slices-pic4]:../_images/Exercise-slices-pic4.png

[Exercise-slices-pic5]:../_images/Exercise-slices-pic5.png
[Exercise-slices-pic6]:../_images/Exercise-slices-pic6.png
