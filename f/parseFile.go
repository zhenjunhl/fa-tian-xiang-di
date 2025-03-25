package f

import (
	"bufio"
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
	defer file.Close()
	switch fileType {
	case "json":
		// 得倒文件内容
		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line = line + scanner.Text()
		}
		// 解析文件内容
		err = json.Unmarshal([]byte(line), &result)
		if err != nil {
			return result, errors.New("解析失败" + err.Error())
		}
	default:
		return result, errors.New("不支持的文件类型")
	}
	return result, nil
}
