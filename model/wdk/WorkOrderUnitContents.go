package wdk

// WorkOrderUnitContents 结构体
type WorkOrderUnitContents struct {
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 库存单位购买库数量
	ExpectStockQuantity string `json:"expect_stock_quantity,omitempty" xml:"expect_stock_quantity,omitempty"`
	// 扩展服务
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
	// 销售单位购买数量
	ExpectSaleQuantity string `json:"expect_sale_quantity,omitempty" xml:"expect_sale_quantity,omitempty"`
	// 履约子单号/子订单号
	WorkUnitContentId string `json:"work_unit_content_id,omitempty" xml:"work_unit_content_id,omitempty"`
	// expect_sale_quantity单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// expect_stock_quantity单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 子订单号
	OrderSubCode string `json:"order_sub_code,omitempty" xml:"order_sub_code,omitempty"`
}
