package f

import (
	"encoding/json"
	"errors"
	"os"
)

func ParseFile[T any](filePath, fileType string) ([]T, error) {
	var result []T
	file, err := os.Open(filePath)
	if err != nil {
		return result, errors.New("打开文件失败")
	}
	switch fileType {
	case "json":
		// 得倒文件内容
		data, err := os.ReadFile(filePath)
		if err != nil {
			return result, errors.New("读取文件失败")
		}
		// 解析文件内容
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			return result, errors.New("解析失败" + err.Error())
		}
	default:
		return result, errors.New("不支持的文件类型")
	}
	defer file.Close()
	return result, nil
}
