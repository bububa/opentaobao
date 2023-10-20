package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderBatchcreateAnswerResponse 结构体
type TaobaoQimenDeliveryorderBatchcreateAnswerResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenDeliveryorderBatchcreateAnswerResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchcreateAnswerResponse)
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateAnswerResponse() 从对象池中获取TaobaoQimenDeliveryorderBatchcreateAnswerResponse
func GetTaobaoQimenDeliveryorderBatchcreateAnswerResponse() *TaobaoQimenDeliveryorderBatchcreateAnswerResponse {
	return poolTaobaoQimenDeliveryorderBatchcreateAnswerResponse.Get().(*TaobaoQimenDeliveryorderBatchcreateAnswerResponse)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerResponse 释放TaobaoQimenDeliveryorderBatchcreateAnswerResponse
func ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerResponse(v *TaobaoQimenDeliveryorderBatchcreateAnswerResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenDeliveryorderBatchcreateAnswerResponse.Put(v)
}
