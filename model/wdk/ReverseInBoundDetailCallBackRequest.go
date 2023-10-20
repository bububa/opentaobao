package wdk

import (
	"sync"
)

// ReverseInBoundDetailCallBackRequest 结构体
type ReverseInBoundDetailCallBackRequest struct {
	// 子单扩展字段
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// oms主单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 实际入库销售数量
	ActualSaleQuantity string `json:"actual_sale_quantity,omitempty" xml:"actual_sale_quantity,omitempty"`
	// 实际入库库存数量
	ActualStockQuantity string `json:"actual_stock_quantity,omitempty" xml:"actual_stock_quantity,omitempty"`
	// oms子单号
	BizSubOrderId string `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// sku编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}

var poolReverseInBoundDetailCallBackRequest = sync.Pool{
	New: func() any {
		return new(ReverseInBoundDetailCallBackRequest)
	},
}

// GetReverseInBoundDetailCallBackRequest() 从对象池中获取ReverseInBoundDetailCallBackRequest
func GetReverseInBoundDetailCallBackRequest() *ReverseInBoundDetailCallBackRequest {
	return poolReverseInBoundDetailCallBackRequest.Get().(*ReverseInBoundDetailCallBackRequest)
}

// ReleaseReverseInBoundDetailCallBackRequest 释放ReverseInBoundDetailCallBackRequest
func ReleaseReverseInBoundDetailCallBackRequest(v *ReverseInBoundDetailCallBackRequest) {
	v.Extension = ""
	v.BizOrderId = ""
	v.ActualSaleQuantity = ""
	v.ActualStockQuantity = ""
	v.BizSubOrderId = ""
	v.SkuCode = ""
	poolReverseInBoundDetailCallBackRequest.Put(v)
}
