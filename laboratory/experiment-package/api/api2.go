package api

import (
	"fmt"
	otherApi "github.com/gudqs7/learn-go/laboratory/experiment-package/other/api"
)

// 测试表明: 小写开头的可见性仅位于同包下
// 测试表明: 当导入的包依赖其他包时, 会先导入其他包, 同时, 其他包下的 init 函数也会先一步执行

func init() {
	fmt.Println("api2.go init invoke")
}

func sayHello() {
	fmt.Println("api2.go sayHello invoke")
	fmt.Print("api2.go sayHello from otherApi: ")
	otherApi.SayHello()
}
