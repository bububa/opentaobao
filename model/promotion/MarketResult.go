package promotion

import (
	"sync"
)

// MarketResult 结构体
type MarketResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的匿名码对象
	Data *CouponActivity `json:"data,omitempty" xml:"data,omitempty"`
}

var poolMarketResult = sync.Pool{
	New: func() any {
		return new(MarketResult)
	},
}

// GetMarketResult() 从对象池中获取MarketResult
func GetMarketResult() *MarketResult {
	return poolMarketResult.Get().(*MarketResult)
}

// ReleaseMarketResult 释放MarketResult
func ReleaseMarketResult(v *MarketResult) {
	v.Message = ""
	v.Data = nil
	poolMarketResult.Put(v)
}
