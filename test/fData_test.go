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
func TestMergeAndDeduplicateStrings(t *testing.T) {
	strings := f.MergeAndDeduplicateStrings([]string{"1", "2"}, []string{"2", "3"})
	log.Println(strings)
}
