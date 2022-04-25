package api

import "fmt"

// 测试表明: 大写开头的可见性表现良好
// 测试表明: init 方法中可调用其他包的函数, 同包的函数

func init() {
	fmt.Println("api3.go init invoke")
	fmt.Print("api3.go init invoke sayHello: ")
	sayHello()
}

func SayHello(name string) string {
	return "Hello " + name
}
