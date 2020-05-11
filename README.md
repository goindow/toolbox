# toolbox
Package toolbox 封装了一些开发中使用的工具

## 索引
- [Dump](#Dump)

### 使用
```go
import (
    "github.com/goindow/toolbox"
)
```

### Dump
- 打印变量的类型和值，支持同时打印多个变量
- 本质是调用 fmt.Printf()，每个变量格式为 "[%T] %v\n"
```go
toolbox.Dump(os.Getwd())

// Output:
// [string] /usr/local/var/go/src/github.com/goindow/toolbox
// [<nil>] <nil>
```

### Trace
- 计算函数/方法运行耗时
```go
// 无参函数
trace(func1)
// 有参函数
trace(func4, 3, true)
// 有参方法
trace(new(reflectStruct).func4, 5, false)

// Output:
// call: main.func1
// time: 52.561µs
//
// call: main.func4
// time: 18.698µs
//
// call: main.(*reflectStruct).func4-fm
// time: 4.168µs
```