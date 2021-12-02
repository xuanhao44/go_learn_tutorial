# 练习：Go语言自定义包_GOPATH

**本文地址**：[xuanhao44/Go_Learn](https://github.com/xuanhao44/Go_Learn)

## 0 教程与简介

### 0.1 简介

包是 Go 语言中代码组成和代码编译的主要方式。关于包的基本信息我们已经在前面介绍过了，本节我们主要来介绍一下如何自定义一个包并使用它。

本文的讲解模式是**用例子来讲规则**。

### 0.2 参考

- [Go语言自定义包 (biancheng.net)](http://c.biancheng.net/view/5123.html)
- [GoLand中报错package xxx is not in GOROOT_THEGREATHXY的博客-CSDN博客](https://blog.csdn.net/THEGREATHXY/article/details/109337283)
- [Go语言import导入包——在代码中使用其他的代码 (biancheng.net)](http://c.biancheng.net/view/91.html)
- ...

### 0.3 前提

- [Go语言包的基本概念](../Go语言包的基本概念/Go语言包的基本概念.md)
- [GOPATH](../GOPATH/GOPATH.md)

实际运用这两篇中的知识，当然要先看。

### 0.4 提示

本文使用的是 `GOPATH`。

请按照 [GOPATH](../GOPATH/GOPATH.md) 一文中的说明修改模式。

```shell
go env -w GO111MODULE=off
```

### 1 任务

创建一个名为 `demo` 的自定义包，并在 `main` 包中使用自定义包 `demo` 中的方法。

`demo.go` 文件的代码如下所示：

```go
package demo

import "fmt"

func PrintStr() {
    fmt.Println("xuanhao44")
}
```

### 2 解决步骤以及讲解

#### 2.1 `demo.go`

从 `package demo` 可知，`demo.go` 属于 `demo` 包，而这个 Go 源码只是与其同名而已。

>  包名为 `main` 的包为**应用程序的入口包**，编译不包含 `main` 包的源码文件时不会得到可执行文件。

所以编译这个 `demo` 包并不会得到可执行文件。如图所示：

![not_main][not_main]

> 一般包的名称就是其源文件所在目录的名称，虽然Go语言没有强制要求包名必须和其所在的目录名同名，但还是建议包名和所在目录同名，这样结构更清晰。
>
> 我们创建的自定义的包需要将其放在 `GOPATH` 的 `src` 目录下。如果项目的目录不在 `GOPATH` 环境变量中，则需要把项目移到 `GOPATH` 所在的目录中，或者将项目所在的目录设置到 `GOPATH` 环境变量中，否则无法完成编译。

所以要在 `src` 文件夹下创建 `demo` 文件夹，将 `demo.go` 放在 `demo` 的文件夹里。

![src_demo_demo.go][src_demo_demo.go]

> **标准包的导入只能使用全路径导入**。
>
> 包的绝对路径就是 `GOROOT/src/`或`GOPATH/src/` 后面包的存放路径。
>
> **标准的 Go 语言代码库**中包含了大量的包，并且在安装 Go 的时候多数会自动安装到系统中。
>
> 我们可以在 `$GOROOT/src/pkg` 目录中查看这些包。

从 `import "fmt"` 可知以**单行全路径**导入的方式导入了 `fmt` 包。显然这个包不在 `GOPATH/src` 中，在 `GOROOT/src` 去找。

> 包中的函数名第一个字母要**大写**，否则无法在外部调用。

所以函数设置为 `func PrintStr()`。

### 2.2 `main.go`

`main.go` 文件的代码如下所示，按照相同的办法创建 `main` 包。

```go
package main

import "demo"

func main() {
    demo.PrintStr()
}
```

使用的是标准引用格式，即使用 `包名.函数名` 的方式。

### 2.3 运行

运行结果如下所示：

```shell
go run main.go
xuanhao44
```

注：如果未修改 Go Mod 模式，则无法编译。错误输出如下：

![not_in_GOROOT][not_in_GOROOT]

*记得改！*

<!-- 图片 -->

[not_main]:.assets/not_main.png
[not_main]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/not_main.png

[src_demo_demo.go]:.assets/src_demo_demo.go.png
[src_demo_demo.go]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/src_demo_demo.go.png

[not_in_GOROOT]:.assets/not_in_GOROOT.png
[not_in_GOROOT]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/not_in_GOROOT.png