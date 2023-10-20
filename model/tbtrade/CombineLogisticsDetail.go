package tbtrade

import (
	"sync"
)

// CombineLogisticsDetail 结构体
type CombineLogisticsDetail struct {
	// 包裹详情（仅支持成分品）
	SendGoodsDetails []SendGoodsDetail `json:"send_goods_details,omitempty" xml:"send_goods_details>send_goods_detail,omitempty"`
	// 运单号
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 物流公司
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 子订单id
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

var poolCombineLogisticsDetail = sync.Pool{
	New: func() any {
		return new(CombineLogisticsDetail)
	},
}

// GetCombineLogisticsDetail() 从对象池中获取CombineLogisticsDetail
func GetCombineLogisticsDetail() *CombineLogisticsDetail {
	return poolCombineLogisticsDetail.Get().(*CombineLogisticsDetail)
}

// ReleaseCombineLogisticsDetail 释放CombineLogisticsDetail
func ReleaseCombineLogisticsDetail(v *CombineLogisticsDetail) {
	v.SendGoodsDetails = v.SendGoodsDetails[:0]
	v.InvoiceNo = ""
	v.LogisticsCompany = ""
	v.SubOrderId = ""
	poolCombineLogisticsDetail.Put(v)
}
