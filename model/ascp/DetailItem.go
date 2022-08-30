package ascp

// DetailItem 结构体
type DetailItem struct {
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 货主编码
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 商家货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 组合货品商家编码
	CombineScItemCode string `json:"combine_sc_item_code,omitempty" xml:"combine_sc_item_code,omitempty"`
	// 组合货品翱象Id
	CombineScItemId string `json:"combine_sc_item_id,omitempty" xml:"combine_sc_item_id,omitempty"`
	// erp配资源唯一编码，卖家唯一
	ErpCode string `json:"erp_code,omitempty" xml:"erp_code,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// sku id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 翱象货品ID
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 货品商家条码
	ScItemBarCode string `json:"sc_item_bar_code,omitempty" xml:"sc_item_bar_code,omitempty"`
	// 卖家ID
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
