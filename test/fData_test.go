package test

import (
	"fmt"
	"testing"

	"github.com/zhenjunhl/fa-tian-xiang-di/fData"
)

func TestGetTimestampHoursAgo(t *testing.T) {
	ago := fData.GetTimestampHoursAgo(1)
	fmt.Println(ago)
}
