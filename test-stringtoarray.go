package main
//1.func Fields(s string) []string，这个函数的作用是按照1：n个空格来分割字符串最后返回的是[]string的切片
代码如下:

import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.Fields("hello widuu golang")) //out  [hello widuu golang]
}

//2.func FieldsFunc(s string, f func(rune) bool) []string一看就了解了，这就是根据自定义函数分割了
代码如下:

import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.FieldsFunc("widuunhellonword", split)) // [widuu hello word]根据n字符分割
}
func split(s rune) bool {
 if s == 'n' {
  return true
 }
 return false
}

//3.func Join(a []string, sep string) string,这个跟php中的implode差不多，这个函数是将一个[]string的切片通过分隔符，分割成一个字符串
代码如下:

import (
 "fmt"
 "strings"
)
func main() {
 s := []string{"hello", "word", "xiaowei"}
 fmt.Println(strings.Join(s, "-")) // hello-word-xiaowei
}

//4.func Split(s, sep string) []string,有join就有Split这个就是把字符串按照指定的分隔符切割成slice
代码如下:

import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.Split("a,b,c,d,e", ",")) //[a b c d e]
}

//5.func SplitAfter(s, sep string) []string,这个函数是在前边的切割完成之后再后边在加上sep分割符
代码如下:

import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.SplitAfter("a,b,c,d", ",")) //[a, b, c, d]
}

//6.func SplitAfterN(s, sep string, n int) []string该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割
代码如下:

import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 4)) //["a," "b," "c," "d,r"]
 fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 5)) //["a," "b," "c," "d," "r"]
}

//7.func SplitN(s, sep string, n int) []string,这个是切割字符串的时候自己定义长度，如果sep为空，那么每一个字符都分割
代码如下:

import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.SplitN("a,b,c", ",", 2)) //[a b,c]
}