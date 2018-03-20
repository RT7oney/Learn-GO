/**
 * 2
 */
// package main

// type student struct {
// 	Name string
// 	Age  int
// }

// func pase_student() {
// 	m := make(map[string]*student)
// 	stus := []student{
// 		{Name: "zhou", Age: 24},
// 		{Name: "li", Age: 23},
// 		{Name: "wang", Age: 22},
// 	}
// 	// 错误
// 	// 与Java的foreach一样，都是使用副本的方法，所以m[stu.Name] = &stu实际上一直指向同一个指针，最终指针的值为遍历的最后一个struct的值拷贝
// 	// for _, stu := range stus {
// 	// 	println(&stu)
// 	// 	m[stu.Name] = &stu
// 	// }
// 	// 正确写法
// 	for i := 0; i < len(stus); i++ {
// 		println(&stus[i])
// 		m[stus[i].Name] = &stus[i]
// 	}
// 	for k, v := range m {
// 		println(k, "=>", v.Name)
// 	}
// }

// func main() {
// 	pase_student()
// }

/**
 * 3
 */
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	wg := sync.WaitGroup{}
// 	wg.Add(q20)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			fmt.Println("A: ", i)
// 			wg.Done()
// 		}()
// 	}
// 	for i := 0; i < 100; i++ {
// 		go func(i int) {
// 			fmt.Println("B: ", i)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

/**
 * 4
 */
// package main

// import (
// 	"fmt"
// 	// "runtime"
// 	// "sync"
// )

// type People struct{}

// func (p *People) ShowA() {
// 	fmt.Println("ShowA")
// 	p.ShowB()
// }

// func (p *People) ShowB() {
// 	fmt.Println("ShowB")
// }

// type Teacher struct {
// 	People
// }

// func (t *Teacher) ShowB() {
// 	fmt.Println("Teacher ShowB")
// }

// func main() {
// 	t := Teacher{}
// 	t.ShowB()
// }

/**
 * 8
 */
// package main

// import (
// 	"sync"
// )

// var flag chan string

// type UserAges struct {
// 	ages map[string]int
// 	sync.Mutex
// }

// func (ua *UserAges) Add(name string, age int) {
// 	ua.Lock()
// 	defer ua.Unlock()
// 	ua.ages = make(map[string]int)
// 	ua.ages[name] = age
// 	<-flag
// }

// func (ua *UserAges) Get(name string) int {
// 	// ua.Lock()
// 	// defer ua.Unlock()
// 	if age, ok := ua.ages[name]; ok {
// 		return age
// 	}
// 	return -1
// }

// func main() {
// 	flag = make(chan string)
// 	ua := UserAges{}
// 	go ua.Add("amy", 12)
// 	flag <- "amy"
// 	age := ua.Get("amy")
// 	println("age", age)
// }

/**
 * 9
 */
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type threadSafeSet struct {
// 	sync.RWMutex
// 	s []interface{}
// }

// func (set *threadSafeSet) Iter() <-chan interface{} {
// 	// ch := make(chan interface{})
// 	ch := make(chan interface{}, len(set.s))
// 	go func() {
// 		set.RLock()
// 		for k, v := range set.s {
// 			ch <- k
// 			println("Iter: ", k, v)
// 		}
// 		close(ch)
// 		set.RUnlock()
// 	}()
// 	return ch
// }

// func main() {
// 	th := threadSafeSet{
// 		s: []interface{}{"1", "2"},
// 	}
// 	v := <-th.Iter()
// 	fmt.Println("ch", v)
// }

/**
 * 10
 */
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type threadSafeSet struct {
// 	sync.RWMutex
// 	s []interface{}
// }

// func (set *threadSafeSet) Iter() <-chan interface{} {
// 	// ch := make(chan interface{})
// 	ch := make(chan interface{}, len(set.s))
// 	go func() {
// 		set.RLock()
// 		for k, v := range set.s {
// 			ch <- k
// 			println("Iter: ", k, v)
// 		}
// 		close(ch)
// 		set.RUnlock()
// 	}()
// 	return ch
// }

// func main() {
// 	th := threadSafeSet{
// 		s: []interface{}{"1", "2"},
// 	}
// 	v := <-th.Iter()
// 	fmt.Println("ch", v)
// }

/**
 * 27
 */
// package main

// func test() []func() {
// 	var funcs []func()
// 	for i := 0; i < 2; i++ {
// 		funcs = append(funcs, func() {
// 			println(&i, i)
// 		})
// 	}
// 	return funcs
// }

// func main() {
// 	funcs := test()
// 	for _, f := range funcs {
// 		f()
// 	}
// }

/********************个人测试*********************/

// package main

// import (
// 	"sync"
// )

// func main() {
// 	wg := &sync.WaitGroup{}
// 	wg.Add(1)
// 	m := 3
// 	hi := "hello"
// 	go func(i int) {
// 		defer wg.Done()
// 		println(i)
// 		println(hi)
// 	}(m)
// 	wg.Wait()
// }

/**********************************************************/

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type demo struct {
// 	m map[string]int
// 	sync.RWMutex
// }

// func main() {

// 	var d demo
// 	d.m = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}

// 	var i, j, k chan int
// 	i = make(chan int)
// 	j = make(chan int)
// 	k = make(chan int)

// 	fmt.Println(d.m)
// 	fmt.Println("-------")

// 	// wg := &sync.WaitGroup{}
// 	// wg.Add(3)
// 	go func(p map[string]int) {
// 		// d.Lock()
// 		// defer wg.Done()
// 		// defer d.Unlock()
// 		println("我把a改成了10")
// 		p["a"] = 10
// 		println("1.1", <-i)
// 		// println("1.2", <-j)
// 		// println("1.3", <-k)
// 		<-i
// 	}(d.m) //因为go语言中的map为引用类型， 虽然go语言函数以传值方式调用，即函数持有的为参数副本，但因是引用类型， 故依然指向映射m ， 类似c的指针.
// 	go func(p map[string]int) {
// 		// d.Lock()
// 		// defer wg.Done()
// 		// defer d.Unlock()
// 		println("我把a改成了11")
// 		p["a"] = 11
// 		// println("2.1", <-i)
// 		println("2.2", <-j)
// 		// println("2.3", <-k)
// 		<-j
// 	}(d.m)
// 	go func(p map[string]int) {
// 		// d.Lock()
// 		// defer wg.Done()
// 		// defer d.Unlock()
// 		println("我把a改成了12")
// 		p["a"] = 12
// 		// println("3.1", <-i)
// 		// println("3.2", <-j)
// 		println("3.3", <-k)
// 		<-k
// 	}(d.m)

// 	// wg.Wait()
// 	i <- 0
// 	j <- 1
// 	k <- 2
// 	fmt.Println(d.m)
// 	/*
// 		因为map为引用类型，所以即使函数传值调用，参数副本依然指向映射m, 所以3个goroutine并发写同一个映射m， 写过多线程程序的同学都知道，对于共享变量，资源，
// 		并发读写会产生竞争的， 故共享资源遭到破坏， 所以要么加锁， 要么用channel排队串行化， 总之要排他访问。
// 		所以切记： 多goroutine读写同一个映射时， 要保护哟， 加锁也可， 利用channel串行化处理也可！！！
// 	*/
// }

/************************************************/

// #示例代码一：
// func funcA() int {
//     x := 5
//     defer func() {
//         x += 1
//     }()
//     return x
// }
// #示例代码二：
// func funcB() (x int) {
//     defer func() {
//         x += 1
//     }()
//     return 5
// }
// <!--more-->
// #示例代码三：
// func funcC() (y int) {
//     x := 5
//     defer func() {
//         x += 1
//     }()
//     return x
// }

// #示例代码四：
// func funcD() (x int) {
//     defer func(x int) {
//         x += 1
//     }(x)
//     return 5
// }

// 解析这几段代码，主要需要理解清楚以下几点知识：
// 1、return语句的处理过程
// return xxx 语句并不是一条原子指令，其在执行的时候会进行语句分解成 返回变量=xxx return，最后执行return
// 2、defer语句执行时机
// 上文说过，defer语句是在函数关闭的时候调用，确切的说是在执行return语句的时候调用，注意，是return 不是return xxx
// 3、函数参数的传递方式
// Go语言中普通的函数参数的传递方式是值传递，即新辟内存拷贝变量值，不包括slice和map，这两种类型是引用传递
// 4、变量赋值的传递方式
// Go语言变量的赋值跟函数参数类似，也是值拷贝，不包括slice和map，这两种类型是内存引用

// 按照以上原则，解析代码：

// #示例代码一：
// func funcA() int {
//     x := 5
//     temp=x      #temp变量表示未显示声明的return变量
//     func() {
//         x += 1
//     }()
//     return
// }
// 返回temp的值，在将x赋值给temp后，temp未发生改变，最终返回值为5

// #示例代码二：
// func funcB() (x int) {
//     x = 5
//     func() {
//         x += 1
//     }()
//     return
// }
// 返回x的值，先对其复制5，接着函数中改变为6，最终返回值为6

// #示例代码三：
// func funcC() (y int) {
//     x := 5
//     y = x       #这里是值拷贝
//     func() {
//         x += 1
//     }()
//     return
// }
// 返回y的值，在将x赋值给y后，y未发生改变，最终返回值为5

// #示例代码四：
// func funcD() (x int) {
//     x := 5
//     func(x int) { #这里是值拷贝
//         x += 1
//     }(x)
//     return
// }
// 返回x的值，传递x到匿名函数中执行的时候，传递的是x的拷贝，其内部修改不影响外部x的值，最终返回值为5

/********************个人测试*********************/

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}
// 	wg := &sync.WaitGroup{}
// 	wg.Add(1)

// 	go func(p map[string]int) {
// 		defer wg.Done()
// 		m["a"] = 111
// 	}(m)

// 	wg.Wait()

// 	fmt.Println(m)
// }

/********************个人测试*********************/

// package main

// import (
// 	"fmt"
// )

// var ch chan int = make(chan int)

// func foo(id int) { //id: 这个routine的标号
//     ch <- id
// }

// type MyError struct{
// 	Code int
// 	Info string
// }

// func (e *MyError) Error() string {
// 	if e.Code == 1 {
// 		return "error"
// 	}
// 	return "another error"
// }

// func (e *MyError) Hehe() string {
// 	if e.Code == 2 {
// 		return "heheheh"
// 	}
// 	return "hahahaha"
// }

// func test(i int) error {
// 	var e MyError
// 	e.Code = i
// 	s := e.Hehe()
// 	fmt.Println(s)
// 	return &e
// }

// // func main() {
// //     // 开启5个routine
// //     for i := 0; i < 5; i++ {
// //         go foo(i)
// //     }

// //     // 取出信道中的数据
// //     for i := 0; i < 5; i++ {
// //         fmt.Print(<- ch)
// //     }
// // }

// // func main() {
// //     ch := make(chan int, 3)
// //     ch <- 1
// //     ch <- 2
// //     ch <- 3

// //     fmt.Println(<-ch) // 1
// //     fmt.Println(<-ch) // 2
// //     fmt.Println(<-ch) // 3
// // }

// func main() {
// 	t := test(0)
// 	fmt.Println(t)
// }

/********************理解interface*********************/

package main

import (
	"fmt"
)

type People interface {
	SayMyName() string
	HowOldAreYou(People) int
}

type Teacher struct {
	Name  string
	Age   int
	Teach string
}

type Student struct {
	Name  string
	Age   int
	Learn string
}

func (t *Teacher) SayMyName() string {
	return "Haha, my name is " + t.Name + " i teach " + t.Teach
}

func (t *Teacher) HowOldAreYou(s People) int {
	Stu_s := s.(*Student)
	return Stu_s.Age
}

func (s *Student) SayMyName() string {
	return "Hehe, my name is " + s.Name + " i learn " + s.Learn
}

func (s *Student) HowOldAreYou(t People) int {
	Tea_t := t.(*Teacher)
	return Tea_t.Age - 5
}

func who_are_you(m interface{}) People {
	switch m.(type) {
	case Teacher:
		var t = m.(Teacher)
		return &t
	case Student:
		var s = m.(Student)
		return &s
	}
	return nil
}

func main() {
	t := Teacher{
		"Zhang",
		45,
		"math",
	}
	s := Student{
		"xiaoming",
		15,
		"chinese",
	}
	// 测试老师
	interface_t := who_are_you(t)
	fmt.Println(interface_t.SayMyName())
	fmt.Println(interface_t.HowOldAreYou(&s))
	// 测试学生
	var p People
	p = &s
	fmt.Println(p.SayMyName())
	var _p *People
	_p = &p
	fmt.Println((*_p).SayMyName())
}

// 使用接口的美感
// 假设某公司有两个员工，一个普通员工和一个高级员工， 但是基本薪资是相同的，高级员工多拿奖金。计算公司为员工的总开支。


package main

import (
	"fmt"
)

// 薪资计算器接口
type SalaryCalculator interface {
	CalculateSalary() int
}
// 普通挖掘机员工
type Contract struct {
	empId  int
	basicpay int
}
// 有蓝翔技校证的员工
type Permanent struct {
	empId  int
	basicpay int
	jj int // 奖金
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}
// 总开支
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("总开支 $%d", expense)
}

func main() {
	pemp1 := Permanent{1,3000,10000}
	pemp2 := Permanent{2, 3000, 20000}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
