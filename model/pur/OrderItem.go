package pur

import (
	"sync"
)

// OrderItem 结构体
type OrderItem struct {
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 税率，如13%传13
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 商城品类
	MallCategoryId string `json:"mall_category_id,omitempty" xml:"mall_category_id,omitempty"`
	// 元，个
	Uom string `json:"uom,omitempty" xml:"uom,omitempty"`
	// 子订单号
	SubPurReqId string `json:"sub_pur_req_id,omitempty" xml:"sub_pur_req_id,omitempty"`
	// sku的ID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 单价(含税）
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 合同ID
	ContractId string `json:"contract_id,omitempty" xml:"contract_id,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品描述
	ItemDescription string `json:"item_description,omitempty" xml:"item_description,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolOrderItem = sync.Pool{
	New: func() any {
		return new(OrderItem)
	},
}

// GetOrderItem() 从对象池中获取OrderItem
func GetOrderItem() *OrderItem {
	return poolOrderItem.Get().(*OrderItem)
}

// ReleaseOrderItem 释放OrderItem
func ReleaseOrderItem(v *OrderItem) {
	v.CurrencyCode = ""
	v.TaxRate = ""
	v.SupplierId = ""
	v.MallCategoryId = ""
	v.Uom = ""
	v.SubPurReqId = ""
	v.SkuId = ""
	v.ItemName = ""
	v.UnitPrice = ""
	v.ContractId = ""
	v.ItemId = ""
	v.ItemDescription = ""
	v.Quantity = 0
	poolOrderItem.Put(v)
}
