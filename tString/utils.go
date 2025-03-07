package tString

import (
	"github.com/bytedance/sonic"
)

// ToString 转换为 String
func ToString(obj interface{}) string {
	str, _ := sonic.MarshalString(obj)
	return str
}
