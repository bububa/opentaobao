package ascp

// ScItemModel 结构体
type ScItemModel struct {
	// 子货品列表(仅对组合货品生效)
	SubScitems []SubScItemModel `json:"sub_scitems,omitempty" xml:"sub_scitems>sub_sc_item_model,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 行业；综合-NORMAL 图书-BOOKS 美妆-BEAUTY
	Industry string `json:"industry,omitempty" xml:"industry,omitempty"`
	// 货主编码
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 货品名称
	ScitemName string `json:"scitem_name,omitempty" xml:"scitem_name,omitempty"`
	// 货品商家编码
	ScitemCode string `json:"scitem_code,omitempty" xml:"scitem_code,omitempty"`
	// 货品id
	ScitemId string `json:"scitem_id,omitempty" xml:"scitem_id,omitempty"`
	// 仓货品编码
	WarehouseScitemCode string `json:"warehouse_scitem_code,omitempty" xml:"warehouse_scitem_code,omitempty"`
	// 条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 货品主图URL
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 存储条件。NORMAL： 常温 COOL： 阴凉 5°C-12°C FRESHNESS：保鲜 0°C-4°C REEFER：冷藏 -18°C-0°C THERMOSTATIC：恒温 18°C-25°C FROZEN：冷冻 &lt;-18°C COOL_AND_DRY：阴凉干燥 5°C-25°C,相对湿度65%以下
	StorageEnvironment string `json:"storage_environment,omitempty" xml:"storage_environment,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 行业属性
	IndustryFeatureMap string `json:"industry_feature_map,omitempty" xml:"industry_feature_map,omitempty"`
	// 长(mm)
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 宽(mm)
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 高(mm)
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 重(g)
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 零售价(人命币-分)
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 是否危险品，1=是，0=否
	Hazardous int64 `json:"hazardous,omitempty" xml:"hazardous,omitempty"`
	// 是否易碎品，1=是，0=否
	Fragile int64 `json:"fragile,omitempty" xml:"fragile,omitempty"`
	// 是否液体，1=是，0=否
	Liquid int64 `json:"liquid,omitempty" xml:"liquid,omitempty"`
	// 是否贵重品，1=是，0=否
	Precious int64 `json:"precious,omitempty" xml:"precious,omitempty"`
	// 采购价列表
	PurchasePrices *PurchasePrice `json:"purchase_prices,omitempty" xml:"purchase_prices,omitempty"`
	// 货品类型；1-普通货品, 2-组合货品
	ScitemType int64 `json:"scitem_type,omitempty" xml:"scitem_type,omitempty"`
}
