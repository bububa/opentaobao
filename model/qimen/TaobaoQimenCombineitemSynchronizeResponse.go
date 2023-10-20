package qimen

import (
	"sync"
)

// TaobaoQimenCombineitemSynchronizeResponse 结构体
type TaobaoQimenCombineitemSynchronizeResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenCombineitemSynchronizeResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemSynchronizeResponse)
	},
}

// GetTaobaoQimenCombineitemSynchronizeResponse() 从对象池中获取TaobaoQimenCombineitemSynchronizeResponse
func GetTaobaoQimenCombineitemSynchronizeResponse() *TaobaoQimenCombineitemSynchronizeResponse {
	return poolTaobaoQimenCombineitemSynchronizeResponse.Get().(*TaobaoQimenCombineitemSynchronizeResponse)
}

// ReleaseTaobaoQimenCombineitemSynchronizeResponse 释放TaobaoQimenCombineitemSynchronizeResponse
func ReleaseTaobaoQimenCombineitemSynchronizeResponse(v *TaobaoQimenCombineitemSynchronizeResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenCombineitemSynchronizeResponse.Put(v)
}
