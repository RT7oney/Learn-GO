package main

import "fmt"

func main(){
     symbol1 := []func()string{test1,test2, test3}
     for _,v := range symbol1{
        fmt.Println( v())
	 }
	 
	//  假如要实现s = "test", $S(), 也就是key=value的方式， 所以改造上边symbol为一个map

	symbol2 := map[string]func()string{"test1":test1, "test2":test2, "test":test3}
    s  := "test1"
    fmt.Println(symbol2[s]())
}


func test1() string {
    return "test1"
}

func test2() string {
    return "test2"
}

func test3() string {
    return "test3"
}