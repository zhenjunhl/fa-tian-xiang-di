package test

import (
	"fmt"
	"testing"

	"github.com/zhenjunhl/fa-tian-xiang-di/f"
)

func TestGetTimestampHoursAgo(t *testing.T) {
	ago := f.GetTimestampHoursAgo(1)
	fmt.Println(ago)
}
