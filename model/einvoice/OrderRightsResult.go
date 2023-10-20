package einvoice

import (
	"sync"
)

// OrderRightsResult 结构体
type OrderRightsResult struct {
	// 赔付列表
	PayoutList []OrderRightsInfo `json:"payout_list,omitempty" xml:"payout_list>order_rights_info,omitempty"`
	// 赔付列表总数，超过系统最大展示数量是以200+的格式返回，
	TotalCount string `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolOrderRightsResult = sync.Pool{
	New: func() any {
		return new(OrderRightsResult)
	},
}

// GetOrderRightsResult() 从对象池中获取OrderRightsResult
func GetOrderRightsResult() *OrderRightsResult {
	return poolOrderRightsResult.Get().(*OrderRightsResult)
}

// ReleaseOrderRightsResult 释放OrderRightsResult
func ReleaseOrderRightsResult(v *OrderRightsResult) {
	v.PayoutList = v.PayoutList[:0]
	v.TotalCount = ""
	poolOrderRightsResult.Put(v)
}
