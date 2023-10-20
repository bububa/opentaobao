package trade

import (
	"sync"
)

// FastBuyPosPayResult 结构体
type FastBuyPosPayResult struct {
	// 返回的错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 错误描述
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFastBuyPosPayResult = sync.Pool{
	New: func() any {
		return new(FastBuyPosPayResult)
	},
}

// GetFastBuyPosPayResult() 从对象池中获取FastBuyPosPayResult
func GetFastBuyPosPayResult() *FastBuyPosPayResult {
	return poolFastBuyPosPayResult.Get().(*FastBuyPosPayResult)
}

// ReleaseFastBuyPosPayResult 释放FastBuyPosPayResult
func ReleaseFastBuyPosPayResult(v *FastBuyPosPayResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.Success = false
	poolFastBuyPosPayResult.Put(v)
}
