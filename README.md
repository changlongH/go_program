# Go 学习记录 #
golang 

# Go 的基础组成 #
1. 包声明 (package)
2. 引入包 (import)
3. 函数 (func)
4. 变量
5. 语句 & 表达式
6. 注释

# go 指令 #
1. go run, go build  go help

# 关键字25 #
```
package     import      go          interface   func
type        var         range       map         chan
const       default     fallthrough struct      defer
return      break       continue    goto        select
if          else        case        switch      for
```

# 预定义符号36 #
```
append      bool        byte        cap         close
copy        false       float32     imag        print
iota        len         make        new         nil        panic
println     real        recover     string      ture        uintptr
float(32 64)     complex(x64 x128)      int(8 16 32 64)         uint(8 16 32 64)
```

# 变量声明 #
1. car v_name v_type
v_name = value

2. var v_name = value

3. v_name := value

- iota 在const出现的时候重置为0, const 中每一行计数一次

# 引用传递和值传递 #
- &a 取变量a地址 其他没啥好说

# 运算符 #
- 也没啥好说和C一样


# 循环语句 #
- for
```
c 一样
for init: condition; post{}

C 的while一样
for condition {}

C 的for {;;}一样

for {}

for key, value := range  oldMap {
    newMap[key] = value
}

etc.
func main() {

    var b int = 15
        var a int

        numbers := [6]int{1, 2, 3, 5} 

    /* for 循环 */
    for a := 0; a < 10; a++ {
        fmt.Printf("a 的值为: %d\n", a)
    }

    for a < b {
        a++
            fmt.Printf("a 的值为: %d\n", a)
    }

    for i,x:= range numbers {
        fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
    }   
}
```

# go 语言函数 #
```
func function_name( [parameter list] ) [return_types] {
       //函数体
}
```
- 可以返回多个值
- 默认值传递
- 函数闭包
- 方法

