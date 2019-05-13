// 当前程序的包名
package main;

// 导入其它的包
import (
	sbt "fmt" // 设置别名
	"math"

	//"os" // 注意 导入的包 必须 使用，若没有使用过，则编译不通过
	//"time"
)

// 常量的定义
const PI = 3.14

// 全局变量的声明与赋值
var name = "gopher"

// 一般类型声明
type newType int


// 结构的声明
type gopher struct {

}

// 接口的声明
type golang interface {

}


// 由 main 函数作为程序入口点启动 【一个程序只能有一个 main 包，并且 main 包里面有且仅有一个 main 函数】
func main() {
	// 函数名 首字母 小写：代表这个函数是 private 的；大写：代表这个函数是 public 的
	sbt.Println("Hello world! 你好，世界") // 使用 fmt 的别名

	var a int // 编译器会默认给未初始化的变量 一个零值
	var b string
	var c float32
	var d float64
	var e byte
	var f []int // 不给大小时，是一个 切片
	var g [1]byte
	sbt.Println(a, b, c, d, d, e, f, "----", g)

	aa, bb, cc, _ := 1, 2, 3, 4 // _ 下划线作为一个 忽略符
	sbt.Println(math.MinInt64)
	sbt.Println(aa, bb, cc)

	sbt.Println()

	var abc int = 65
	bbb := string(abc)
	sbt.Println(bbb)
}
