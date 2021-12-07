# 练习：rot13Reader

## 题目

有种常见的模式是一个 [`io.Reader`](https://go-zh.org/pkg/io/#Reader) 包装另一个 `io.Reader`，然后通过某种方式修改其数据流。

例如，[`gzip.NewReader`](https://go-zh.org/pkg/compress/gzip/#NewReader) 函数接受一个 `io.Reader`（已压缩的数据流）并返回一个同样实现了 `io.Reader` 的 `*gzip.Reader`（解压后的数据流）。

编写一个实现了 `io.Reader` 并从另一个 `io.Reader` 中读取数据的 `rot13Reader`，通过应用 [rot13](http://en.wikipedia.org/wiki/ROT13) 代换密码对数据流进行修改。

`rot13Reader` 类型已经提供。实现 `Read` 方法以满足 `io.Reader`。

```go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

## 解答



<!-- 网址或引用 -->
