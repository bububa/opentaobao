package uscesl

// ItemChangeBo 结构体
type ItemChangeBo struct {
	// 溯源码URL
	SourceCode string `json:"source_code,omitempty" xml:"source_code,omitempty"`
	// 销售单位
	PriceUnit string `json:"price_unit,omitempty" xml:"price_unit,omitempty"`
	// 品牌
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品类
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 等级
	Rank string `json:"rank,omitempty" xml:"rank,omitempty"`
	// (建议)零售价，单位分，整数
	SuggestPrice string `json:"suggest_price,omitempty" xml:"suggest_price,omitempty"`
	// 库存sku id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 能效
	EnergyEfficiency string `json:"energy_efficiency,omitempty" xml:"energy_efficiency,omitempty"`
	// 优惠开始时间
	PromotionStart string `json:"promotion_start,omitempty" xml:"promotion_start,omitempty"`
	// 淘系二级商品类目ID
	ForestSecondCatId string `json:"forest_second_cat_id,omitempty" xml:"forest_second_cat_id,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureC string `json:"customize_feature_c,omitempty" xml:"customize_feature_c,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureD string `json:"customize_feature_d,omitempty" xml:"customize_feature_d,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureE string `json:"customize_feature_e,omitempty" xml:"customize_feature_e,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureF string `json:"customize_feature_f,omitempty" xml:"customize_feature_f,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureG string `json:"customize_feature_g,omitempty" xml:"customize_feature_g,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureH string `json:"customize_feature_h,omitempty" xml:"customize_feature_h,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureI string `json:"customize_feature_i,omitempty" xml:"customize_feature_i,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureJ string `json:"customize_feature_j,omitempty" xml:"customize_feature_j,omitempty"`
	// 二维码图片URL
	ItemQrCode string `json:"item_qr_code,omitempty" xml:"item_qr_code,omitempty"`
	// 是否优惠
	IfPromotion bool `json:"if_promotion,omitempty" xml:"if_promotion,omitempty"`
	// 优惠结束时间
	PromotionEnd string `json:"promotion_end,omitempty" xml:"promotion_end,omitempty"`
	// 额外扩展信息
	ExtraAttribute string `json:"extra_attribute,omitempty" xml:"extra_attribute,omitempty"`
	// 原价，分，整数
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 货架位
	PositonCode string `json:"positon_code,omitempty" xml:"positon_code,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureB string `json:"customize_feature_b,omitempty" xml:"customize_feature_b,omitempty"`
	// 自定义扩展属性
	CustomizeFeatureA string `json:"customize_feature_a,omitempty" xml:"customize_feature_a,omitempty"`
	// 型号
	ModelNum string `json:"model_num,omitempty" xml:"model_num,omitempty"`
	// 规格
	SaleSpec string `json:"sale_spec,omitempty" xml:"sale_spec,omitempty"`
	// 实际销售价格，分，整数
	AcctionPrice string `json:"acction_price,omitempty" xml:"acction_price,omitempty"`
	// 商品条码
	ItemBarCode string `json:"item_bar_code,omitempty" xml:"item_bar_code,omitempty"`
	// 会员价格，分，整数
	MemberPrice string `json:"member_price,omitempty" xml:"member_price,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 优惠文案
	PromotionText string `json:"promotion_text,omitempty" xml:"promotion_text,omitempty"`
	// 商品状态
	ItemStatus int64 `json:"item_status,omitempty" xml:"item_status,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 一级商品类目id
	ForestFirstCatId string `json:"forest_first_cat_id,omitempty" xml:"forest_first_cat_id,omitempty"`
	// 商品短标题
	ShortTitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 是否溯源
	IfSourceCode bool `json:"if_source_code,omitempty" xml:"if_source_code,omitempty"`
	// 产地
	ProductionPlace string `json:"production_place,omitempty" xml:"production_place,omitempty"`
	// 商品变更状态
	ItemChangeStatus string `json:"item_change_status,omitempty" xml:"item_change_status,omitempty"`
	// 促销原因
	PromotionReason string `json:"promotion_reason,omitempty" xml:"promotion_reason,omitempty"`
}
