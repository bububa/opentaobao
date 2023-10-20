package wlbimports

import (
	"sync"
)

// ParcelOrderStatusResponse 结构体
type ParcelOrderStatusResponse struct {
	// 小包LP
	LogisticsOrderCode string `json:"logistics_order_code,omitempty" xml:"logistics_order_code,omitempty"`
	// 状态，init:初始化、inbound_normal：入库正常、inbound_abnormal：入库异常
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolParcelOrderStatusResponse = sync.Pool{
	New: func() any {
		return new(ParcelOrderStatusResponse)
	},
}

// GetParcelOrderStatusResponse() 从对象池中获取ParcelOrderStatusResponse
func GetParcelOrderStatusResponse() *ParcelOrderStatusResponse {
	return poolParcelOrderStatusResponse.Get().(*ParcelOrderStatusResponse)
}

// ReleaseParcelOrderStatusResponse 释放ParcelOrderStatusResponse
func ReleaseParcelOrderStatusResponse(v *ParcelOrderStatusResponse) {
	v.LogisticsOrderCode = ""
	v.Status = ""
	poolParcelOrderStatusResponse.Put(v)
}
