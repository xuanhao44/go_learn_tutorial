# Go Modules

## 0 教程与简介

### 0.1 简介

- -

### 0.2 参考

- [关于Go Modules，看这一篇文章就够了 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/105556877)
- [使用 Go 模块 - go.dev](https://go.dev/blog/using-go-modules)
- [Go Modules 终极入门_Go中国-CSDN博客](https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/104568182/)
- ...

## 1 一句话介绍 Go Modules

`GOPATH` 不好用，之后出现过各种包管理工具，而 Go Modules 是 Go 官方出品的，先进的包管理器。

## 2 使用 Go Modules 前提

1. 执行 `go env` 命令，查看 `GO111MODULE` 环境变量。这个环境变量来作为 Go Modules 的开关，其允许设置以下参数：

   - auto：只要项目包含了 go.mod 文件的话启用 Go Modules，目前在Go1.11至 Go1.14 中仍然是默认值。

   - on：启用 Go Modules，推荐设置，将会是未来版本中的默认值。

   - off：禁用 Go Modules，不推荐设置。

   如果需要对 `GO111MODULE` 的值进行变更，推荐通过`go env`命令进行设置：

   ```shell
   go env -w GO111MODULE=on
   ```

2. 设置 Go 模块代理（Go module proxy）

   ```shell
   go env -w GOPROXY=https://goproxy.cn,direct
   ```

## 3 实际创建项目

1. 创建名为 `try` 的项目。

2. 进入项目文件夹，执行 `go mod init`，项目文件夹中会生成一个 go.mod 文件（空空如也）。

   (看起来不需要 go 源码也可以先 init)

   ```.mod
   module try
   
   go 1.17
   ```

3. 在项目文件夹中添加 go 源码：`main.go`

   显然 `golang.org/x/tour/pic` 是我们需要使用的包。

   ```go
   package main
   
   import (
   	"golang.org/x/tour/pic"
   	"math/rand"
   	"time"
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
   	pic.Show(Pic(3, 4))
   }
   ```

4. 在项目根目录执行 `go get golang.org/x/tour/pic` 命令。

   ![go_modules_go_get][go_modules_go_get]

5. 查看 go.mod 文件，基本内容如下：

   ```.mod
   module try
   
   go 1.17
   
   require golang.org/x/tour v0.1.0 // indirect
   ```
   
6. 查看 go.sum 文件，基本内容如下：
   
   ```.sum
   golang.org/x/tour v0.1.0 h1:OWzbINRoGf1wwBhKdFDpYwM88NM0d1SL/Nj6PagS6YE=
   golang.org/x/tour v0.1.0/go.mod h1:DUZC6G8mR1AXgXy73r8qt/G5RsefKIlSj6jBMc8b9Wc=
   ```

### 最后

原理不敢在这里讲啊。建议还是看参考的好。以后接着更新。

<!-- 图片 -->

[go_modules_go_get]:../_images/go_modules_go_get.png
