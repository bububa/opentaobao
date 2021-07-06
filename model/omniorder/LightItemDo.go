package omniorder

// LightItemDo 结构体
type LightItemDo struct {
	// 关联门店id
	StoreIds []int64 `json:"store_ids,omitempty" xml:"store_ids>int64,omitempty"`
	// sku信息列表
	Skus []ItemLightPublishSkuDto `json:"skus,omitempty" xml:"skus>item_light_publish_sku_dto,omitempty"`
	// 商品图片，最少1张，最多5张
	Images []ItemLightPublishImageDto `json:"images,omitempty" xml:"images>item_light_publish_image_dto,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品卖点
	Subtitle string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// 吊牌价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 销售价
	Pretium string `json:"pretium,omitempty" xml:"pretium,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 商品条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 扩展字段
	ExtendAttr string `json:"extend_attr,omitempty" xml:"extend_attr,omitempty"`
	// 商品所属商家id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 商品淘系id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品叶子类目，参见taobao.omniitem.category.get接口返回值
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}
