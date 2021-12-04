# 练习：修改映射

## 题目

实现 `WordCount`。它应当返回一个映射，其中包含字符串 `s` 中每个“单词”的个数。函数 `wc.Test` 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 [strings.Fields](https://go-zh.org/pkg/strings/#Fields) 很有帮助。

## 解答

前提：[练习：导入Go语言远程包](Go-details/Import-remote-pakcages.md)

### strings.Fields

功能大概是将一个字符串按照空格分解为多个子串并返回子串的数组(并不是多值返回)。

具体得可以去看 [strings.Fields](https://go-zh.org/pkg/strings/#Fields)

示例 1：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
```

输出结果：

```shell
Fields are: ["foo" "bar" "baz"]
```

示例 2：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  foo bar  baz   "
	index := strings.Fields(s)
	fmt.Println(len(index))
}
```

输出结果：

```shell
3
```

### 本地测试

```go
package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	// 获取子串数组 赋值到 index
	index := strings.Fields(s)
	// 映射
	m := make(map[string]int)
	// 空间换时间
	for i := range index {
		m[index[i]]++
	}
	return m
}

func main() {
	s := "I am learning Go!"
	fmt.Println(WordCount(s))
}
```

输出结果：

```go
map[Go!:1 I:1 am:1 learning:1]
```

### 最终代码


```go
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	var m = make(map[string]int)
	index := strings.Fields(s)
	m[s] = len(index)
	return m
}

func main() {
	wc.Test(WordCount)
}
```

输出结果：

```shell
PASS
 f("I am learning Go!") = 
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") = 
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") = 
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") = 
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
```

<!-- 网址或引用 -->
