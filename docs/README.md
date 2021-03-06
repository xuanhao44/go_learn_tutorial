# Go 学习笔记

> 在这里记录自己学 Go 的历程，长期更新。

## 参考

暂且不提 Go 的文档、手册以及语言规范。

1. [Go 语言之旅 (go-zh.org)](https://tour.go-zh.org/list)
2. [Go语言入门教程，Golang入门教程（非常详细） (biancheng.net)](http://c.biancheng.net/golang/)
3. [《Go 编程基础》(github.com)](https://github.com/unknwon/go-fundamental-programming)
4. [《The Way to Go》中文译本《Go 入门指南》 (github.com)](https://github.com/Unknwon/the-way-to-go_ZH_CN)
5. [《Go Web 编程》 (github.com)](https://github.com/astaxie/build-web-application-with-golang)
6. ...

## 学习顺序

### 基础

1. Go 本地安装和环境配置：[VSCode-Go](Go-local-environment-setting/VSCode-Go.md)
2. 入门：[Go 语言之旅](A-tour-of-Go/README.md)
3. 学习的同时可以看 Go 细节闲话的基础部分：

   - [Go语言包的基本概念](Go-details/Packages.md)

   - [GOPATH](Go-details/GOPATH.md)

   - [练习：Go语言自定义包_GOPATH](Go-details/Define-import-packages-byGOPATH.md)

   - [Go源码编译过程](Go-details/Go-compile.md)

   - [练习：导入Go语言远程包](Go-details/Import-remote-pakcages.md)

   - [Go Modules](Go-details/Go-Modules.md)
4. 还有一些其他的基础知识：
   - [OS 基础](other/OS.md)

### 网络编程

由于要学网络相关，故学[《Go Web 编程》](https://github.com/astaxie/build-web-application-with-golang)，此书安装 Go 和 Go 基础的部分可以拿来复习和查缺补漏，有了前面的基础理解的会很快。

## 流水账心得

按照学习的阶段写心得。

### 0

进度：0

我也许可以选择找一本讲解完整的书籍来学习，但我也想要经历自己一点点收集资料、一点点尝试的过程。尽管我也知道这两者并不矛盾，但在看书的时候却总是让书去代替我思考，也不愿意尝试上面的练习和拓展的部分。可见实践和理论之间还是有一定距离的。在本地安装完 Go 以及配置好环境后，正在学习 Go 语言之旅以及其他内容。学完之后可能就会自己创建项目了。

### 1

进度：

1. 本地环境配置完成
2. 语言基本学完除了接口和并发的部分
3. 对包、`GOPATH`、导入包有了一定的了解

一直在学 Go 语言之旅，有些难受，比较劝退。有时候会想用这个来入门是不是个错误。让我第一次羡慕其他有着舒适入门流程的指导书籍，但转念一想，这又不是坏事。

我有想过如果我准备齐全指导书籍，收藏了众多有价值的讲解，像是万事俱备一样开始学习会不会更好。但实际经验告诉我这不是好事——这就像一个刚开始运动的人提前买好了各式各样的运动装备从而觉得自己为了运动付出了许多一样，为你提供了虚假的满足感。类似的例子还有"到我的收藏夹里吃灰去吧"的文章，收藏了就觉得自己开始学了，学了好多...我觉得这样是不好的。事实上，我的 Python 的学习状况就是这样。(逃)

到现在学习了一段时间后，再来审视学习的过程。Go 语言之旅不能满足初学者对 Go 的学习，因此不得不寻找其他教程辅助；在此期间也构建了自己的笔记。

现在再去看一些文章对 Go 学习资料的说明，就会有"哇，相见恨晚"的感觉。这和在学之前看说明的感觉完全不同。在不真正通过学习知道这个资料有多重要之前，看着别人的文章说这个很重要，我相信也是没有什么实感的。

至于构建自己的笔记，我的看法是这样的：

1. 不论形式，一定要有
2. 尽量自己写，不要 copy 太多别人的

有时候觉得别人的笔记很全，没必要自己写；费劲心思写了看起来很全的笔记，之后却看到了自己的上位替代。但不是自己的终究不是自己的，自己构建笔记的过程本身就蕴含着自己的思考。所以构建笔记这件事还是不能停。

## 文档组织

仓库文件的组织也是一个新的尝试。用了 [docsify](https://docsify.js.org/#/zh-cn/) 组织文档，比较方便。

~~为了避免耦合，每篇文章都会做到尽量独立。~~那是不可能的，不然为啥要用工具来组织。
