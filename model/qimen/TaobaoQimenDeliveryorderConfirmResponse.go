package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderConfirmResponse 结构体
type TaobaoQimenDeliveryorderConfirmResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenDeliveryorderConfirmResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderConfirmResponse)
	},
}

// GetTaobaoQimenDeliveryorderConfirmResponse() 从对象池中获取TaobaoQimenDeliveryorderConfirmResponse
func GetTaobaoQimenDeliveryorderConfirmResponse() *TaobaoQimenDeliveryorderConfirmResponse {
	return poolTaobaoQimenDeliveryorderConfirmResponse.Get().(*TaobaoQimenDeliveryorderConfirmResponse)
}

// ReleaseTaobaoQimenDeliveryorderConfirmResponse 释放TaobaoQimenDeliveryorderConfirmResponse
func ReleaseTaobaoQimenDeliveryorderConfirmResponse(v *TaobaoQimenDeliveryorderConfirmResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenDeliveryorderConfirmResponse.Put(v)
}
