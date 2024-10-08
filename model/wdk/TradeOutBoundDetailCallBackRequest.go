package wdk

import (
	"sync"
)

// TradeOutBoundDetailCallBackRequest 结构体
type TradeOutBoundDetailCallBackRequest struct {
	// 称重品商品实称重量/g
	SkuWeights []float64 `json:"sku_weights,omitempty" xml:"sku_weights>float64,omitempty"`
	// 缺货出销售数量
	OutOfStockSaleQuantity string `json:"out_of_stock_sale_quantity,omitempty" xml:"out_of_stock_sale_quantity,omitempty"`
	// 渠道子单号
	ChannelSubOrderNo string `json:"channel_sub_order_no,omitempty" xml:"channel_sub_order_no,omitempty"`
	// 扩展字段
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 实际出库销售数量
	ActualSaleQuantity string `json:"actual_sale_quantity,omitempty" xml:"actual_sale_quantity,omitempty"`
	// 实际出库库存数量
	ActualStockQuantity string `json:"actual_stock_quantity,omitempty" xml:"actual_stock_quantity,omitempty"`
	// 缺货出库存数量
	OutOfStockStockQuantity string `json:"out_of_stock_stock_quantity,omitempty" xml:"out_of_stock_stock_quantity,omitempty"`
	// oms子单号
	BizSubOrderId string `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// 商品sku编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 是否缺货出
	IsOutStock bool `json:"is_out_stock,omitempty" xml:"is_out_stock,omitempty"`
}

var poolTradeOutBoundDetailCallBackRequest = sync.Pool{
	New: func() any {
		return new(TradeOutBoundDetailCallBackRequest)
	},
}

// GetTradeOutBoundDetailCallBackRequest() 从对象池中获取TradeOutBoundDetailCallBackRequest
func GetTradeOutBoundDetailCallBackRequest() *TradeOutBoundDetailCallBackRequest {
	return poolTradeOutBoundDetailCallBackRequest.Get().(*TradeOutBoundDetailCallBackRequest)
}

// ReleaseTradeOutBoundDetailCallBackRequest 释放TradeOutBoundDetailCallBackRequest
func ReleaseTradeOutBoundDetailCallBackRequest(v *TradeOutBoundDetailCallBackRequest) {
	v.SkuWeights = v.SkuWeights[:0]
	v.OutOfStockSaleQuantity = ""
	v.ChannelSubOrderNo = ""
	v.Extension = ""
	v.ActualSaleQuantity = ""
	v.ActualStockQuantity = ""
	v.OutOfStockStockQuantity = ""
	v.BizSubOrderId = ""
	v.SkuCode = ""
	v.IsOutStock = false
	poolTradeOutBoundDetailCallBackRequest.Put(v)
}
