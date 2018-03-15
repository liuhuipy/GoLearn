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
* 可见性规则
    ```
    Go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包调用。
    根据约定，函数名首字母小写即为private；函数名首字母大写即为public。
    ```
* 多个常量、全局变量和类型的声明、定义或赋值可和import多个包一样使用简写。
    ```
    # 常量的定义
    const {
        PI     = 3.14
        const1 = "1"
        const2 = 2
        const3 = 3
    }
    # 全局变量的声明与赋值
    var {
        name  = "gopher"
        name1 = "1"
        name2 = 2
        name3 = 3
    }
    # 一般类型声明
    type {
        newType int
        type1   float32
        type2   string
        type3   byte
    }
    ```
* 单个变量的声明与赋值
    ```
    - 变量的声明格式：var <变量名称> <变量类型>
    var a int
    - 变量的赋值格式：<变量名称> = <表达式>
    a = 123
    - 变量声明同时赋值：var <变量名称> [变量类型] = <表达式>
    var b int = 321
    - 变量声明同时赋值可省略变量类型，系统会自己判断
    var c = 321
    - 变量声明赋值最简写法
    d := 123
    ```
* 多个变量的声明与赋值
    ```
    - 全局变量的声明可使用var()的方式进行简写
    - 全局变量的声明不可以省略var，但可使用并行方式
    - 所有变量都可以使用类型推断
    - 局部变量不可以使用var()的方式简写，只能使用并行方式
    var {   // 全局变量
        // 使用常规方式
        aaa = "hello"
        // 使用并行方式以及类型推断
        bbb, ccc = 111, 222
    }

    // 局部变量
    // 多个变量的声明
    var a, b, c, d int
    // 多个变量的赋值
    a, b, c, d = 1, 2, 3, 4

    ```

## Go基本数据类型
* 布尔型(bool)：
    ```
    - 长度：1字节
    - 取值范围：true，false
    - 注意事项：不可以用数字代表true或false
    ```
* 整形(int/uint)：
    ```
    - 根据运行平台可能为32或64位
    ```
* 8位整形(int8/uint8)：
    ```
    - 长度：1字节
    - 取值范围：-128～127/0～255
    ```
* 字节型：byte(uint8别名)
* 16位整形(int16/uint16)：
    ```
    - 长度：2字节
    - 取值范围：-32768～32767/0～65535
    ```
* 32位整形(int32(rune)/uint32)：
    ```
    - 长度：4字节
    - 取值范围：-2^32/2~2^32/2-1/0~2^32-1
    ```
* 64位整形(int64/uint64)：
    ```
    - 长度：8字节
    - 取值范围：-2^64/2~2^64/2-1/0~2^64-1
    ```
* 浮点型(float32/float64)：
    ```
    - 长度：4/8字节
    - 小数点：精确到7/15小数位
    ```
* 复数(complex64/complex128)：
    ```
    - 长度：8/16字节
    ```
* 足够保存指针的32位或64位整数型(nintptr)
* 其它值类型：
    ```
    - array、struct、string
    ```
* 引用类型：
    ```
    - slice、map、chan
    ```
* 接口类型(interface)
* 函数类型(func)

* 类型转换
    - Go语言中不存在隐式转换，所有类型转换必须显式声明
    - 转换只能发生在两种相互兼容的类型之间
    - 类型转换的格式： <Value> [:]= <TypeOfValueA>(<ValueB>)
        ```
        // 在相互兼容的两种类型之间进行转换
        var a float32 = 1.1
        b := int(a)
        // 以下表达式无法通过编译
        var c bool = true
        d := int(c)
        ```




