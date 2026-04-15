package test

import (
	"fmt"
	"testing"

	"github.com/zhenjunhl/fa-tian-xiang-di/f"
)

func TestName(t *testing.T) {
	arr := f.SplitNumberByArr(10, 9)
	fmt.Println(arr)
}
