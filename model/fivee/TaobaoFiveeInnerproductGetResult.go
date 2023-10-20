package fivee

import (
	"sync"
)

// TaobaoFiveeInnerproductGetResult 结构体
type TaobaoFiveeInnerproductGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *InnerProduct `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFiveeInnerproductGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeInnerproductGetResult)
	},
}

// GetTaobaoFiveeInnerproductGetResult() 从对象池中获取TaobaoFiveeInnerproductGetResult
func GetTaobaoFiveeInnerproductGetResult() *TaobaoFiveeInnerproductGetResult {
	return poolTaobaoFiveeInnerproductGetResult.Get().(*TaobaoFiveeInnerproductGetResult)
}

// ReleaseTaobaoFiveeInnerproductGetResult 释放TaobaoFiveeInnerproductGetResult
func ReleaseTaobaoFiveeInnerproductGetResult(v *TaobaoFiveeInnerproductGetResult) {
	v.Message = ""
	v.Code = 0
	v.Data = nil
	v.Success = false
	poolTaobaoFiveeInnerproductGetResult.Put(v)
}
