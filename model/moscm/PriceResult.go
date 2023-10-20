package moscm

import (
	"sync"
)

// PriceResult 结构体
type PriceResult struct {
	// 如果出错，返回错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 价格对象
	PriceDto *PriceDto `json:"price_dto,omitempty" xml:"price_dto,omitempty"`
	// 单挑成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPriceResult = sync.Pool{
	New: func() any {
		return new(PriceResult)
	},
}

// GetPriceResult() 从对象池中获取PriceResult
func GetPriceResult() *PriceResult {
	return poolPriceResult.Get().(*PriceResult)
}

// ReleasePriceResult 释放PriceResult
func ReleasePriceResult(v *PriceResult) {
	v.Message = ""
	v.PriceDto = nil
	v.Success = false
	poolPriceResult.Put(v)
}
