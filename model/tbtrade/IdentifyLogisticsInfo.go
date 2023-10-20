package tbtrade

import (
	"sync"
)

// IdentifyLogisticsInfo 结构体
type IdentifyLogisticsInfo struct {
	// 物流公司
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 运单号
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 退款单号
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 子订单号
	DetailOrderId int64 `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
	// 阶段号
	StageNo int64 `json:"stage_no,omitempty" xml:"stage_no,omitempty"`
	// 是否退货
	Refund bool `json:"refund,omitempty" xml:"refund,omitempty"`
}

var poolIdentifyLogisticsInfo = sync.Pool{
	New: func() any {
		return new(IdentifyLogisticsInfo)
	},
}

// GetIdentifyLogisticsInfo() 从对象池中获取IdentifyLogisticsInfo
func GetIdentifyLogisticsInfo() *IdentifyLogisticsInfo {
	return poolIdentifyLogisticsInfo.Get().(*IdentifyLogisticsInfo)
}

// ReleaseIdentifyLogisticsInfo 释放IdentifyLogisticsInfo
func ReleaseIdentifyLogisticsInfo(v *IdentifyLogisticsInfo) {
	v.LogisticsCompany = ""
	v.InvoiceNo = ""
	v.RefundId = ""
	v.DetailOrderId = 0
	v.StageNo = 0
	v.Refund = false
	poolIdentifyLogisticsInfo.Put(v)
}
