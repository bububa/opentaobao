package trade

import (
	"sync"
)

// FastBuyPosCloseResult 结构体
type FastBuyPosCloseResult struct {
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 错误信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 关单结果状态: 1为成功,2为失败
	ResultResult int64 `json:"result_result,omitempty" xml:"result_result,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFastBuyPosCloseResult = sync.Pool{
	New: func() any {
		return new(FastBuyPosCloseResult)
	},
}

// GetFastBuyPosCloseResult() 从对象池中获取FastBuyPosCloseResult
func GetFastBuyPosCloseResult() *FastBuyPosCloseResult {
	return poolFastBuyPosCloseResult.Get().(*FastBuyPosCloseResult)
}

// ReleaseFastBuyPosCloseResult 释放FastBuyPosCloseResult
func ReleaseFastBuyPosCloseResult(v *FastBuyPosCloseResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.ResultResult = 0
	v.Success = false
	poolFastBuyPosCloseResult.Put(v)
}
