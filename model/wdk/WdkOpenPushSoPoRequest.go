package wdk

import (
	"sync"
)

// WdkOpenPushSoPoRequest 结构体
type WdkOpenPushSoPoRequest struct {
	// 淘系子订单列表，一次最多20个
	SubTradeOrderNoList []string `json:"sub_trade_order_no_list,omitempty" xml:"sub_trade_order_no_list>string,omitempty"`
	// 销售=10，销退=20
	SalesType int64 `json:"sales_type,omitempty" xml:"sales_type,omitempty"`
}

var poolWdkOpenPushSoPoRequest = sync.Pool{
	New: func() any {
		return new(WdkOpenPushSoPoRequest)
	},
}

// GetWdkOpenPushSoPoRequest() 从对象池中获取WdkOpenPushSoPoRequest
func GetWdkOpenPushSoPoRequest() *WdkOpenPushSoPoRequest {
	return poolWdkOpenPushSoPoRequest.Get().(*WdkOpenPushSoPoRequest)
}

// ReleaseWdkOpenPushSoPoRequest 释放WdkOpenPushSoPoRequest
func ReleaseWdkOpenPushSoPoRequest(v *WdkOpenPushSoPoRequest) {
	v.SubTradeOrderNoList = v.SubTradeOrderNoList[:0]
	v.SalesType = 0
	poolWdkOpenPushSoPoRequest.Put(v)
}
