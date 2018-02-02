# Go编程基础
Go是一门并发支持、垃圾回收的编译型系统编程语言，旨在创造一门具有在静态编译语言的高性能和动态语言的高效
开发之间拥有良好平衡点的一门编程语言。
Go的主要特点：
    * 类型安全和内存安全；
    * 以非常直观和极低代价的方案实现高并发；
    * 高效的垃圾回收机制；
    * 快速编译；
    * 为多核计算机提供性能提升的方案；
    * UTF-8编码支持。

* Go常用命令
    ```
    go get：获取远程包（需提前安装git或其他代码管理工具）；
    go run：直接运行程序；
    go build：测试编译，检查是否有编译错误；
    go fmt：格式化源码；
    go install：编译包文件并编译整个程序；
    go test：运行测试文件；
    go doc：查看文档。
    ```

* 先写个hello.go吧！
    ```
    package main

    import (
    	"fmt"
    )

    func main() {
    	fmt.Println("Hello World!你好，世界！")
    }
    ```

* Go内置关键字（25个均为小写）
    ```
    break       default     func        interface       select
    case        defer       go          map             struct
    chan        else        goto        package         switch
    const       fallthrough if          range           type
    continue    for         import      return          var
    ```

* Go注释方法
    ```
    //：单行注释
    /* */：多行注释
    ```

* Go程序的一般结构：basic_structure.go
    ```
    Go程序是通过package来组织的（与python类似）；
    只有package名称为main的包可以包含main函数；
    一个可执行程序有且仅有一个main包；
    通过import关键字来导入其他非main包；
    通过const关键字来进行常量的定义；
    通过在函数体外部使用var关键字来进行全局变量的声明与赋值；
    通过type关键字来进行结构(struct)或接口(interface)的声明；
    通过func关键字来进行函数的声明；
    如果导入一个包但是没有用会出现未使用包的异常。
    举个例子，编辑print.go：
    // 当前程序的包名
    package main

    // 导入其它的包
    import (
    	"fmt"
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
    	fmt.Println("Hello,World!")
    }

    # 使用别名
    import std "fmt"
    这样可以使用std.Println
    ```