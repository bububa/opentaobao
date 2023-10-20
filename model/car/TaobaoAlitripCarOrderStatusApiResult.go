package car

import (
	"sync"
)

// TaobaoAlitripCarOrderStatusApiResult 结构体
type TaobaoAlitripCarOrderStatusApiResult struct {
	// 其它数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoAlitripCarOrderStatusApiResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderStatusApiResult)
	},
}

// GetTaobaoAlitripCarOrderStatusApiResult() 从对象池中获取TaobaoAlitripCarOrderStatusApiResult
func GetTaobaoAlitripCarOrderStatusApiResult() *TaobaoAlitripCarOrderStatusApiResult {
	return poolTaobaoAlitripCarOrderStatusApiResult.Get().(*TaobaoAlitripCarOrderStatusApiResult)
}

// ReleaseTaobaoAlitripCarOrderStatusApiResult 释放TaobaoAlitripCarOrderStatusApiResult
func ReleaseTaobaoAlitripCarOrderStatusApiResult(v *TaobaoAlitripCarOrderStatusApiResult) {
	v.Data = ""
	v.Message = ""
	v.Code = 0
	poolTaobaoAlitripCarOrderStatusApiResult.Put(v)
}
