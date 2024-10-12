package stringcode

import (
	"fmt"
	"strings"
)

// splitInput 分割输入字符串，返回一个字符串数组
func SplitLongString(LongString string) []string {
	parts := strings.Split(LongString, ",")
	var result []string

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		if trimmedPart != "" {
			result = append(result, trimmedPart)
		}
	}

	if len(result) == 0 {
		fmt.Println("输入字符串不包含任何用逗号分隔的有效部分")
	}

	return result
}
