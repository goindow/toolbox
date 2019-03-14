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
- 打印变量的类型和值，支持同时打印多个变量，
- 本质是调用 fmt.Printf()，每个变量格式为 "[%T] %v\n"
```go
toolbox.Dump(os.Getwd())

// Output:
// [string] /usr/local/var/go/src/github.com/goindow/toolbox
// [<nil>] <nil>
```