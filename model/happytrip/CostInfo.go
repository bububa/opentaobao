package happytrip

import (
	"sync"
)

// CostInfo 结构体
type CostInfo struct {
	// 费用明细
	Detail []CostDetailInfo `json:"detail,omitempty" xml:"detail>cost_detail_info,omitempty"`
	// 总费用，折后金额
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 原价，如果订单有折扣这里为折扣前的价格，如果没有折扣和total_price字段保持一致
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
}

var poolCostInfo = sync.Pool{
	New: func() any {
		return new(CostInfo)
	},
}

// GetCostInfo() 从对象池中获取CostInfo
func GetCostInfo() *CostInfo {
	return poolCostInfo.Get().(*CostInfo)
}

// ReleaseCostInfo 释放CostInfo
func ReleaseCostInfo(v *CostInfo) {
	v.Detail = v.Detail[:0]
	v.TotalPrice = ""
	v.OriginalPrice = ""
	poolCostInfo.Put(v)
}
