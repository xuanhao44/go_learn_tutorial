# README

## 简介

Go 语言中文教程：[Go 语言之旅 (go-zh.org)](https://tour.go-zh.org/list)，可以在线学。本文基本把教程全文搬运过来了，所以会比较长。但是细节上略有不同，添加了自己的一些笔记，增添了目录结构，以供自己查阅。

- [包、变量和函数](A-tour-of-Go/Packages-variables-and-functions.md)：学习 Go 程序的基本结构。

- [流程控制语句：for、if、else、switch 和 defer](A-tour-of-Go/Flow-control-statements.md)：学习如何使用条件、循环、分支和推迟语句来控制代码的流程。

- [更多类型：struct、slice 和映射](A-tour-of-Go/Structs-slices-and-maps.md)：学习如何基于现有类型定义新的类型：本节课涵盖了结构体、数组、切片和映射。

- [方法和接口](A-tour-of-Go/Methods-and-interfaces.md)：学习如何为类型定义方法；如何定义接口；以及如何将所有内容贯通起来。本节课包含了方法和接口，可以用这种构造来定义对象及其行为。

- [并发](A-tour-of-Go/Concurrency.md)：作为语言的核心部分，Go 提供了并发的特性。这一部分概览了 goroutine 和 channel，以及如何使用它们来实现不同的并发模式。Go 将并发结构作为核心语言的一部分提供。本节课程通过一些示例介绍并展示了它们的用法。

Go 学习之旅中穿插着练习，将其抽出单独处理。

- [循环与函数](A-tour-of-Go/Exercise-loops-and-Functions.md)
- [切片](A-tour-of-Go/Exercise-slices.md)
- [映射](A-tour-of-Go/Exercise-maps.md)
- [斐波纳契闭包](A-tour-of-Go/Exercise-fibonaci-closure.md)
- [练习参考答案](A-tour-of-Go/Exercise-answer.md)

## 评价

不幸的是，这份中文教程学起来挺费劲的。学到方法和接口的时候难受极了。更让人失望的是，这份教程即使是在一个初学者看来也是相当简略的。照着教程学是没有问题的(毕竟你只是运行一下他的代码！)，遇到自己动手的练习就问题百出了。所以还是有必要自己去查询资料。

这里先把这份教程的**《接下来去哪》**放出来。再次重申，**不要在一棵树上吊死**。

> 你可以从[安装 Go](https://go-zh.org/doc/install/) 开始。
>
> 一旦安装了 Go，Go [文档](https://go-zh.org/doc/)是一个极好的 应当继续阅读的内容。 它包含了参考、指南、视频等等更多资料。
>
> 了解如何组织 Go 代码并在其上工作，参阅[此视频](https://www.youtube.com/watch?v=XCsL89YtqCs)，或者阅读[如何编写 Go 代码](https://go-zh.org/doc/code.html)。
>
> 如果你需要标准库方面的帮助，请参考[包手册](https://go-zh.org/pkg/)。如果是语言本身的帮助，阅读[语言规范](https://go-zh.org/ref/spec)是件令人愉快的事情。
>
> 进一步探索 Go 的并发模型，参阅 [Go 并发模型](https://www.youtube.com/watch?v=f6kdp27TYZs)([幻灯片](https://talks.go-zh.org/2012/concurrency.slide))以及[深入 Go 并发模型](https://www.youtube.com/watch?v=QDDwwePbDtw)([幻灯片](https://talks.go-zh.org/2013/advconc.slide))并阅读[通过通信共享内存](https://go-zh.org/doc/codewalk/sharemem/)的代码之旅。
>
> 想要开始编写 Web 应用，请参阅[一个简单的编程环境](https://vimeo.com/53221558)([幻灯片](https://talks.go-zh.org/2012/simple.slide))并阅读[编写 Web 应用](https://go-zh.org/doc/articles/wiki/)的指南。
>
> [函数：Go 中的一等公民](https://go-zh.org/doc/codewalk/functions/)展示了有趣的函数类型。
>
> [Go 博客](https://blog.go-zh.org/)有着众多关于 Go 的文章和信息。
>
> [mikespook 的博客](https://www.mikespook.com/tag/golang/)中有大量中文的关于 Go 的文章和翻译。
>
> 开源电子书 [Go Web 编程](https://github.com/astaxie/build-web-application-with-golang)和 [Go 入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN)能够帮助你更加深入的了解和学习 Go 语言。
>
> 访问 [go-zh.org](https://go-zh.org/) 了解更多内容。
