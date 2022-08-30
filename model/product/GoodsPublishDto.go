package product

// GoodsPublishDto 结构体
type GoodsPublishDto struct {
	// 卖家账号信息商品属性对象数组
	SellerAccountPropertyList []GoodsPropertyValueDto `json:"seller_account_property_list,omitempty" xml:"seller_account_property_list>goods_property_value_dto,omitempty"`
	// 商品属性对象数组
	GoodsPropertyList []GoodsPropertyValueDto `json:"goods_property_list,omitempty" xml:"goods_property_list>goods_property_value_dto,omitempty"`
	// 商品图片url列表
	ImageUrlList []GoodsPublishImageDto `json:"image_url_list,omitempty" xml:"image_url_list>goods_publish_image_dto,omitempty"`
	// 外部商品ID，用于标识外部系统每次提交过来的商品
	ExternalGoodsId string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
	// 游戏属性对象
	GameProperty *GamePropertyDto `json:"game_property,omitempty" xml:"game_property,omitempty"`
	// 二级类目ID
	SecondCategoryId int64 `json:"second_category_id,omitempty" xml:"second_category_id,omitempty"`
	// 一级类目ID
	FirstCategoryId int64 `json:"first_category_id,omitempty" xml:"first_category_id,omitempty"`
	// 商品基本信息
	GoodsBaseInfo *GoodsBaseInfoDto `json:"goods_base_info,omitempty" xml:"goods_base_info,omitempty"`
	// 是否支持找回包赔
	SupportRetrieveCompensation bool `json:"support_retrieve_compensation,omitempty" xml:"support_retrieve_compensation,omitempty"`
	// 是否支持议价
	CanBargain bool `json:"can_bargain,omitempty" xml:"can_bargain,omitempty"`
}
