package wdk

import (
	"sync"
)

// LogisticsTraceCallbackRequest 结构体
type LogisticsTraceCallbackRequest struct {
	// 经营店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 维度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 订单编码
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolLogisticsTraceCallbackRequest = sync.Pool{
	New: func() any {
		return new(LogisticsTraceCallbackRequest)
	},
}

// GetLogisticsTraceCallbackRequest() 从对象池中获取LogisticsTraceCallbackRequest
func GetLogisticsTraceCallbackRequest() *LogisticsTraceCallbackRequest {
	return poolLogisticsTraceCallbackRequest.Get().(*LogisticsTraceCallbackRequest)
}

// ReleaseLogisticsTraceCallbackRequest 释放LogisticsTraceCallbackRequest
func ReleaseLogisticsTraceCallbackRequest(v *LogisticsTraceCallbackRequest) {
	v.StoreId = ""
	v.Longitude = ""
	v.Latitude = ""
	v.UpdateTime = ""
	v.BizOrderId = 0
	poolLogisticsTraceCallbackRequest.Put(v)
}
