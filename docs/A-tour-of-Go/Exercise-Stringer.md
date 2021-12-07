# 练习：Stringer

## 题目

通过让 `IPAddr` 类型实现 `fmt.Stringer` 来打印点号分隔的地址。

例如，`IPAddr{1, 2, 3, 4}` 应当打印为 `"1.2.3.4"`。

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

## 解答

### 分析源码

```go
type IPAddr [4]byte
```

定义了名为 `IPAddr` 的类型，是拥有 4 个字符(byte)的数组。

```go
hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
```

定义了从 `string` 类型到 `IPAddr` 的映射，用的是较为简略的写法。

```go
for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
```

采用了 Range 来遍历映射。`name` 为当前 `host` 元素的下标，`ip` 为 `host` 元素的副本。

### 最终代码

```go
package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
     return fmt.Sprintf("%v.%v.%v.%v",i[0], i[1], i[2], i[3])
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```



<!-- 网址或引用 -->
