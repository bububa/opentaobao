package ascpchannel

import (
	"sync"
)

// Orderlines 结构体
type Orderlines struct {
	// 店铺编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 实际成交价
	ActualPrice string `json:"actual_price,omitempty" xml:"actual_price,omitempty"`
	// 零售价=实际成交价+单件商品折扣金额
	RetailPrice string `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商品ID,前端宝贝ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 交易平台子订单编码
	SubSourceOrderCode string `json:"sub_source_order_code,omitempty" xml:"sub_source_order_code,omitempty"`
	// 交易平台订单
	SourceOrderCode string `json:"source_order_code,omitempty" xml:"source_order_code,omitempty"`
	// 单据行
	OrderLineNo string `json:"order_line_no,omitempty" xml:"order_line_no,omitempty"`
	// 应发商品数量
	PlanQty int64 `json:"plan_qty,omitempty" xml:"plan_qty,omitempty"`
}

var poolOrderlines = sync.Pool{
	New: func() any {
		return new(Orderlines)
	},
}

// GetOrderlines() 从对象池中获取Orderlines
func GetOrderlines() *Orderlines {
	return poolOrderlines.Get().(*Orderlines)
}

// ReleaseOrderlines 释放Orderlines
func ReleaseOrderlines(v *Orderlines) {
	v.ShopCode = ""
	v.ActualPrice = ""
	v.RetailPrice = ""
	v.ItemName = ""
	v.ItemCode = ""
	v.ItemId = ""
	v.SubSourceOrderCode = ""
	v.SourceOrderCode = ""
	v.OrderLineNo = ""
	v.PlanQty = 0
	poolOrderlines.Put(v)
}
