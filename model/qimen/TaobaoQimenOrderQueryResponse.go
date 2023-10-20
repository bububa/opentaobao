package qimen

import (
	"sync"
)

// TaobaoQimenOrderQueryResponse 结构体
type TaobaoQimenOrderQueryResponse struct {
	// success|failure，必填
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单列表
	OrderLines *OrderLines `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
}

var poolTaobaoQimenOrderQueryResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderQueryResponse)
	},
}

// GetTaobaoQimenOrderQueryResponse() 从对象池中获取TaobaoQimenOrderQueryResponse
func GetTaobaoQimenOrderQueryResponse() *TaobaoQimenOrderQueryResponse {
	return poolTaobaoQimenOrderQueryResponse.Get().(*TaobaoQimenOrderQueryResponse)
}

// ReleaseTaobaoQimenOrderQueryResponse 释放TaobaoQimenOrderQueryResponse
func ReleaseTaobaoQimenOrderQueryResponse(v *TaobaoQimenOrderQueryResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.OrderLines = nil
	poolTaobaoQimenOrderQueryResponse.Put(v)
}
