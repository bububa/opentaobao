package wdk

// SfB2CFmsCallbackContent 结构体
type SfB2CFmsCallbackContent struct {
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
	// 快递包裹信息
	Packages []ExpressPackage `json:"packages,omitempty" xml:"packages>express_package,omitempty"`
}
