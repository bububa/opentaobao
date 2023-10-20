package qimen

import (
	"sync"
)

// TaobaoQimenReceiverinfoQueryResponse 结构体
type TaobaoQimenReceiverinfoQueryResponse struct {
	// success|failure，必填
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单收件人 ID, string (50)
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 出库单号, string (50) , 必填
	DeliveryOrderCode string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
	// 收货人信息
	ReceiverInfo *ReceiverInfo `json:"receiverInfo,omitempty" xml:"receiverInfo,omitempty"`
}

var poolTaobaoQimenReceiverinfoQueryResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReceiverinfoQueryResponse)
	},
}

// GetTaobaoQimenReceiverinfoQueryResponse() 从对象池中获取TaobaoQimenReceiverinfoQueryResponse
func GetTaobaoQimenReceiverinfoQueryResponse() *TaobaoQimenReceiverinfoQueryResponse {
	return poolTaobaoQimenReceiverinfoQueryResponse.Get().(*TaobaoQimenReceiverinfoQueryResponse)
}

// ReleaseTaobaoQimenReceiverinfoQueryResponse 释放TaobaoQimenReceiverinfoQueryResponse
func ReleaseTaobaoQimenReceiverinfoQueryResponse(v *TaobaoQimenReceiverinfoQueryResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.Oaid = ""
	v.DeliveryOrderCode = ""
	v.ReceiverInfo = nil
	poolTaobaoQimenReceiverinfoQueryResponse.Put(v)
}
