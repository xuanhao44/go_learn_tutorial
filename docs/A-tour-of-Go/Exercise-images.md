# 练习：图像

## 题目

还记得之前编写的 [练习：切片-图片生成器](A-tour-of-Go/Exercise-slices.md) 吗？我们再来编写另外一个，不过这次它将会返回一个 `image.Image` 的实现而非一个数据切片。

定义你自己的 `Image` 类型，实现 [必要的方法](https://go-zh.org/pkg/image/#Image) 并调用 `pic.ShowImage`。

`Bounds` 应当返回一个 `image.Rectangle` ，例如 `image.Rect(0, 0, w, h)`。

`ColorModel` 应当返回 `color.RGBAModel`。

`At` 应当返回一个颜色。上一个图片生成器的值 `v` 对应于此次的 `color.RGBA{v, v, 255, 255}`。

```go
package main

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
```

## 解答



<!-- 网址或引用 -->
