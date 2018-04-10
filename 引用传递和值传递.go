package main

import "fmt"

func main() {
	//一个小坑()
	//小坑的正确写法()
	把ab两个不定长的排好序的数组把b中的值逐个排序放入a中()
}

//一个小坑
func 一个小坑() {
	a := []string{"1", "2", "3", "4", "5"}
	tmp := a[3:]
	fmt.Println("一开始的tmp", tmp)
	a[3] = "aaa"
	fmt.Println("之后的tmp", tmp)
	//因为slice和map都是引用传递
}

func 小坑的正确写法() {
	var tmp = make([]string, 0)
	a := []string{"1", "2", "3", "4", "5"}
	tmp = append(tmp, a[3:]...)
	fmt.Println("一开始的tmp", tmp)
	a[3] = "aaa"
	fmt.Println("之后的tmp", tmp)
}

func 把ab两个不定长的排好序的数组把b中的值逐个排序放入a中() {
	a := []int{1, 2, 3, 4, 7, 8}
	b := []int{1, 2, 3, 5, 6, 9, 10}
	var i,j int
	for {
		if j >= len(b) {
			break
		}
		if i >= len(a) {
			a = append(a, b[j])
			j++
			continue
		}
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			var tmp = make([]int, 0)
			tmp = append(tmp, a[i:]...)
			a = append(a[:i], b[j])
			a = append(a, tmp...)
			i++
			j++
		case a[i] == b[j]:
			i++
			j++
		}
 	}
 	fmt.Println("====", a)
}
