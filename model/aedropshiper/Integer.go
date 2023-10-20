package aedropshiper

import (
	"sync"
)

// Integer 结构体
type Integer struct {
	// product small image urls
	ProductSmallImageUrls []string `json:"product_small_image_urls,omitempty" xml:"product_small_image_urls>string,omitempty"`
	// target sale price
	TargetSalePrice string `json:"target_sale_price,omitempty" xml:"target_sale_price,omitempty"`
	// evaluate rate
	EvaluateRate string `json:"evaluate_rate,omitempty" xml:"evaluate_rate,omitempty"`
	// target original price
	TargetOriginalPrice string `json:"target_original_price,omitempty" xml:"target_original_price,omitempty"`
	// second level category name
	SecondLevelCategoryName string `json:"second_level_category_name,omitempty" xml:"second_level_category_name,omitempty"`
	// product video url
	ProductVideoUrl string `json:"product_video_url,omitempty" xml:"product_video_url,omitempty"`
	// sale price
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// sale price in target currency
	TargetSalePriceCurrency string `json:"target_sale_price_currency,omitempty" xml:"target_sale_price_currency,omitempty"`
	// shop url
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// product detail url
	ProductDetailUrl string `json:"product_detail_url,omitempty" xml:"product_detail_url,omitempty"`
	// product title
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// first level category name
	FirstLevelCategoryName string `json:"first_level_category_name,omitempty" xml:"first_level_category_name,omitempty"`
	// product main image url
	ProductMainImageUrl string `json:"product_main_image_url,omitempty" xml:"product_main_image_url,omitempty"`
	// platform product type:ALL,PLAZA,TMALL
	PlatformProductType string `json:"platform_product_type,omitempty" xml:"platform_product_type,omitempty"`
	// target original price currency
	TargetOriginalPriceCurrency string `json:"target_original_price_currency,omitempty" xml:"target_original_price_currency,omitempty"`
	// ship to days
	ShipToDays string `json:"ship_to_days,omitempty" xml:"ship_to_days,omitempty"`
	// sale price currency
	SalePriceCurrency string `json:"sale_price_currency,omitempty" xml:"sale_price_currency,omitempty"`
	// original price
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// original price currency
	OriginalPriceCurrency string `json:"original_price_currency,omitempty" xml:"original_price_currency,omitempty"`
	// discount
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// lastest volume
	LastestVolume int64 `json:"lastest_volume,omitempty" xml:"lastest_volume,omitempty"`
	// seller id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// shop id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// first level category id
	FirstLevelCategoryId int64 `json:"first_level_category_id,omitempty" xml:"first_level_category_id,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// second level category id
	SecondLevelCategoryId int64 `json:"second_level_category_id,omitempty" xml:"second_level_category_id,omitempty"`
}

var poolInteger = sync.Pool{
	New: func() any {
		return new(Integer)
	},
}

// GetInteger() 从对象池中获取Integer
func GetInteger() *Integer {
	return poolInteger.Get().(*Integer)
}

// ReleaseInteger 释放Integer
func ReleaseInteger(v *Integer) {
	v.ProductSmallImageUrls = v.ProductSmallImageUrls[:0]
	v.TargetSalePrice = ""
	v.EvaluateRate = ""
	v.TargetOriginalPrice = ""
	v.SecondLevelCategoryName = ""
	v.ProductVideoUrl = ""
	v.SalePrice = ""
	v.TargetSalePriceCurrency = ""
	v.ShopUrl = ""
	v.ProductDetailUrl = ""
	v.ProductTitle = ""
	v.FirstLevelCategoryName = ""
	v.ProductMainImageUrl = ""
	v.PlatformProductType = ""
	v.TargetOriginalPriceCurrency = ""
	v.ShipToDays = ""
	v.SalePriceCurrency = ""
	v.OriginalPrice = ""
	v.OriginalPriceCurrency = ""
	v.Discount = ""
	v.LastestVolume = 0
	v.SellerId = 0
	v.ShopId = 0
	v.FirstLevelCategoryId = 0
	v.ProductId = 0
	v.SecondLevelCategoryId = 0
	poolInteger.Put(v)
}
