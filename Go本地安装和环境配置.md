# Go 本地安装环境配置

## 0 参考和简介

这两篇文章非常详细的讲解了 Go 的本地安装和环境配置，跟他自己说的一样，踩的坑也是真的多。但这位老哥两篇文章没有整合到一起，看着有点难受。这里记录一下我自己的安装过程。

- [在VsCode中搭建Go开发环境，手把手教你配置](https://blog.csdn.net/AdolphKevin/article/details/90274378?spm=1001.2014.3001.5501)
- [VSCode搭建Go开发环境](https://blog.csdn.net/AdolphKevin/article/details/105480530)

Go 的这些文件的组织确实还是很重要的，只能说我在学的同时不断地学习如何改进文件组织吧。但是这篇文章不会写更多的东西了。

## 1 Go 的安装

安装包下载地址，这三个都 *OK*。

- https://golang.org/dl/
- https://golang.google.cn/dl/
- https://studygolang.com/dl

*我这里是 Windows 安装，就不提 Linux 了。*

Windows 安装包名：**`go1.17.3.windows-amd64.msi`**

![windows安装包][windows安装包]

安装时**注意安装目录**，我安装在 **`c:\Program Files\Go`** 目录下。

在控制台输入

```shell
$ go version
```

如果显示了版本号，即代表安装成功。

![go_version][go_version]

在设置环境之前，可以先在**任何地方**创建一个 go 文件，比如 `hello.go`，内容为：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

在命令行输入：（右键 — "在 *Windows* 终端中打开"，默认的一般是 *PowerShell*）

```shell
$ go run hello.go
```

结果：

![go_run][go_run]

## 2 Go 环境变量的配置

要配置的环境变量有 `GOPATH `和 系统变量 `Path`。

据说 Go 在 1.6 版本后，这些环境变量都不需要我们手动配置了。安装的时候都会自动配置好。

事实上确实是这样，系统变量 `Path` 已经设置好了，不需要我们去调整。

![Path][Path]

但是 `GOPATH` 是需要调整的。因为 `GOPATH` 路径是我们的工作区，就是我们放代码的地方。（和上面 Go 的安装目录不是一个东西）

代码别乱放，本来这些语言的配置文件就乱的吓人，就别继续在自己的电脑里乱丢垃圾了。（笑）

例如我的就设为：`D:\OneDrive\Codefield\CODE_Go`

### 2.1 设置 `GOPATH` 路径

打开命令行输入：

```shell
$ go env -w GOPATH=我们自己的工作区路径
```

但结果报了错：

```shell
does not override conflicting OS environment variable
```

不知道是什么原因，哎。

![go_env][go_env]

那只能手动到环境变量的地方设置 `GOPATH`。怎么找到设置环境变量？我的评价是：善用搜索。

![搜索环境变量][搜索环境变量]

（这是比尔盖茨给你的搜索按钮.jpg）

系统环境变量和账户环境变量选哪个都 *OK*，但选择账户的环境变量直接一些。

![账户还是帐户？微软，我的文盲][账户还是帐户]

一看原来已经有了。（原本是被设置在了用户文件夹里的，图示是已经修改之后的 `GOPATH`）

赶紧把 `GOPATH` 修改为自己的工作目录吧！

![GOPATH][GOPATH]

注：如果要变更 `GOPATH` 的路径，建议把原来 `GOPATH` 里的 `pkg`、`bin`、`src` 三个文件夹也复制过去，不然到时候开发时又要重新配置依赖包，不知情的可能还认为出 *BUG* 了呢。

### 2.2 打开 GoMOD，配置代理

打开控制台输入命令。

```shell
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

### 2.3 在 VS Code 中安装 Go 插件

进入 *Extensions* 后直接搜索 go，即可安装。

![Go拓展][Go拓展]

### 2.4 创建第一个项目

在我们的工作目录下，创建一个 `src` 文件夹。再在 `src` 文件夹下创建一个项目文件夹 `hello` ，项目文件夹 `hello` 下创建一个 `hello.go` 文件。

用记事本打开，输入以下代码并保存：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

（这确实是你的第一个 Go 的 HelloWorld，但是到能把它在 VS Code 里运行起来还有一段距离，不过也可以先使用 `go run` 命令编译源码，在命令行里查看结果。

### 2.5 安装 Go 的依赖包

用 VSCode 打开 你的 `GOPATH` 文件夹。（文件夹右键 — "通过 Code 打开"）

![通过_Code_打开][通过_Code_打开]

安装了 Go 插件后的 VS Code 现在打开 go 文件后，会自动安装我们自己的必要的环境依赖。

（没自动的话就点点点右下角的 `Install All`）

又是几百个 *Failed*，看着真的很糟心。

![问题百出][问题百出]

但也不是全部都失败了。可以看到 `GOPATH` 那里还是下了一部分的，在 `GOPATH` 里创建了 `bin`、`pkg` 文件夹。

![pkg][pkg]

*But*，没想到等一会就安装好了！也没有其他问题。惊喜，感觉就像等了一辈子。

![安装依赖完毕][安装依赖完毕]

### 2.6 Go Modules 的使用

Go Mod 是一个 Go 语言依赖库管理器。我们要把它安置在 hello 项目里。

进入我们的 `hello` 文件夹，并且执行 `go mod init` 即可。

![go_mod_init][go_mod_init]

可以看到多出了一个 `go.mod`文件（*PS*：我没看见博主说的 `go.sum` 文件）

![go.mod1][go.mod1]

`go.mod` 文件记录我们依赖库以及版本号。

用 VS Code 打开 `go.mod`：但现在可以看到我们现在什么都库都还没有安装，空空如也。

![go_mod2][go_mod2]

此时我们这个 `hello` 项目，就采用了 `go mod` 进行库依赖管理。

使用起来也非常简单，常用命令就一个 `go mod tidy`，通俗来说就是将当前的库源码文件所依赖的包，全部安装并记录下来，多的包就删掉，少了的就自动补上，以后编写完代码后执行 `go mod tidy` 即可。

更多的使用方式可以执行 `go help mod` 进行查看。

![go_help_mod][go_help_mod]

### 2.7 VS Code 的 `launch.json` 配置以及 `setting.json` 配置

`launch.json`

```javascript
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "LaunchGo",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "remotePath": "",
            "port": 5546,
            "host": "127.0.0.1",
            "program": "${fileDirname}",
            "env": {
                "GOPATH": "D:\\OneDrive\\Codefield\\CODE_Go",
                "GOROOT": "C:\\Program Files\\Go"
            },
            "args": [],
            //"showLog": true
        }
    ]
```

`setting.json`

```javascript
{
    "editor.wordWrap": "on",
    "go.useLanguageServer": false,
    "editor.minimap.renderCharacters": false,
    "editor.minimap.enabled": false,
    "terminal.external.osxExec": "iTerm.app",
    "go.docsTool": "gogetdoc",
    "go.testFlags": ["-v","-count=1"],
    "go.buildTags": "",
    "go.buildFlags": [],
    "go.lintFlags": [],
    "go.vetFlags": [],
    "go.coverOnSave": false,
    "go.useCodeSnippetsOnFunctionSuggest": false,
    "go.formatTool": "goreturns",
    "go.gocodeAutoBuild": false,
    "go.goroot": "C:\\Program Files\\Go",
    "go.gopath": "D:\\OneDrive\\Codefield\\CODE_Go",
    "go.autocompleteUnimportedPackages": true,
    "go.formatOnSave": true,
    "window.zoomLevel": 0,
    "debug.console.fontSize": 16,
    "debug.console.lineHeight": 30
}
```

出现红色异常的地方直接删除即可，比如删一两个逗号之类的。

记得修改 `GOPATH`、`GOROOT`、`go.goroot`、`go.gopath`，改成自己的地址。从 Windows 那里复制来的地址的右斜线记得改成双右斜线，别复制完就走了。

配置好后，直接按 <kbd>F5</kbd> 就能进入调试状态。

## 3 实验调试

结束，爽按 <kbd>F5</kbd>，开始调试，看着控制台的输出很开心。

![控制台输出][控制台输出]

### 3.1 创建第二个项目

开一个新项目，编译运行：`packages.go`

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}
```

流程：

1. 在 `src` 文件夹下创建 `try` 文件夹。

2. 在 `try` 文件夹下创建 `packages.go` 文件夹。

3. 复制 `.vscode` 文件夹。

4. 在 `try` 文件夹下执行 `go mod init` 命令。

5. 实验调试。

结果没问题，非常顺利。目前来说，以后创建项目的大体流程都是一样的。

### 3.2 项目也可以不放在 `src` 中

项目不放在 `src` 文件夹中也是 *OK* 的。随便在哪里按照上面 3.1 除去第一步的流程创建一个项目，再把 `launch.json` 的 `port` 和 `host` 那两行删除就能正常的编译运行了。

## 4 总结

充其量算一个不求甚解的流程介绍，照着做下来也不是很懂为什么。据我目前的经验来说这是不大好的。我的评价是：最好继续**研究原理**。

现在这个教程只是非常非常简单的傻瓜设置，还有很多的局限性。

- 一个项目暂时还只能放一个源代码
- 包的引用形同摆设
- ...

<!-- 图片 -->

[windows安装包]:.assets/windows安装包.png

[windows安装包]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/windows%E5%AE%89%E8%A3%85%E5%8C%85.png

[go_version]:.assets/go_version.png
[go_version]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go_version.png
[go_run]:.assets/go_run.png
[go_run]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go_run.png

[Path]:.assets/Path.png
[Path]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/Path.png

[go_env]:.assets/go_env.png
[go_env]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go_env.png
[账户还是帐户]:.assets/账户还是帐户.png
[账户还是帐户]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/%E8%B4%A6%E6%88%B7%E8%BF%98%E6%98%AF%E5%B8%90%E6%88%B7.png
[搜索环境变量]:.assets/搜索环境变量.png
[搜索环境变量]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/%E6%90%9C%E7%B4%A2%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F.png

[GOPATH]:.assets/GOPATH.png
[GOPATH]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/GOPATH.png

[Go拓展]:.assets/Go拓展.png
[Go拓展]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/Go%E6%8B%93%E5%B1%95.png

[通过_Code_打开]:.assets/通过_Code_打开.png
[通过_Code_打开]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/%E9%80%9A%E8%BF%87_Code_%E6%89%93%E5%BC%80.png

[问题百出]:.assets/问题百出.png
[问题百出]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/%E9%97%AE%E9%A2%98%E7%99%BE%E5%87%BA.png

[pkg]:.assets/pkg.png
[pkg]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/pkg.png

[安装依赖完毕]:.assets/安装依赖完毕.png
[安装依赖完毕]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/%E5%AE%89%E8%A3%85%E4%BE%9D%E8%B5%96%E5%AE%8C%E6%AF%95.png

[go_mod_init]:.assets/go_mod_init.png
[go_mod_init]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go_mod_init.png

[go.mod1]:.assets/go.mod1.png
[go.mod1]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go.mod1.png

[go_mod2]:.assets/go_mod2.png
[go_mod2]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go_mod2.png

[go_help_mod]:.assets/go_help_mod.png
[go_help_mod]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/go_help_mod.png

[控制台输出]:.assets/控制台输出.png

[控制台输出]:https://typora-1304621073.cos.ap-guangzhou.myqcloud.com/typora/%E6%8E%A7%E5%88%B6%E5%8F%B0%E8%BE%93%E5%87%BA.png
