# Go 源码编译过程

## 0 教程与简介

### 0.1 简介

本文从原理角度考察 Go 源码是如何被编译的。主要内容：

- 编译
- 编译相关的命令
- `import` 机制

### 0.2 参考

- [初探 Go 的编译命令执行过程 - 简书 (jianshu.com)](https://www.jianshu.com/p/35a4ec1b3067)

- [Go 命令教程 (hyper0x.github.io)](https://hyper0x.github.io/go_command_tutorial/#/)

  **写的太好了 我没什么好说的了**

## 1 `go help`

我想如果要系统的学习命令，大概第一个就应该是 `help`，这在任何地方都是通用的。例如 Git 查看所有命令的命令是 `git help`。

![help命令][help命令]

尽管你可以在网上查命令是如何输入的，但我认为 `help` 是最快的方法。

#### 1.1 使用命令

使用方法已经显示在命令行中：

```shell
Usage:

        go <command> [arguments]
```

简单翻译就是 `go 命令 对象`，如`go run hello.go`，命令是 `run`，对象是 `hello.go`。

#### 1.2 命令列表

下面介绍命令。目前 Go 最新版 1.8.3 里面基本命令只有以下的16个。

```shell
The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages
```

#### 1.3 查看命令更多信息

```shell
Use "go help <command>" for more information about a command.
```

如 `go help run`：

```
usage: go run [build flags] [-exec xprog] package [arguments...]

Run compiles and runs the named main Go package.
Typically the package is specified as a list of .go source files from a single
directory, but it may also be an import path, file system path, or pattern
matching a single known package, as in 'go run .' or 'go run my/cmd'.

If the package argument has a version suffix (like @latest or @v1.0.0),
"go run" builds the program in module-aware mode, ignoring the go.mod file in
the current directory or any parent directory, if there is one. This is useful
for running programs without affecting the dependencies of the main module.

If the package argument doesn't have a version suffix, "go run" may run in
module-aware mode or GOPATH mode, depending on the GO111MODULE environment
variable and the presence of a go.mod file. See 'go help modules' for details.
If module-aware mode is enabled, "go run" runs in the context of the main
module.

By default, 'go run' runs the compiled binary directly: 'a.out arguments...'.
If the -exec flag is given, 'go run' invokes the binary using xprog:
        'xprog a.out arguments...'.
If the -exec flag is not given, GOOS or GOARCH is different from the system
default, and a program named go_$GOOS_$GOARCH_exec can be found
on the current search path, 'go run' invokes the binary using that program,
for example 'go_js_wasm_exec a.out arguments...'. This allows execution of
cross-compiled programs when a simulator or other execution method is
available.

The exit status of Run is not the exit status of the compiled binary.

For more about build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build.
```

*好吧，我承认这个看起来还是比较困难的。知道就行了。*

## 2 `go env`

命令 `go env` 用于打印Go语言的环境信息。比较好理解的是 `GOROOT`，`GOPATH`，`GO111MODULE`，`GOPROXY` 等。在不清楚某项环境变量是如何设置之前可以使用该命令来查看。

<!-- 图片 -->

[help命令]:../_images/help命令.png

[help命令]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/help%E5%91%BD%E4%BB%A4.png
