package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelMainRefundCreateData 结构体
type AlibabaAscpChannelMainRefundCreateData struct {
	// 外部退款单
	OutRefundNo string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
	// 退款单
	RefundNo string `json:"refund_no,omitempty" xml:"refund_no,omitempty"`
	// 外部子订单
	OutSubOrderNo string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
	// 子订单
	SubSaleOrderNo string `json:"sub_sale_order_no,omitempty" xml:"sub_sale_order_no,omitempty"`
	// 订单
	SaleOrderNo string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}

var poolAlibabaAscpChannelMainRefundCreateData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelMainRefundCreateData)
	},
}

// GetAlibabaAscpChannelMainRefundCreateData() 从对象池中获取AlibabaAscpChannelMainRefundCreateData
func GetAlibabaAscpChannelMainRefundCreateData() *AlibabaAscpChannelMainRefundCreateData {
	return poolAlibabaAscpChannelMainRefundCreateData.Get().(*AlibabaAscpChannelMainRefundCreateData)
}

// ReleaseAlibabaAscpChannelMainRefundCreateData 释放AlibabaAscpChannelMainRefundCreateData
func ReleaseAlibabaAscpChannelMainRefundCreateData(v *AlibabaAscpChannelMainRefundCreateData) {
	v.OutRefundNo = ""
	v.RefundNo = ""
	v.OutSubOrderNo = ""
	v.SubSaleOrderNo = ""
	v.SaleOrderNo = ""
	v.SkuId = 0
	v.ProductId = 0
	v.RefundFee = 0
	poolAlibabaAscpChannelMainRefundCreateData.Put(v)
}
