package api

import (
	"fmt"
	otherApi "github.com/gudqs7/learn-go/laboratory/experiment-package/other/api"
)

// 测试表明: 小写开头的可见性仅位于同包下

func init() {
	fmt.Println("api2.go init invoke")
}

func sayHello() {
	fmt.Println("api2.go sayHello invoke")
	fmt.Print("api2.go sayHello from otherApi: ")
	otherApi.SayHello()
}
