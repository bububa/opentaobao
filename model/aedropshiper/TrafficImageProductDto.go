package aedropshiper

// TrafficImageProductDto 结构体
type TrafficImageProductDto struct {
	// commodity thumbnail address list
	ProductSmallImageUrls []string `json:"product_small_image_urls,omitempty" xml:"product_small_image_urls>string,omitempty"`
	// original price
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// original price currency
	OriginalPriceCurrency string `json:"original_price_currency,omitempty" xml:"original_price_currency,omitempty"`
	// discount
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// latest volume
	LastestVolume string `json:"lastest_volume,omitempty" xml:"lastest_volume,omitempty"`
	// target sale price
	TargetSalePrice string `json:"target_sale_price,omitempty" xml:"target_sale_price,omitempty"`
	// evaluate rate
	EvaluateRate string `json:"evaluate_rate,omitempty" xml:"evaluate_rate,omitempty"`
	// target original price
	TargetOriginalPrice string `json:"target_original_price,omitempty" xml:"target_original_price,omitempty"`
	// second level category name
	SecondLevelCategoryName string `json:"second_level_category_name,omitempty" xml:"second_level_category_name,omitempty"`
	// first level category id
	FirstLevelCategoryId string `json:"first_level_category_id,omitempty" xml:"first_level_category_id,omitempty"`
	// product vedio url
	ProductVideoUrl string `json:"product_video_url,omitempty" xml:"product_video_url,omitempty"`
	// product id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// sale price
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// target sale price currency
	TargetSalePriceCurrency string `json:"target_sale_price_currency,omitempty" xml:"target_sale_price_currency,omitempty"`
	// second level category id
	SecondLevelCategoryId string `json:"second_level_category_id,omitempty" xml:"second_level_category_id,omitempty"`
	// shop url
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// product title
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// product detail url
	ProductDetailUrl string `json:"product_detail_url,omitempty" xml:"product_detail_url,omitempty"`
	// first level category name
	FirstLevelCategoryName string `json:"first_level_category_name,omitempty" xml:"first_level_category_name,omitempty"`
	// product main image url
	ProductMainImageUrl string `json:"product_main_image_url,omitempty" xml:"product_main_image_url,omitempty"`
	// platform product type
	PlatformProductType string `json:"platform_product_type,omitempty" xml:"platform_product_type,omitempty"`
	// target original price currency
	TargetOriginalPriceCurrency string `json:"target_original_price_currency,omitempty" xml:"target_original_price_currency,omitempty"`
	// sale price currency
	SalePriceCurrency string `json:"sale_price_currency,omitempty" xml:"sale_price_currency,omitempty"`
	// seller id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// shop id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}
