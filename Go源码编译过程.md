# Go 源码编译过程

**本文地址**：[xuanhao44/Go_Learn](https://github.com/xuanhao44/Go_Learn)

## 0 教程与简介

### 0.1 简介

本文从原理角度考察 Go 源码是如何被编译的。主要内容：

- 编译
- 编译相关的命令
- `import` 机制

### 0.2 参考

- [初探 Go 的编译命令执行过程 - 简书 (jianshu.com)](https://www.jianshu.com/p/35a4ec1b3067)

## 1 源代码是如何编译成可执行文件的

## 2 Go 编译相关命令

`go install`

`go install` 用于生成可执行文件。

`go install` 将编译的中间文件放在 `GOPATH` 的 `pkg` 目录下，以及固定地将编译结果放在 `GOPATH` 的 `bin` 目录下。如果当前目录是非主包，则`go install` 直接把编译结果安装到 `$GOPATH/pkg`，并缓存，如果包未做更改，下次编译则直接使用缓存。 

## 3 `import` 的机制

<!-- 图片 -->