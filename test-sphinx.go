package main

import (
	"fmt"
	"github.com/yunge/sphinx"
)

func main() {
    // 链接参数
    opts := &sphinx.Options{
        Host: "127.0.0.1",
        Port: 9312,
        Timeout: 5000,
    }

	// 创建客户端
	/*************分割线***************/
    // spClient := &sphinx.Client{Options: opts}
    // if err := spClient.Error(); err != nil {
    //     fmt.Println("19---", err)
    // }
	/*********************************/
	spClient := sphinx.NewClient(opts)
 
	if err := spClient.Error(); err != nil {
        fmt.Println("19---", err)
    }

    // 打开链接
    if err := spClient.Open(); err != nil {
        fmt.Println("24---", err)
    }

    // 获取实例信息
    status, err := spClient.Status()
    if err != nil {
        fmt.Println("30---", err)
    }

    for _, row := range status {
        fmt.Println("结果", row[0], row[1])
    }

    // 查询
	res, err := spClient.Query("几粒", "index_one", "this test")
	if err != nil {
		fmt.Println("40---", err)
	}

	fmt.Println("最终结果", res)
}