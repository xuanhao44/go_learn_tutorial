# 练习：斐波纳契闭包

## 题目

让我们用函数做些好玩的事情。

实现一个 `fibonacci` 函数，它返回一个函数（闭包），该闭包返回一个[斐波纳契数列](https://zh.wikipedia.org/wiki/斐波那契数列) `(0, 1, 1, 2, 3, 5, ...)`。

## 解答

### 斐波那契数定义

$$
F_{0}=0 \\
F_{1}=1 \\
F_{n}=F_{{n-1}}+F_{{n-2}} (n≧2)
$$

### 代码

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
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

<!-- 网址或引用 -->