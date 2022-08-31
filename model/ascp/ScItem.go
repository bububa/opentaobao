package ascp

// ScItem 结构体
type ScItem struct {
	// 仓库货品编码
	WarehouseScItemRelation []WarehouseScItemRelation `json:"warehouse_sc_item_relation,omitempty" xml:"warehouse_sc_item_relation>warehouse_sc_item_relation,omitempty"`
	// 采购价格
	PurchasePrices []PurchasePrice `json:"purchase_prices,omitempty" xml:"purchase_prices>purchase_price,omitempty"`
	// ERP货品id（支持组合货品、普通货品）
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 货品商家编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 货品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 扩展字段
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 行业。综合-NORMAL 图书-BOOKS 美妆-BEAUTY
	Industry string `json:"industry,omitempty" xml:"industry,omitempty"`
	// 货主编码；优仓分配，长度不会超过32位，不含特殊字符
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 仓库编码。需要下发仓库时，必填
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 货品主图。填写URL
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 存储条件。NORMAL： 常温 COOL： 阴凉 5°C-12°C FRESHNESS：保鲜 0°C-4°C REEFER：冷藏 -18°C-0°C THERMOSTATIC：恒温 18°C-25°C FROZEN：冷冻 &lt;-18°C COOL_AND_DRY：阴凉干燥 5°C-25°C,相对湿度65%以下
	StorageEnvironment string `json:"storage_environment,omitempty" xml:"storage_environment,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 币种，USD-美元，CNY-人民币，RUB-卢布，JPY-日元，EUR-欧元，GBP-英镑，HKD-港币，NZD-新西兰元，SGD-新加坡元，AUD-澳元，KRW-韩元，THB-泰铢
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 仓库货品编码
	WarehouseScItemCode string `json:"warehouse_sc_item_code,omitempty" xml:"warehouse_sc_item_code,omitempty"`
	// 长(毫米)
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 宽(毫米)
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 高(毫米)
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 重量(毫克)
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 零售价，单位:分，如:20007，表示:20元7分
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 是否危险品，1=是，0=否
	Hazardous int64 `json:"hazardous,omitempty" xml:"hazardous,omitempty"`
	// 是否易碎品，1=是，0=否
	Fragile int64 `json:"fragile,omitempty" xml:"fragile,omitempty"`
	// 是否液体，1=是，0=否
	Liquid int64 `json:"liquid,omitempty" xml:"liquid,omitempty"`
	// 是否贵重品，1=是，0=否
	Precious int64 `json:"precious,omitempty" xml:"precious,omitempty"`
	// 是否需要下发仓。0表示否，1表示是
	NeedNotifyWarehouse int64 `json:"need_notify_warehouse,omitempty" xml:"need_notify_warehouse,omitempty"`
}
