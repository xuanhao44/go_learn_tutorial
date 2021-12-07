# 练习参考答案

*目前尚不完整*

<!-- tabs:start -->

#### **Exercise: Loops and Functions #43**

```go
/* Exercise: Loops and Functions #43 */
package main
 
import (
    "fmt"
    "math"
)
 
func Sqrt(x float64) float64 {
    z := float64(2.)
    s := float64(0)
    for {
        z = z - (z*z - x)/(2*z)
        if math.Abs(s-z) < 1e-15 {
            break
        }
        s = z
    }
    return s
}
 
func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
```

#### **Exercise: Maps #44**

```go
/* Exercise: Maps #44 */
package main
 
import (
    "tour/wc"
    "strings"
)
 
func WordCount(s string) map[string]int {
    ss := strings.Fields(s)
    num := len(ss)
    ret := make(map[string]int)
    for i := 0; i < num; i++ {
        (ret[ss[i]])++
    }
    return ret
}
 
func main() {
    wc.Test(WordCount)
}
```

#### **Exercise: Slices #45**

```go
/* Exercise: Slices #45 */
package main
 
import "tour/pic"
 
func Pic(dx, dy int) [][]uint8 {
    ret := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        ret[i] = make([]uint8, dx)
        for j := 0; j < dx; j++ {
            ret[i][j] = uint8(i^j+(i+j)/2)
        }
    }
    return ret
}
 
func main() {
    pic.Show(Pic)
}
```

#### **Exercise: Fibonacci closure #46**

```go
/* Exercise: Fibonacci closure #46 */
package main
 
import "fmt"
 
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    x := 0
    y := 1
    return func() int {
        x,y = y,x+y
        return x
    }
}
 
func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
```

#### **Advanced Exercise: Complex cube roots #47**

```go
/* Advanced Exercise: Complex cube roots #47 */
package main
 
import (
    "fmt"
    "math/cmplx"
)
 
func Cbrt(x complex128) complex128 {
    z := complex128(2)
    s := complex128(0)
    for {
        z = z - (cmplx.Pow(z,3) - x)/(3 * (z * z))
        if cmplx.Abs(s-z) < 1e-17 {
            break
        }
        s = z
    }
    return z
}
 
func main() {
    fmt.Println(Cbrt(2))
}
```

#### **Exercise: Errors #57**

```go
/* Exercise: Errors #57 */
package main
 
import (
    "fmt"
)
 
type ErrNegativeSqrt float64
 
func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negativ number: %g", float64(e))
}
 
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    }
    return 0, nil
}
 
func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
```

#### **Exercise: Images #58**

```go
/* Exercise: Images #58 */
package main
 
import (
    "image"
    "tour/pic"
    "image/color"
)
 
type Image struct{
    Width, Height int
    colr uint8    
}
 
func (r *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, r.Width, r.Height)
}
 
func (r *Image) ColorModel() color.Model {
    return color.RGBAModel
}
 
func (r *Image) At(x, y int) color.Color {
    return color.RGBA{r.colr+uint8(x), r.colr+uint8(y), 255, 255}
}
 
func main() {
    m := Image{100, 100, 128}
    pic.ShowImage(&m)
}
```

#### **Exercise: Rot13 Reader #59: 'You cracked the code!'**

```go
/* Exercise: Rot13 Reader #59: 'You cracked the code!' */
package main
 
import (
    "io"
    "os"
    "strings"
)
 
type rot13Reader struct {
    r io.Reader
}
 
func (rot *rot13Reader) Read(p []byte) (n int, err error) {
    n,err = rot.r.Read(p)
    for i := 0; i < len(p); i++ {
        if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
            p[i] += 13
        } else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
            p[i] -= 13
        }
    }
    return
}
 
func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
```

#### **Exercise: Equivalent Binary Trees #67**

```go
/* Exercise: Equivalent Binary Trees #67 */
package main
 
import (
    "tour/tree"
    "fmt"
)
 
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    _walk(t, ch)
    close(ch)
    
}
 
func _walk(t *tree.Tree, ch chan int) {
    if t != nil {
        _walk(t.Left, ch)
        ch <- t.Value
        _walk(t.Right, ch)
    }
}
 
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for i := range ch1 {
        if i != <- ch2 {
            return false
        }
    }
    return true    
}
 
func main() {
    //tree.New(2)
    ch := make(chan int)
    go Walk(tree.New(1), ch)
    for v := range ch {
        fmt.Print(v)
    }
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
```

#### **Exercise: Web Crawler #69**

```go
/* Exercise: Web Crawler #69 */
package main
 
import (
    "fmt"
)
 
type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}
 
var store map[string]bool
 
func Krawl(url string, fetcher Fetcher, Urls chan []string) {
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("found: %s %q\n", url, body)
    }
    Urls <- urls
}
 
func Crawl(url string, depth int, fetcher Fetcher) {
    Urls := make(chan []string)
    go Krawl(url, fetcher, Urls)
    band := 1
    store[url] = true // init for level 0 done
    for i := 0; i < depth; i++ {
        for band > 0 {
            band--
            next := <- Urls
            for _, url := range next {
                if _, done := store[url] ; !done {
                    store[url] = true
                    band++
                    go Krawl(url, fetcher, Urls)
                }
            }
        }
    }
    return
}
 
func main() {
    store = make(map[string]bool)
    Crawl("http://golang.org/", 4, fetcher)
}
 
 
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult
 
type fakeResult struct {
    body string
    urls     []string
}
 
func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}
 
// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
```

<!-- tabs:end -->