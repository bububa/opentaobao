package nrt

// TopHomeItemDto 结构体
type TopHomeItemDto struct {
	// 类目属性
	CProps []CategoryPropDto `json:"c_props,omitempty" xml:"c_props>category_prop_dto,omitempty"`
	// 图片信息
	Images []ItemImageDto `json:"images,omitempty" xml:"images>item_image_dto,omitempty"`
	// SKU
	Skus []SkuDto `json:"skus,omitempty" xml:"skus>sku_dto,omitempty"`
	// 摊位ID
	BoothId string `json:"booth_id,omitempty" xml:"booth_id,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 经销商编码
	DealerCode string `json:"dealer_code,omitempty" xml:"dealer_code,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// features
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 卖场ID
	MallId string `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 修改时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 叶子类目ID
	CId int64 `json:"c_id,omitempty" xml:"c_id,omitempty"`
	// 系统自动生成
	Ext *MacallineItemExtDto `json:"ext,omitempty" xml:"ext,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品类型，1：主商品，2：摊位商品，3：城市站商品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 主商品ID
	MainItemId int64 `json:"main_item_id,omitempty" xml:"main_item_id,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
