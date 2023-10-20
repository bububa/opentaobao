package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelSubRefundCreateData 结构体
type AlibabaAscpChannelSubRefundCreateData struct {
	// 外部退款单号
	OutRefundNo string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
	// 供应链渠道退款单号
	RefundNo string `json:"refund_no,omitempty" xml:"refund_no,omitempty"`
	// 外部子订单号
	OutSubOrderNo string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
	// 子订单号
	SubSaleOrderNo string `json:"sub_sale_order_no,omitempty" xml:"sub_sale_order_no,omitempty"`
	// 供应链渠道订单号
	SaleOrderNo string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolAlibabaAscpChannelSubRefundCreateData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSubRefundCreateData)
	},
}

// GetAlibabaAscpChannelSubRefundCreateData() 从对象池中获取AlibabaAscpChannelSubRefundCreateData
func GetAlibabaAscpChannelSubRefundCreateData() *AlibabaAscpChannelSubRefundCreateData {
	return poolAlibabaAscpChannelSubRefundCreateData.Get().(*AlibabaAscpChannelSubRefundCreateData)
}

// ReleaseAlibabaAscpChannelSubRefundCreateData 释放AlibabaAscpChannelSubRefundCreateData
func ReleaseAlibabaAscpChannelSubRefundCreateData(v *AlibabaAscpChannelSubRefundCreateData) {
	v.OutRefundNo = ""
	v.RefundNo = ""
	v.OutSubOrderNo = ""
	v.SubSaleOrderNo = ""
	v.SaleOrderNo = ""
	v.RefundFee = 0
	v.SkuId = 0
	v.ProductId = 0
	poolAlibabaAscpChannelSubRefundCreateData.Put(v)
}
