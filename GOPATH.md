# GOPATH

**本文地址**：[xuanhao44/Go_Learn](https://github.com/xuanhao44/Go_Learn)

## 0 教程与简介

### 0.1 简介

本文简单介绍 `GOPATH`，未涉及命令行的编译命令。

### 0.2 参考

- [Go语言GOPATH详解（Go语言工作目录） (biancheng.net)](http://c.biancheng.net/view/88.html)
- [Go语言GOPATH是什么 - 艾里_Simple - 博客园](https://www.cnblogs.com/ailiailan/p/13454139.html)
- [GoLand中报错package xxx is not in GOROOT_THEGREATHXY的博客-CSDN博客](https://blog.csdn.net/THEGREATHXY/article/details/109337283)
- ...

## 1 `GOPATH` 定义

`GOPATH` 是 Go 语言中使用的一个环境变量，它使用绝对路径提供项目的**工作目录（**也称为**工作区）**。

> 工作目录是一个工程开发的相对参考目录，好比当你要在公司编写一套服务器代码，你的工位所包含的桌面、计算机及椅子就是你的工作区。

## 2 使用 `GOPATH` 的工程结构

`GOPATH` 目录一般为：

```shell
  --bin      # 存放编译后的可执行文件
  --pkg      # 依赖包编译后的*.a文件
  --src      # 存放源码文件，以代码包为组织形式
```

- 在 `GOPATH` 指定的工作目录下，代码总是会保存在 `$GOPATH/src` 目录下。
- 在工程经过编译后，会将产生的二进制可执行文件放在 `$GOPATH/bin` 目录下。
- 生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下。

如果需要将整个源码添加到版本管理工具（*VersionControlSystem*，*VCS*）中时，只需要添加 `GOPATH/src` 目录的源码即可。`bin` 和 `pkg` 目录的内容都可以由 `src` 目录生成。

## 3 命令行设置 GOPATH

选择一个目录，在目录中的命令行中执行下面的指令：

```shell
export GOPATH=`pwd`
```

使用 `export` 指令可以将当前目录的值设置到环境变量 `GOPATH` 中。

该指令中的 `pwd` 将输出当前的目录，使用反引号将 `pwd` 指令括起来表示命令行替换，也就是说，使用 `pwd` 将获得 `pwd` 返回的当前目录的值。

*一个项目在开始之前，首先就要指定 `GOPATH`。*

## 4 在多项目工程中使用GOPATH

上面描述的 `GOPATH` 其实是通过修改系统全局的环境变量来实现的。然而，这种设置全局 `GOPATH` 的方法可能会导致当前项目错误引用了其他目录的 Go 源码文件从而造成编译输出错误的版本或编译报出一些无法理解的错误提示。

比如说，将某项目代码保存在 `/home/davy/projectA` 目录下，将该目录设置为 `GOPATH`。

随着开发进行，需要再次获取一份工程项目的源码，此时源码保存在 `/home/davy/projectB` 目录下，如果此时需要编译 `projectB` 目录的项目，但开发者忘记设置 GOPATH 而直接使用命令行编译，则当前的 `GOPATH` 指向的是 `/home/davy/projectA` 目录，而不是开发者编译时期望的 `projectB` 目录。编译完成后，开发者就会将错误的工程版本发布到外网。

因此，建议大家无论是使用命令行或者使用集成开发环境编译 Go 源码时，GOPATH 跟随项目设定。在 Jetbrains 公司的 GoLand 集成开发环境（IDE）中的 `GOPATH` 设置分为全局 `GOPATH` 和项目 `GOPATH`，如下图所示。

![Goland][Goland]

Global `GOPATH` 代表全局 `GOPATH`，一般来源于系统环境变量中的 `GOPATH`；

Project `GOPATH` 代表项目所使用的 `GOPATH`，该设置会被保存在工作目录的 **.idea** 目录下，不会被设置到环境变量的 `GOPATH` 中，但会在编译时使用到这个目录。

建议在开发时填写项目 `GOPATH`，**每一个项目尽量只设置一个 `GOPATH`，不使用多个 `GOPATH` 和全局的 `GOPATH`**。

## 5 GoMod 和 `GOPATH` 是冲突的

打开 Goland 创建的项目的目录，发现它是采用的 Go Mod。虽然现在没讲 Go Mod，但是必须要知道 Go Mod 是和 `GOPATH` 冲突的，二者只能选其一，不然编译会出现问题。

![Goland_gomod][Goland_gomod]

如果要使用 `GOPATH` 包管理方案，就在命令行内输入 ：

```shell
go env -w GO111MODULE=off
```

反之要使用 Go Mod 包管理方案则 `on`。

## 5 总结

知道了 `GOPATH` 的定义：工作空间。使用 `GOPATH` 的项目有固定的工程结构。有全局和项目 `GOPATH` 之分。每一个项目的 `GOPATH` 都必须要设定好。GoMod 和 `GOPATH` 是冲突的。

但是遗憾的是到目前也只是在了解粗浅的概念。仍然不能自己随心所欲的创建项目，限制真的相当的多。

<!-- 图片 -->

[Goland]:.assets/Goland.png
[Goland]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/Goland.png
[Goland_gomod]:.assets/Goland_gomod.png
[Goland_gomod]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/Goland_gomod.png
