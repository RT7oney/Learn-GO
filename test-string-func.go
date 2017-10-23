package main

import "fmt"

func main(){
     symbol := []func()string{test1,test2, test3}
     for _,v := range symbol{
        fmt.Println( v())
	 }
	 
	 symbol := map[string]func()string{"test1":test1, "test2":test2, "test":test3}
	 s  := "test1"
	 fmt.Println(symbol[s]())
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