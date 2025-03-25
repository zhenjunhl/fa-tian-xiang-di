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
		ID    int64  `json:"_id"`
		Input string `json:"input"`
	}
	file, err := f.ParseFile[name]("/Users/bytedance/code/byteFaas/ai-insights-webscoket/biz/local/BN32517/json.json", "json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file[0])
}
