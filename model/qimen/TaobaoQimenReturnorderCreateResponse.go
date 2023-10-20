package qimen

import (
	"sync"
)

// TaobaoQimenReturnorderCreateResponse 结构体
type TaobaoQimenReturnorderCreateResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 仓储系统退货单编码
	ReturnOrderId string `json:"returnOrderId,omitempty" xml:"returnOrderId,omitempty"`
}

var poolTaobaoQimenReturnorderCreateResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnorderCreateResponse)
	},
}

// GetTaobaoQimenReturnorderCreateResponse() 从对象池中获取TaobaoQimenReturnorderCreateResponse
func GetTaobaoQimenReturnorderCreateResponse() *TaobaoQimenReturnorderCreateResponse {
	return poolTaobaoQimenReturnorderCreateResponse.Get().(*TaobaoQimenReturnorderCreateResponse)
}

// ReleaseTaobaoQimenReturnorderCreateResponse 释放TaobaoQimenReturnorderCreateResponse
func ReleaseTaobaoQimenReturnorderCreateResponse(v *TaobaoQimenReturnorderCreateResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.ReturnOrderId = ""
	poolTaobaoQimenReturnorderCreateResponse.Put(v)
}
