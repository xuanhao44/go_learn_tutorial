# 练习：导入Go语言远程包

## 0 教程与简介

### 0.1 简介

导入包，抽象的说是在使用另一个包的内容。具体一点的说是在使用另一个包的代码。那么如果要使用远程包，需要做的当然是将其下载到本地，并进行一定的配置。下面介绍如何导入远程包。

### 0.2 参考

- [手把手教你导入Go语言第三方库_Golang_脚本之家 (jb51.net)](https://www.jb51.net/article/218754.htm)
- [go get安装第三方包的前提条件和步骤 - Go语言中文网 - Golang中文社区 (studygolang.com)](https://studygolang.com/articles/5840)
- ...

### 0.3 前提

- [Go语言包的基本概念](Go语言包的基本概念.md)
- [GOPATH](GOPATH.md)

## 1 任务

在某份 Go 源代码中导入一个包：

```go
import "golang.org/x/tour/pic"
```

## 2 远程包的路径格式

Go 语言的代码被托管于 Github.com 网站，该网站是基于 Git 代码管理工具的，很多有名的项目都在该网站托管代码。其他类似的托管网站还有 code.google.com、bitbucket.org 等。

这些网站的项目包路径都有一个共同的标准。

```html
github.com/golang/go
```

这个路径共由 3 个部分组成：

- `github.com` 网站域名：表示代码托管的网站，类似于电子邮件 @ 后面的服务器地址。
- `golang` 作者或机构：表明这个项目的归属，一般为网站的用户名，如果需要找到这个作者下的所有项目，可以直接在网站上通过搜索“域名/作者”进行查看。这部分类似于电子邮件 @ 前面的部分。
- `go` 项目名：每个网站下的作者或机构可能会同时拥有很多的项目，图中标示的部分表示项目名称。

## 3 解决步骤

有两种方法：分为**命令行自动安装**和**手动下载然后安装**。

### 3.1 命令行自动安装

```shell
go get golang.org/x/tour/pic
```

需要注意的是有前提要求：

1. 已经设置好环境变量 `GOPATH`
2. 已经安装 Git

### 3.2 手动下载安装

1. 到第三方类包所在的网址，手动下载源码(或源码压缩包)

2. 解压到 `gopath/src` 里面的路径，相关路径需要自己手动设置

   ![手动设置路径][手动设置路径]

3. 然后执行 `go install golang.org/x/tour/pic` 安装这个包

### 3.3 代理

实话说这个不应该算导入包的步骤之一，但是我们有着特殊的原因。上面的两个步骤基本上都是需要代理的。

比如，如果不开启代理则无法使用 `go get` 安装包。

![go_get_超时][go_get_超时]

比如，想要手动下载源码压缩包至少要能访问第三方类包所在的网址。

![pic包网址][pic包网址]

<!-- 图片 -->

[手动设置路径]:.assets/手动设置路径.png
[手动设置路径]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/%E6%89%8B%E5%8A%A8%E8%AE%BE%E7%BD%AE%E8%B7%AF%E5%BE%84.png

[go_get_超时]:.assets/go_get_超时.png

[go_get_超时]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go_get_%E8%B6%85%E6%97%B6.png

[pic包网址]:.assets/pic包网址.png
[pic包网址]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/pic%E5%8C%85%E7%BD%91%E5%9D%80.png
