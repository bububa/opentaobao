package qimen

import (
	"sync"
)

// TaobaoQimenReturnorderConfirmResponse 结构体
type TaobaoQimenReturnorderConfirmResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenReturnorderConfirmResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnorderConfirmResponse)
	},
}

// GetTaobaoQimenReturnorderConfirmResponse() 从对象池中获取TaobaoQimenReturnorderConfirmResponse
func GetTaobaoQimenReturnorderConfirmResponse() *TaobaoQimenReturnorderConfirmResponse {
	return poolTaobaoQimenReturnorderConfirmResponse.Get().(*TaobaoQimenReturnorderConfirmResponse)
}

// ReleaseTaobaoQimenReturnorderConfirmResponse 释放TaobaoQimenReturnorderConfirmResponse
func ReleaseTaobaoQimenReturnorderConfirmResponse(v *TaobaoQimenReturnorderConfirmResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenReturnorderConfirmResponse.Put(v)
}
