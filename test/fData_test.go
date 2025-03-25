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
	type GeneralAnnotationRecordType struct {
		Id int64 `json:"id"`
	}
	file, err := f.ParseFile[GeneralAnnotationRecordType]("aas", "json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file[0])
}
