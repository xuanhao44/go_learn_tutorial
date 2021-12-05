# 流程控制语句：for、if、else、switch 和 defer

## 1 for

Go 只有一种循环结构：`for` 循环。

基本的 `for` 循环由三部分组成，它们用分号隔开：

- **初始化语句：在第一次迭代前执行**
- **条件表达式：在每次迭代前求值**
- **后置语句：在每次迭代的结尾执行**

初始化语句通常为**一句短变量声明**，该变量声明仅在 `for` 语句的作用域中可见。

一旦条件表达式的布尔值为 `false`，循环迭代就会终止。

**注意**：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面的三个构成部分外**没有小括号**， 大括号 `{ }` 则是必须的。

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```

初始化语句和后置语句是可选的。

```go
package main

import "fmt"

func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}
```

### 1.1 for 是 Go 中的 “while”

此时你可以**去掉分号**，因为 C 的 `while` 在 Go 中叫做 `for`。

```go
package main

import "fmt"

func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
```

### 1.2 无限循环

如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。

```go
package main

func main() {
    for {
    }
}
```

输出：

```shell
timeout running program
```

## 2 if

Go 的 `if` 语句与 `for` 循环类似，表达式外**无需小括号** `( )` ，而大括号 `{ }` 则是必须的。

```go
package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
}
```

### 2.1 if 的简短语句

同 `for` 一样， `if` 语句可以在条件表达式前**执行一个简单的语句**。该语句声明的变量作用域仅在 `if` 之内。

```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
```

（在最后的 `return` 语句处使用 `v` 看看。）

错误输出：

```shell
./prog.go:12:9: undefined: v
```

显示 `v` 未定义。

### 2.2 if 和 else

在 `if` 的简短语句中声明的变量同样可以在任何对应的 `else` 块中使用。

（在 `main` 的 `fmt.Println` 调用开始前，两次对 `pow` 的调用均已执行并返回其各自的结果。）

```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // 这里开始就不能使用 v 了
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
```

##  练习：循环与函数

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

详细解答：[循环与函数](A-tour-of-Go/Exercise-loops-and-Functions.md)

## 3 switch

`switch` 是编写一连串 `if - else` 语句的简便方法。它运行**第一个值等于条件表达式的 case 语句**。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go **只运行选定的 case，而非之后所有的 case**。 实际上，Go **自动提供了在这些语言中每个 case 后面所需的 `break` 语句**。 除非以 `fallthrough` 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 **switch 的 case 无需为常量，且取值不必为整数**。

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.\n", os)
    }
}
```

程序输出的结果根据你运行程序的平台而定。

### 3.1 switch 的求值顺序

switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

（例如，

```go
switch i {
case 0:
case f():
}
```

在 `i==0` 时 `f` 不会被调用。）

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
}
```

*注意：* Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始，该值的意义留给读者去发现。

### 3.2 没有条件的 switch

没有条件的 switch 同 `switch true` 一样。

**这种形式能将一长串 if-then-else 写得更加清晰。**

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
```

## 4 defer

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
```

输出结果：

```go
hello
world
```

*有点异步的感觉了*

推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

更多关于 defer 语句的信息，请阅读[此博文](http://blog.go-zh.org/defer-panic-and-recover)。

```go
package main

import "fmt"

func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
```
