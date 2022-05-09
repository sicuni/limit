package limit

import (
	"fmt"
	"testing"
)

func TestInitLimit(t *testing.T) {
	limit := InitLimit()

	str := "111"
	if limit.Exists(str) {
		fmt.Println("limit str:", str)
	}
	// 此处被limit
	if limit.Exists(str) {
		fmt.Println("limit str:", str)
	}

	key := 0
	if limit.Exists(key) {
		fmt.Println("limit key:", key)
	}
	limit.Delete(key) // 解除限制
	if limit.Exists(key) {
		fmt.Println("limit key:", key)
	}
}
