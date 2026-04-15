package f

import (
	"os"
		"encoding/csv"
	"reflect"
)

// WriteToFileIfExists 写入文件
/*
 * @param filePath 文件路径
 * @param content 内容
 * @return void
 */
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

// ReadCSVToStruct 解析CSV
/*
 * @param filePath 文件路径
 */
func ReadCSVToStruct[T any](filePath string) ([]T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
reader.FieldsPerRecord = -1 
	allRecords, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(allRecords) <= 1 {
		return []T{}, nil
	}

	headers := allRecords[0]
	headerIndex := make(map[string]int)
	for i, header := range headers {
		headerIndex[header] = i
	}

	var result []T
	tType := reflect.TypeOf((*T)(nil)).Elem()

	for i := 1; i < len(allRecords); i++ {
		record := allRecords[i]
		item := reflect.New(tType).Elem()

		for j := 0; j < tType.NumField(); j++ {
			field := tType.Field(j)
			csvTag := field.Tag.Get("csv")
			if csvTag == "" {
				csvTag = field.Name
			}

			if idx, ok := headerIndex[csvTag]; ok && idx < len(record) {
				fieldValue := item.Field(j)
				if fieldValue.CanSet() {
					fieldValue.SetString(record[idx])
				}
			}
		}

		result = append(result, item.Interface().(T))
	}

	return result, nil
}

