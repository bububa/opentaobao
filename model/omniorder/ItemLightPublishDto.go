package omniorder

// ItemLightPublishDto 结构体
type ItemLightPublishDto struct {
	// 商品条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 叶子类目ID
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// extendAttr
	ExtendAttr string `json:"extend_attr,omitempty" xml:"extend_attr,omitempty"`
	// images
	Images []ItemLightPublishImageDto `json:"images,omitempty" xml:"images>item_light_publish_image_dto,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 销售价
	Pretium string `json:"pretium,omitempty" xml:"pretium,omitempty"`
	// 吊牌价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// skus
	Skus []ItemLightPublishSkuDto `json:"skus,omitempty" xml:"skus>item_light_publish_sku_dto,omitempty"`
	// 副标题
	Subtitle string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 卖家ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 商品描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
}
