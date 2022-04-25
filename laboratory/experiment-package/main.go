package main

import (
	"fmt"
	"github.com/gudqs7/learn-go/laboratory/experiment-package/api"
	otherApi "github.com/gudqs7/learn-go/laboratory/experiment-package/other/api"
)

func main() {
	// 1.测试 init 方法的执行情况
	//   结果: 导入一个包, 该包下每个文件的 init 函数(init函数不会冲突)都会执行, import 是有序的(包有序, 文件亦有序)
	// 2.测试包与目录的关系
	//   结果: 目录下一级文件中无法包含多个不同的 package, 否则会造成编译报错; 包名与当前目录名称不要求一致, 但 import 时以目录名为准
	// 3.测试同名包冲突的情况
	//   结果: 完整包名不重复, 最后包名重复时, import 时添加别名区分即可
	// 4.测试同包下同函数名冲突的情况
	//   结果: init 同文件下, 同包下, 可包含多个, 执行顺序与代码顺序等同;  go 不支持函数重载, 函数名重复即编译报错, 不关系参数是否相同
	// 5.测试函数名称首字母大小写对函数可见性的影响
	//   结果: 小写时, 仅同包下代码可见; 大写时, import 后可见
	// 6.测试 import 时包名的生成规则
	//   结果: go.mod 中 module 的名称 加上从 go.mod 所在的目录开始算起(不包含), 文件位于的相对路径;
	//          如 module 改为 laboratory, 则 import otherApi 时为 "laboratory/experiment-package/other/api"
	fmt.Println("main invoke")
	hello := api.SayHello("wq")
	fmt.Println("say hello res: ", hello)
	otherApi.SayHello()
}
