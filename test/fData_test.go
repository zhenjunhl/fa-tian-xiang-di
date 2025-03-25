package test

import (
	"log"
	"testing"

	"github.com/zhenjunhl/fa-tian-xiang-di/f"
)

func TestGetTimestampHoursAgo(t *testing.T) {
	// ago := f.GetTimestampHoursAgo(1)
	// fmt.Println(ago)
}
func TestName(t *testing.T) {
	var a = "ABCD"
	log.Println(a[:2])
}
func TestParseFile(t *testing.T) {
	type name struct {
		ID int64 `json:"_id"`
	}
	file, err := f.ParseFile[name]("./json.json", "json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file)
}
