package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderBatchconfirmResponse 结构体
type TaobaoQimenDeliveryorderBatchconfirmResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenDeliveryorderBatchconfirmResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchconfirmResponse)
	},
}

// GetTaobaoQimenDeliveryorderBatchconfirmResponse() 从对象池中获取TaobaoQimenDeliveryorderBatchconfirmResponse
func GetTaobaoQimenDeliveryorderBatchconfirmResponse() *TaobaoQimenDeliveryorderBatchconfirmResponse {
	return poolTaobaoQimenDeliveryorderBatchconfirmResponse.Get().(*TaobaoQimenDeliveryorderBatchconfirmResponse)
}

// ReleaseTaobaoQimenDeliveryorderBatchconfirmResponse 释放TaobaoQimenDeliveryorderBatchconfirmResponse
func ReleaseTaobaoQimenDeliveryorderBatchconfirmResponse(v *TaobaoQimenDeliveryorderBatchconfirmResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenDeliveryorderBatchconfirmResponse.Put(v)
}
