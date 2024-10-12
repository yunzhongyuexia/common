package mapcode

// 查找并返回与给定键对应的值的数组
func FindValues(keysToCheck []any, aMap map[any]any) []any {
	values := make([]any, 0, len(keysToCheck))
	for _, key := range keysToCheck {
		if v, ok := aMap[key]; ok {
			values = append(values, v)
		}
	}
	return values
}

// 将[]string转换为[]any
func StringArrayToAnyArray(stringArray []string) []any {
	var anyArray []any
	for _, key := range stringArray {
		anyArray = append(anyArray, any(key))
	}
	return anyArray
}

// 将map[string]int转换为map[any]any
func StringIntMapToAnyMap(stringIntMap map[string]int) map[any]any {
	anyMap := make(map[any]any)
	for k, v := range stringIntMap {
		anyMap[any(k)] = any(v)
	}
	return anyMap
}
