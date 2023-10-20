package car

import (
	"sync"
)

// TaobaoAlitripCarOrderAcceptApiResult 结构体
type TaobaoAlitripCarOrderAcceptApiResult struct {
	// 其它数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码 0成功 其它见文档
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 错误码 0成功 其它见文档
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

var poolTaobaoAlitripCarOrderAcceptApiResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderAcceptApiResult)
	},
}

// GetTaobaoAlitripCarOrderAcceptApiResult() 从对象池中获取TaobaoAlitripCarOrderAcceptApiResult
func GetTaobaoAlitripCarOrderAcceptApiResult() *TaobaoAlitripCarOrderAcceptApiResult {
	return poolTaobaoAlitripCarOrderAcceptApiResult.Get().(*TaobaoAlitripCarOrderAcceptApiResult)
}

// ReleaseTaobaoAlitripCarOrderAcceptApiResult 释放TaobaoAlitripCarOrderAcceptApiResult
func ReleaseTaobaoAlitripCarOrderAcceptApiResult(v *TaobaoAlitripCarOrderAcceptApiResult) {
	v.Data = ""
	v.Message = ""
	v.Code = 0
	v.MessageCode = 0
	poolTaobaoAlitripCarOrderAcceptApiResult.Put(v)
}
