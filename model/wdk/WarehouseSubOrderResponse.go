package wdk

import (
	"sync"
)

// WarehouseSubOrderResponse 结构体
type WarehouseSubOrderResponse struct {
	// 渠道子单号
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 库存单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 交易子单号
	BizSubOrderId int64 `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// 库存单位购买数量(应拣数量)
	BuyStockQuantity float64 `json:"buy_stock_quantity,omitempty" xml:"buy_stock_quantity,omitempty"`
	// 库存单位拣货数量
	PickStockQuantity float64 `json:"pick_stock_quantity,omitempty" xml:"pick_stock_quantity,omitempty"`
}

var poolWarehouseSubOrderResponse = sync.Pool{
	New: func() any {
		return new(WarehouseSubOrderResponse)
	},
}

// GetWarehouseSubOrderResponse() 从对象池中获取WarehouseSubOrderResponse
func GetWarehouseSubOrderResponse() *WarehouseSubOrderResponse {
	return poolWarehouseSubOrderResponse.Get().(*WarehouseSubOrderResponse)
}

// ReleaseWarehouseSubOrderResponse 释放WarehouseSubOrderResponse
func ReleaseWarehouseSubOrderResponse(v *WarehouseSubOrderResponse) {
	v.OutSubOrderId = ""
	v.SkuCode = ""
	v.SkuName = ""
	v.StockUnit = ""
	v.BizSubOrderId = 0
	v.BuyStockQuantity = 0
	v.PickStockQuantity = 0
	poolWarehouseSubOrderResponse.Put(v)
}
