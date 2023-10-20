package wdk

import (
	"sync"
)

// SfB2CFmsCallbackContent 结构体
type SfB2CFmsCallbackContent struct {
	// 快递包裹信息
	Packages []ExpressPackage `json:"packages,omitempty" xml:"packages>express_package,omitempty"`
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
}

var poolSfB2CFmsCallbackContent = sync.Pool{
	New: func() any {
		return new(SfB2CFmsCallbackContent)
	},
}

// GetSfB2CFmsCallbackContent() 从对象池中获取SfB2CFmsCallbackContent
func GetSfB2CFmsCallbackContent() *SfB2CFmsCallbackContent {
	return poolSfB2CFmsCallbackContent.Get().(*SfB2CFmsCallbackContent)
}

// ReleaseSfB2CFmsCallbackContent 释放SfB2CFmsCallbackContent
func ReleaseSfB2CFmsCallbackContent(v *SfB2CFmsCallbackContent) {
	v.Packages = v.Packages[:0]
	v.SkuCode = ""
	v.OutOfStockStockQuantity = ""
	v.OutOfStockSaleQuantity = ""
	v.ActualStockQuantity = ""
	v.ActualSaleQuantity = ""
	v.WorkUnitContentId = ""
	poolSfB2CFmsCallbackContent.Put(v)
}
