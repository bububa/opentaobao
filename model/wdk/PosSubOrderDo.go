package wdk

import (
	"sync"
)

// PosSubOrderDo 结构体
type PosSubOrderDo struct {
	// 库存单位，必填
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 库存单位购买数量，必填
	BuyAmountStock string `json:"buy_amount_stock,omitempty" xml:"buy_amount_stock,omitempty"`
	// sku编码，必填
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 外部子订单号，全局唯一，子单和主单不能重复，可以包含字母
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 子单实付金额，单位分
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 销售单位购买数量。对于标品，和库存单位库存单位购买数量一样
	BuyAmountSale int64 `json:"buy_amount_sale,omitempty" xml:"buy_amount_sale,omitempty"`
	// 子单原价金额，单位分
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 子单优惠金额，单位分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 商品单价，单位分
	SkuPrice int64 `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
}

var poolPosSubOrderDo = sync.Pool{
	New: func() any {
		return new(PosSubOrderDo)
	},
}

// GetPosSubOrderDo() 从对象池中获取PosSubOrderDo
func GetPosSubOrderDo() *PosSubOrderDo {
	return poolPosSubOrderDo.Get().(*PosSubOrderDo)
}

// ReleasePosSubOrderDo 释放PosSubOrderDo
func ReleasePosSubOrderDo(v *PosSubOrderDo) {
	v.StockUnit = ""
	v.BuyAmountStock = ""
	v.SkuCode = ""
	v.OutOrderId = ""
	v.SaleUnit = ""
	v.SkuName = ""
	v.PayFee = 0
	v.BuyAmountSale = 0
	v.OriginFee = 0
	v.DiscountFee = 0
	v.SkuPrice = 0
	poolPosSubOrderDo.Put(v)
}
