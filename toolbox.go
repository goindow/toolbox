// Package toolbox 封装了一些开发调试工具
package toolbox

import (
	"fmt"
	"time"
	"reflect"
	"runtime"
)

// Dump 打印变量的类型和值，支持同时打印多个变量，
//
// 本质是调用 fmt.Printf()，每个变量格式为 "[%T] %v\n"
func Dump(v ...interface{}) {
	var format string
	slice := make([]interface{}, 0, len(v))
	for _, j := range v {
		format += "[%T] %v\n"
		slice = append(slice, j, j)
	}
	fmt.Printf(format, slice...)
}

// Trace 计算函数/方法运行耗时
//
// 无参函数，trace(func1)
// 有参函数，trace(func4, 3, true)
// 有参方法，trace(new(reflectStruct).func4, 5, false)
func Trace(f interface{}, args ...interface{}) {
	// 1.处理函数
	// 获取 f 反射类型，reflect.Type
	// rftF := reflect.TypeOf(f)
	// 获取 f 反射值，reflect.Value
	rfvF := reflect.ValueOf(f)
	// 获取 f 反射分类，reflect.Kind
	if rfvF.Kind() != reflect.Func {
		panic(fmt.Sprintf("the first argument must be func[function|method]"))
	}
	// 获取 f 函数名（完整的调用栈名）
	name := runtime.FuncForPC(rfvF.Pointer()).Name()
	// 2.处理参数
	// 转换 args，[]interface{} -> []reflec.Value
	var rfvArgs = make([]reflect.Value, len(args))
	for i, j := range args {
		rfvArgs[i] = reflect.ValueOf(j)
	}
	// 3.调用方法
	// 耗时计时
	begin := time.Now()
	_ = rfvF.Call(rfvArgs)
	//toolbox.Dump(rfvF, name, args, rfvArgs, ret)
	fmt.Printf("\ncall: %v\ntime: %v\n\n", name, time.Now().Sub(begin))
}
