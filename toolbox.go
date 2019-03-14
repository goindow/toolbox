// Package toolbox 封装了一些开发调试工具
package toolbox

import (
	"fmt"
)

// Dump 打印变量的类型和值，支持同时打印多个变量
func Dump(v ...interface{}) {

	var format string
	slice := make([]interface{}, 0, len(v))

	for _, j := range v {
		format += "[%T] %v\n"
		slice = append(slice, j, j)
	}
	fmt.Printf(format, slice...)
}
