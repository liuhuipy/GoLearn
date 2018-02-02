// 当前程序的包名
package main

// 导入其它的包
import (
	std "fmt"
)

// 定义常量
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类似声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由main函数作为程序的入口
func main() {
	std.Println("Hello,World!")
}