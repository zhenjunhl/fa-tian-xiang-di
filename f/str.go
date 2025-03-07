package f

import (
	"github.com/bytedance/sonic"
)

// ToString 转换为 String
/*
 * @param obj 要转换的对象
 * @return string 转换后的字符串
 */
func ToString(obj interface{}) string {
	str, _ := sonic.MarshalString(obj)
	return str
}
