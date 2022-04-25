package otherApi

import "fmt"

// 测试表明包名不一定等于目录名
// 测试表明: init 的执行顺序受 import 包时的顺序影响
// 测试表明: 最后包名重复时 import 需添加别名区分

func init() {
	fmt.Println("other/api.go init invoke")
}

func SayHello() {
	fmt.Println("[other] Hello!")
}
