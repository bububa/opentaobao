package qimen

import (
	"sync"
)

// TaobaoQimenReturnpackageReportRequest 结构体
type TaobaoQimenReturnpackageReportRequest struct {
	// 订单
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
	// 包裹列表
	Packages *Packages `json:"packages,omitempty" xml:"packages,omitempty"`
}

var poolTaobaoQimenReturnpackageReportRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnpackageReportRequest)
	},
}

// GetTaobaoQimenReturnpackageReportRequest() 从对象池中获取TaobaoQimenReturnpackageReportRequest
func GetTaobaoQimenReturnpackageReportRequest() *TaobaoQimenReturnpackageReportRequest {
	return poolTaobaoQimenReturnpackageReportRequest.Get().(*TaobaoQimenReturnpackageReportRequest)
}

// ReleaseTaobaoQimenReturnpackageReportRequest 释放TaobaoQimenReturnpackageReportRequest
func ReleaseTaobaoQimenReturnpackageReportRequest(v *TaobaoQimenReturnpackageReportRequest) {
	v.Order = nil
	v.Packages = nil
	poolTaobaoQimenReturnpackageReportRequest.Put(v)
}
