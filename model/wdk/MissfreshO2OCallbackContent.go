package wdk

import (
	"sync"
)

// MissfreshO2OCallbackContent 结构体
type MissfreshO2OCallbackContent struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 缺货出库存数量
	OutOfStockStockQuantity string `json:"out_of_stock_stock_quantity,omitempty" xml:"out_of_stock_stock_quantity,omitempty"`
	// 缺货出销售数量
	OutOfStockSaleQuantity string `json:"out_of_stock_sale_quantity,omitempty" xml:"out_of_stock_sale_quantity,omitempty"`
	// 实际库存拣货数量
	ActualStockQuantity string `json:"actual_stock_quantity,omitempty" xml:"actual_stock_quantity,omitempty"`
	// 实际销售拣货数量
	ActualSaleQuantity string `json:"actual_sale_quantity,omitempty" xml:"actual_sale_quantity,omitempty"`
	// 作业内容单号
	WorkUnitContentId string `json:"work_unit_content_id,omitempty" xml:"work_unit_content_id,omitempty"`
	// 是否缺货出
	IsShortage bool `json:"is_shortage,omitempty" xml:"is_shortage,omitempty"`
}

var poolMissfreshO2OCallbackContent = sync.Pool{
	New: func() any {
		return new(MissfreshO2OCallbackContent)
	},
}

// GetMissfreshO2OCallbackContent() 从对象池中获取MissfreshO2OCallbackContent
func GetMissfreshO2OCallbackContent() *MissfreshO2OCallbackContent {
	return poolMissfreshO2OCallbackContent.Get().(*MissfreshO2OCallbackContent)
}

// ReleaseMissfreshO2OCallbackContent 释放MissfreshO2OCallbackContent
func ReleaseMissfreshO2OCallbackContent(v *MissfreshO2OCallbackContent) {
	v.SkuCode = ""
	v.OutOfStockStockQuantity = ""
	v.OutOfStockSaleQuantity = ""
	v.ActualStockQuantity = ""
	v.ActualSaleQuantity = ""
	v.WorkUnitContentId = ""
	v.IsShortage = false
	poolMissfreshO2OCallbackContent.Put(v)
}
