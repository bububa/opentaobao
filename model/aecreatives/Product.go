package aecreatives

// Product 结构体
type Product struct {
	// app专享售价
	AppSalePrice string `json:"app_sale_price,omitempty" xml:"app_sale_price,omitempty"`
	// 佣金率
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 折扣比例
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 一级类目ID
	FirstLevelCategoryId int64 `json:"first_level_category_id,omitempty" xml:"first_level_category_id,omitempty"`
	// 一级类目名称
	FirstLevelCategoryName string `json:"first_level_category_name,omitempty" xml:"first_level_category_name,omitempty"`
	// 商品原始报价
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 商品访问地址
	ProductDetailUrl string `json:"product_detail_url,omitempty" xml:"product_detail_url,omitempty"`
	// 商品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 商品主图地址
	ProductMainImageUrl string `json:"product_main_image_url,omitempty" xml:"product_main_image_url,omitempty"`
	// 商品小图地址列表
	ProductSmallImageUrls []string `json:"product_small_image_urls,omitempty" xml:"product_small_image_urls>string,omitempty"`
	// 商品标题
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// 商品视频地址
	ProductVideoUrl string `json:"product_video_url,omitempty" xml:"product_video_url,omitempty"`
	// 推广链接
	PromotionLink string `json:"promotion_link,omitempty" xml:"promotion_link,omitempty"`
	// 商品售价
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 商品二级类目ID
	SecondLevelCategoryId int64 `json:"second_level_category_id,omitempty" xml:"second_level_category_id,omitempty"`
	// 商品二级类目名称
	SecondLevelCategoryName string `json:"second_level_category_name,omitempty" xml:"second_level_category_name,omitempty"`
	// 店铺ID
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 店铺地址
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// 按目标币种转换后的APP专享价
	TargetAppSalePrice string `json:"target_app_sale_price,omitempty" xml:"target_app_sale_price,omitempty"`
	// 按目标币种转换后的原始价格
	TargetOriginalPrice string `json:"target_original_price,omitempty" xml:"target_original_price,omitempty"`
	// 按目标币种转换后的通用售价币种
	TargetSalePrice string `json:"target_sale_price,omitempty" xml:"target_sale_price,omitempty"`
	// 最近销量
	LastestVolume int64 `json:"lastest_volume,omitempty" xml:"lastest_volume,omitempty"`
	// 爆品佣金率
	HotProductCommissionRate string `json:"hot_product_commission_rate,omitempty" xml:"hot_product_commission_rate,omitempty"`
	// 平台商品类型：ALL,PLAZA,TMALL
	PlatformProductType string `json:"platform_product_type,omitempty" xml:"platform_product_type,omitempty"`
	// 好评率
	EvaluateRate string `json:"evaluate_rate,omitempty" xml:"evaluate_rate,omitempty"`
	// app专享售价币种
	AppSalePriceCurrency string `json:"app_sale_price_currency,omitempty" xml:"app_sale_price_currency,omitempty"`
	// 原始售价币种
	OriginalPriceCurrency string `json:"original_price_currency,omitempty" xml:"original_price_currency,omitempty"`
	// 按目标币种转换后的原始价格币种
	TargetOriginalPriceCurrency string `json:"target_original_price_currency,omitempty" xml:"target_original_price_currency,omitempty"`
	// 按目标币种转换后的售价
	TargetSalePriceCurrency string `json:"target_sale_price_currency,omitempty" xml:"target_sale_price_currency,omitempty"`
	// 按目标币种转换后的APP专享价币种
	TargetAppSalePriceCurrency string `json:"target_app_sale_price_currency,omitempty" xml:"target_app_sale_price_currency,omitempty"`
	// 商品售价币种
	SalePriceCurrency string `json:"sale_price_currency,omitempty" xml:"sale_price_currency,omitempty"`
	// JV佣金率
	RelevantMarketCommissionRate string `json:"relevant_market_commission_rate,omitempty" xml:"relevant_market_commission_rate,omitempty"`
	// code信息
	PromoCodeInfo *PromoCodeDto `json:"promo_code_info,omitempty" xml:"promo_code_info,omitempty"`
	// 可达国家与到达时间
	ShipToDays string `json:"ship_to_days,omitempty" xml:"ship_to_days,omitempty"`
	// 佣金率
	CommisionRate string `json:"commision_rate,omitempty" xml:"commision_rate,omitempty"`
}
