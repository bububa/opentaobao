package qimen

import (
	"sync"
)

// TaobaoQimenEntryorderConfirmResponse 结构体
type TaobaoQimenEntryorderConfirmResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenEntryorderConfirmResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderConfirmResponse)
	},
}

// GetTaobaoQimenEntryorderConfirmResponse() 从对象池中获取TaobaoQimenEntryorderConfirmResponse
func GetTaobaoQimenEntryorderConfirmResponse() *TaobaoQimenEntryorderConfirmResponse {
	return poolTaobaoQimenEntryorderConfirmResponse.Get().(*TaobaoQimenEntryorderConfirmResponse)
}

// ReleaseTaobaoQimenEntryorderConfirmResponse 释放TaobaoQimenEntryorderConfirmResponse
func ReleaseTaobaoQimenEntryorderConfirmResponse(v *TaobaoQimenEntryorderConfirmResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenEntryorderConfirmResponse.Put(v)
}
