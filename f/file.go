package f

import (
	"os"
)

// WriteToFileIfExists 写入文件
func WriteToFileIfExists(
	filePath string,
	content string,
) {
	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 创建文件
		_, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
	}

	// 如果文件存在，打开文件以追加模式写入内容
	file, err := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		return
	}
}
