package aliospay

import (
	"sync"
)

// QueryTradeRequest 结构体
type QueryTradeRequest struct {
	// 请求唯一id，不可重复，服务端会根据此参数防重放
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 语言,en表示英文，zh表示中文
	Lang string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 发送请求的时间戳
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// alios支付订单id
	PayOrderId string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
}

var poolQueryTradeRequest = sync.Pool{
	New: func() any {
		return new(QueryTradeRequest)
	},
}

// GetQueryTradeRequest() 从对象池中获取QueryTradeRequest
func GetQueryTradeRequest() *QueryTradeRequest {
	return poolQueryTradeRequest.Get().(*QueryTradeRequest)
}

// ReleaseQueryTradeRequest 释放QueryTradeRequest
func ReleaseQueryTradeRequest(v *QueryTradeRequest) {
	v.TraceId = ""
	v.Lang = ""
	v.Time = ""
	v.BizOrderId = ""
	v.PayOrderId = ""
	poolQueryTradeRequest.Put(v)
}
