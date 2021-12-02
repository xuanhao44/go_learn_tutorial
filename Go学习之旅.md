# Go学习之旅

**本文地址**：[xuanhao44/Go_Learn](https://github.com/xuanhao44/Go_Learn)

## 0 教程与简介

### 0.1 简介

Go 语言中文教程：[Go 语言之旅 (go-zh.org)](https://tour.go-zh.org/list)，可以在线学。本文基本把教程全文搬运过来了，所以会比较长。

但是细节上略有不同，添加了自己的一些笔记，增添了目录结构，以供自己查阅。

但不幸的是，这份教程即使是在一个初学者看来也是相当简略的。照着教程学是没有问题的(毕竟你只是运行一下他的代码！)，遇到自己动手的练习就问题百出了。所以还是有必要自己去查询资料。

### 0.2 目录

- 基础

  - [包、变量和函数](https://tour.go-zh.org/basics)

    学习 Go 程序的基本结构。

  - [流程控制语句：for、if、else、switch 和 defer](https://tour.go-zh.org/flowcontrol)

    学习如何使用条件、循环、分支和推迟语句来控制代码的流程。

  - [更多类型：struct、slice 和映射](https://tour.go-zh.org/moretypes)

    学习如何基于现有类型定义新的类型：本节课涵盖了结构体、数组、切片和映射。

- [方法和接口](https://tour.go-zh.org/methods)

  学习如何为类型定义方法；如何定义接口；以及如何将所有内容贯通起来。本节课包含了方法和接口，可以用这种构造来定义对象及其行为。

- [并发](https://tour.go-zh.org/concurrency)

  作为语言的核心部分，Go 提供了并发的特性。这一部分概览了 goroutine 和 channel，以及如何使用它们来实现不同的并发模式。Go 将并发结构作为核心语言的一部分提供。本节课程通过一些示例介绍并展示了它们的用法。

### 0.3 练习

Go学习之旅中穿插着练习，将其抽出单独处理。

1. [循环与函数](../练习/循环与函数/循环与函数.md)
2. [切片](../练习/切片/切片.md)

练习答案参考[Go 学习之旅 （tour.golang.org） Google官方答案 - 掘金 (juejin.cn)](https://juejin.cn/post/6844903806648451080)

## 1 [包、变量和函数](https://tour.go-zh.org/basics)

### 1.1 包

每个 Go 程序都是由包构成的。

程序从 `main` 包开始运行。

本程序通过导入路径 `"fmt"` 和 `"math/rand"` 来使用这两个包。

按照约定，包名与导入路径的最后一个元素一致。例如，`"math/rand"` 包中的源码均以 `package rand` 语句开始。

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(1))
}
```

*注意：* 此程序的运行环境是固定的，因此 `rand.Intn` 总是会返回相同的数字。（要得到不同的数字，需为生成器提供不同的种子数，参见 [`rand.Seed`](https://go-zh.org/pkg/math/rand/#Seed)。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）

#### 1.1.1 导入

此代码用圆括号组合了导入，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

```go
import "fmt"
import "math"
```

不过使用分组导入语句是更好的形式。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

尝试：将语句修改为：

```go
fmt.Printf("Now you have a problems.\n")
```

再试着执行一次。

得到错误输出：

```shell
./prog.go:5:2: imported and not used: "math"
```

使用标准格式引用包，但是代码中却没有使用包，编译器会报错。

#### 1.1.2 导出名

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，`Pizza` 就是个已导出名，`Pi` 也同样，它导出自 `math` 包。

`pizza` 和 `pi` 并未以大写字母开头，所以它们是未导出的。

在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

执行代码，观察错误输出。

然后将 `math.pi` 改名为 `math.Pi` 再试着执行一次。

错误代码

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
}
```

错误输出

```shell
./prog.go:9:14: cannot refer to unexported name math.pi
```

应将 `math.pi` 改名为 `math.Pi` 。

### 1.2 函数

函数可以没有参数或接受多个参数。

在本例中，`add` 接受两个 `int` 类型的参数。

注意**类型在变量名 之后**。

（参考 [这篇关于 Go 语法声明的文章](http://blog.go-zh.org/gos-declaration-syntax)了解这种类型声明形式出现的原因。）

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

1. `add` 接受两个 `int` 类型的参数，写法为：`x int, y int`。
2. `add` 函数本身也是把返回值的类型放在 `add` 这个函数名的后面。

#### 1.2.1 缩写

`x int, y int` 被缩写为 `x, y int`。当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

在本例中，

```go
x int, y int
```

被缩写为

```go
x, y int
```

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

#### 1.2.2 多值返回

函数可以返回任意数量的返回值。

`swap` 函数返回了两个字符串。

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

#### 1.2.2 命名返回值

Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

没有参数的 `return` 语句返回已命名的返回值。也就是 `直接` 返回。

直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return y, x
}

func main() {
	fmt.Println(split(17))
}
```

### 1.3 变量

#### 1.3.1 变量声明

`var` 语句用于**声明一个变量列表**，跟函数的参数列表一样，**类型在最后**。

就像在这个例子中看到的一样，`var` 语句可以出现在**包**或**函数**级别。

- `var c, python, java bool` 在包 `package main` 中
- `var i int` 在函数 `func main()` 中

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

#### 1.3.2 变量的初始化

变量声明可以包含初始值，每个变量对应一个。*写法好怪*

```go
var i, j int = 1, 2
```

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

```go
var c, python, java = true, false, "no!"
```

这一句中 `c` 和 `python` 被赋予了 `true` 和 `false` 的初始化值，所以它们从从初始值中获得类型 `bool`；`java` 被赋予了 `"no!"` 的初始化值，所以它们从从初始值中获得类型 `string`。

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

#### 1.3.3 短变量声明

**在函数中**，简洁赋值语句 `:=` 可在类型明确的地方代替 `var` 声明。

函数外的每个语句都必须以关键字开始（`var`, `func` 等等），因此 `:=` 结构不能在函数外使用。*PS：这两者不可以共用*

共用后错误输出：

```shell
./prog.go:7:12: syntax error: unexpected :=, expecting =
```

#### 1.3.4 基本类型

Go 的基本类型有

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
```

本例展示了几种类型的变量。 同导入语句一样，**变量声明也可以“分组”成一个语法块**。

`int`, `uint` 和 `uintptr` 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 `int` 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

*居然有复数，太赞了!*

#### 1.3.5 零值

没有明确初始值的变量声明会被赋予它们的 **零值**。

零值是：

- 数值类型为 `0`，
- 布尔类型为 `false`，
- 字符串为 `""`（空字符串）。

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

#### 1.3.6 类型转换

表达式 `T(v)` 将值 `v` 转换为类型 `T`。

一些关于数值的转换：

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

与 C 不同的是，Go 在不同类型的项之间赋值时**需要显式转换**。试着移除例子中 `float64` 或 `uint` 的转换看看会发生什么。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = f
	fmt.Println(x, y, z)
}
```

尝试：将 `int` 类型的变量的值赋给 `uint` 类型的变量 `z`

```go
var z uint = x
```

出现了错误输出：

```shell
./prog.go:9:6: cannot use x (type int) as type uint in assignment
```

所以不管是即使是变量类型提升的情况也是需要显式转换的，与 c 不同。

#### 1.3.7 类型推导

在声明一个变量而不指定其类型时（即使用不带类型的 `:=` 语法或 `var =` 表达式语法），变量的类型由右值推导得出。

当右值声明了类型时，新变量的类型与其相同：

```go
var i int
j := i // j 也是一个 int
```

不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 `int`, `float64` 或 `complex128` 了，这取决于常量的精度：

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

尝试修改示例代码中 `v` 的初始值，并观察它是如何影响类型的。

```go
package main

import "fmt"

func main() {
	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)
}
```

`v := 42`，输出：`v is of type int`

`v := 32.0`，输出：`v is of type float64`

`v := 0.867 + 0.5i`，输出：`v is of type complex128`

现在你知道如何输出变量类型了：**`%T`**

####  1.3.8 常量

常量的声明与变量类似，只不过是使用 `const` 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 `:=` 语法声明。

```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

#### 1.3.9 数值常量

数值常量是高精度的 **值**。

一个未指定类型的常量由上下文来决定其类型。

（`int` 类型最大可以存储一个 64 位的整数，有时会更小。）

（`int` 可以存放最大64位的整数，根据平台不同有时会更少。）

```go
package main

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

再尝试一下输出 `needInt(Big)` 吧。

```go
fmt.Println(needInt(Big))
```

得到错误输出：

```shell
./prog.go:21:21: constant 1267650600228229401496703205376 overflows int
```

这说明 `Big` 常量太大，超出了 `int` 的范围。

## 2 [流程控制语句：for、if、else、switch 和 defer](https://tour.go-zh.org/flowcontrol)

### 2.1 for

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

#### 2.1.1 for 是 Go 中的 “while”

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

#### 2.1.2 无限循环

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

### 2.2 if

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

#### 2.2.1 if 的简短语句

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

#### 2.2.2 if 和 else

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

#### 2.2.3 练习：循环与函数

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

详细解答：[循环与函数](../练习/循环与函数/循环与函数.md)

### 2.3 switch

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

#### 2.3.1 switch 的求值顺序

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

#### 2.3.2 没有条件的 switch

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

### 2.4 defer

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

## 3 [更多类型：struct、slice 和映射](https://tour.go-zh.org/moretypes)

### 3.1 指针

Go 拥有指针。指针保存了值的内存地址。

类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`。

```go
var p *int
```

`&` 操作符会生成一个指向其操作数的指针。

```go
i := 42
p = &i
```

`*` 操作符表示指针指向的底层值。

```go
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```

这也就是通常所说的“间接引用”或“重定向”。

与 C 不同，Go **没有指针运算**。

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
```

### 3.2 结构体

一个结构体（`struct`）就是一组字段（field）。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

#### 3.2.1 结构体字段

结构体字段使用**点号**来访问。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

#### 3.2.2 结构体指针

结构体字段可以通过结构体指针来访问。

如果我们有一个指向结构体的指针 `p`，那么可以通过 `(*p).X` 来访问其字段 `X`。不过这么写太啰嗦了，所以语言也允许我们使用**隐式间接引用**，直接写 `p.X` 就可以。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

#### 3.2.3 结构体文法

结构体文法通过**直接列出字段的值来新分配一个结构体**。

使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 `&` 返回一个指向结构体的指针。

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

### 3.3 数组

类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组。

表达式

```go
var a [10]int
```

会将变量 `a` 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此**数组不能改变大小**。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

### 3.4 切片

每个数组的大小都是固定的。而切片则为数组元素提供**动态大小的、灵活的视角**。在实践中，**切片比数组更常用**。

类型 `[]T` 表示一个元素类型为 `T` 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

```go
a[low : high]
```

它会选择**一个半开区间，包括第一个元素，但排除最后一个元素**。

以下表达式创建了一个切片，它包含 `a` 中下标从 1 到 3 的元素：

```go
a[1:4]
```

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

#### 3.4.1  切片就像数组的引用

切片并不存储任何数据，它只是描述了底层数组中的一段。

**更改切片的元素会修改其底层数组中对应的元素**。与它共享底层数组的切片都会观测到这些修改。

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

#### 3.4.2 切片文法

切片文法类似于没有长度的数组文法。

这是一个数组文法：

```go
[3]bool{true, true, false}
```

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

```go
[]bool{true, true, false}
```

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

#### 3.4.3 切片的默认行为

在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的默认值为 `0`，上界则是该切片的长度。

对于数组

```go
var a [10]int
```

来说，以下切片是等价的：

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

#### 3.4.4 切片的长度与容量

切片拥有 **长度** 和 **容量**。

- 切片的长度就是它所包含的元素个数。

- 切片的容量是**从它的第一个元素开始数，到其底层数组元素末尾的个数**。

切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

你可以通过重新切片来扩展一个切片，给它提供足够的容量。试着修改示例程序中的切片操作，向外扩展它的容量，看看会发生什么。

*显然容量是不能超出底层数组的长度的，也就是拓展长度也是有极限的*

把 切片 `s` 拓展到 10：

```go
s = s[:10]
```

错误输出：

```shell
panic: runtime error: slice bounds out of range [:10] with capacity 6
```

#### 3.4.5 nil 切片

切片的零值是 `nil`。

nil 切片的长度和容量为 0 且没有底层数组。

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

#### 3.4.6 用 make 创建切片

切片可以用内建函数 `make` 来创建，这也是你创建动态数组的方式。

`make` 函数会分配一个**元素为零值**的数组并返回一个引用了它的切片：

```go
a := make([]int, 5)  // len(a)=5
```

要指定它的容量，需向 `make` 传入第三个参数：

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
```

#### 3.4.7 切片的切片

切片可包含任何类型，甚至包括其它的切片。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

#### 3.4.8 向切片追加元素

为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 `append` 函数。内建函数的[文档](https://go-zh.org/pkg/builtin/#append)对此函数有详细的介绍。

```go
func append(s []T, vs ...T) []T
```

`append` 的第一个参数 `s` 是一个元素类型为 `T` 的切片，其余类型为 `T` 的值将会追加到该切片的末尾。

`append` 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 `s` 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

（要了解关于切片的更多内容，请阅读文章 [Go 切片：用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals)。）

#### 3.4.9 Range

`for` 循环的 `range` 形式可遍历切片或映射。

当使用 `for` 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

可以将下标或值赋予 `_` 来忽略它。

```go
for i, _ := range pow
for _, value := range pow
```

若你只需要索引，忽略第二个变量即可。

```go
for i := range pow
```

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

#### 3.4.10 练习：切片

实现 `Pic`。它应当返回一个长度为 `dy` 的切片，其中每个元素是一个长度为 `dx`，元素类型为 `uint8` 的切片。当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 `(x+y)/2`, `x*y`, `x^y`, `x*log(y)` 和 `x%(y+1)`。

（提示：需要使用循环来分配 `[][]uint8` 中的每个 `[]uint8`；请使用 `uint8(intValue)` 在类型之间转换；你可能会用到 `math` 包中的函数。）


详细解答：[切片](../练习/切片/切片.md)

### 3.5 映射

#### 3.5.1 映射的文法

#### 3.5.2 修改映射

#### 3.5.3 练习：映射

### 3.6 函数值

#### 3.6.1 函数的闭包

#### 3.6.3 练习：斐波纳契闭包

## 4 [方法和接口](https://tour.go-zh.org/methods)

### 4.1 方法

### 4.2 接口

## 5 [并发](https://tour.go-zh.org/concurrency)
