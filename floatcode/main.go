package floatcode

import (
	"math/big"
)

// 使用big.Float提供任意精度的浮点数运算
func FloatArrayAdd(values []float64) float64 {
	sum := new(big.Float).SetUint64(0)
	for _, v := range values {
		fv := new(big.Float).SetFloat64(v)
		sum.Add(sum, fv)
	}
	floatSum, _ := sum.Float64()
	return floatSum
}
