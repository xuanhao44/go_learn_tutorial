# 练习：斐波纳契闭包

## 题目

让我们用函数做些好玩的事情。

实现一个 `fibonacci` 函数，它返回一个函数（闭包），该闭包返回一个[斐波纳契数列](https://zh.wikipedia.org/wiki/斐波那契数列) `(0, 1, 1, 2, 3, 5, ...)`。

## 解答

### 斐波那契数定义

$$
F_{1} = 1 \\
F_{n} = F_{n-1} + F_{n-2} (n≧2)
$$

###  多重赋值

在初始化赋值时可以有：

```go
var i, j int = 1, 2
```

那么如何理解：

```go
x, y = y, x
```

```go
x, y = y, x+y
```

代码 1：

```go
package main

import "fmt"

func main() {
	x := 3
	y := 5
	x, y = y, x
	fmt.Println(x, y)
}
```

输出：

```shell
5 3
```

代码 2：

```go
package main

import "fmt"

func main() {
	x := 3
	y := 5
	x, y = y, x+y
	fmt.Println(x, y)
}
```

输出：

```
5 8
```

### 最终代码

这里直接给出了官方的答案。

```go
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 1; i <= 10; i++ {
		fmt.Println(f())
	}
}
```

我们把前 3 次循环展示出来。
$$
i = 1, x = y = 1, y = x+y = 0 + 1 = 1, return \, 1; \\
i = 2, x = y = 1, y = x+y = 1 + 1 = 2, return \, 1; \\
i = 3, x = y = 2, y = x+y = 1 + 2 = 3, return \, 2;
$$

每次函数返回的是结果，也就是 x，每次都会计算下一次的结果，也就是 y。

$x = y$ 解释：

所以每次的 x 只需要承接 y。

$y = x+y$ 解释：

具体根据：$F_{n+1} = F_{n} + F_{n-1}$

等式右侧的 x 意味着这一次的结果，即 $F_{n}$

等式右侧的 y 意味着上一次的结果，即 $F_{n-1}$

等式左侧的 y 意味着下一次的结果，即 $F_{n+1}$

<!-- 网址或引用 -->