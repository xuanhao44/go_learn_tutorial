# go_learn_tutorial

> 一个 go 学习项目。

## docsify

使用了 [docsify](https://docsify.js.org/#/zh-cn/) 组织文档。创建的理由是看到了一份用 docsify 构建的极简的 Go 命令行手册，马上就被这份简洁所折服。

使用方法和优缺点依靠官方教程和搜索可以轻松的完成，这里不作赘述。下面简单记录一下定制化的部分(index.html)。

1. 添加了网站 favicon.ico
2. 添加了封面页、侧边栏(目录)、导航栏、404 页面
3. 加载了 [docsify-katex](https://upupming.site/docsify-katex/docs/)、[docsify-tabs](https://jhildenbiddle.github.io/docsify-tabs/#/?id=docsify-tabs)、[分页导航Pagination](https://docsify.js.org/#/zh-cn/plugins?id=pagination)、代码高亮、代码复制到剪贴板、全文搜索、PWA 离线查看、字数统计等 js 插件
4. 注意到了禁用 Emoji 解析的细节
5. ...
6. 尝试但未处理好直接运行代码的插件

这份写好的 index.html 是可以复用的。
