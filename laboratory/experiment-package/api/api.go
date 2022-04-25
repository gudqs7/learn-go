package api

import "fmt"

// 测试表明: 支持多个 init 方法, 执行顺序从上往下

func init() {
	fmt.Println("api.go init invoke")
}

func init() {
	fmt.Println("api.go init2 invoke")
}
