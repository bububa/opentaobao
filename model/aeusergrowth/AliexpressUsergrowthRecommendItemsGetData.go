package aeusergrowth

import (
	"sync"
)

// AliexpressUsergrowthRecommendItemsGetData 结构体
type AliexpressUsergrowthRecommendItemsGetData struct {
	// product id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// sale price,will return local pirce
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// product photo url
	Photo string `json:"photo,omitempty" xml:"photo,omitempty"`
	// shop name
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// product name
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// store url
	ShopDomain string `json:"shop_domain,omitempty" xml:"shop_domain,omitempty"`
	// detailUrl
	DetailUrl string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
	// shopId
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// sub category
	SubCategory string `json:"sub_category,omitempty" xml:"sub_category,omitempty"`
	// shop rating
	ShopRating string `json:"shop_rating,omitempty" xml:"shop_rating,omitempty"`
	// product rating
	RatingValue string `json:"rating_value,omitempty" xml:"rating_value,omitempty"`
	// category Name
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// rating count
	RatingCount int64 `json:"rating_count,omitempty" xml:"rating_count,omitempty"`
	// photo height
	PhotoHeight int64 `json:"photo_height,omitempty" xml:"photo_height,omitempty"`
	// photo width
	PhotoWidth int64 `json:"photo_width,omitempty" xml:"photo_width,omitempty"`
	// delivery freeShipping
	FreeShipping bool `json:"free_shipping,omitempty" xml:"free_shipping,omitempty"`
}

var poolAliexpressUsergrowthRecommendItemsGetData = sync.Pool{
	New: func() any {
		return new(AliexpressUsergrowthRecommendItemsGetData)
	},
}

// GetAliexpressUsergrowthRecommendItemsGetData() 从对象池中获取AliexpressUsergrowthRecommendItemsGetData
func GetAliexpressUsergrowthRecommendItemsGetData() *AliexpressUsergrowthRecommendItemsGetData {
	return poolAliexpressUsergrowthRecommendItemsGetData.Get().(*AliexpressUsergrowthRecommendItemsGetData)
}

// ReleaseAliexpressUsergrowthRecommendItemsGetData 释放AliexpressUsergrowthRecommendItemsGetData
func ReleaseAliexpressUsergrowthRecommendItemsGetData(v *AliexpressUsergrowthRecommendItemsGetData) {
	v.ProductId = ""
	v.SalePrice = ""
	v.Photo = ""
	v.ShopName = ""
	v.ProductName = ""
	v.ShopDomain = ""
	v.DetailUrl = ""
	v.ShopId = ""
	v.SubCategory = ""
	v.ShopRating = ""
	v.RatingValue = ""
	v.Category = ""
	v.RatingCount = 0
	v.PhotoHeight = 0
	v.PhotoWidth = 0
	v.FreeShipping = false
	poolAliexpressUsergrowthRecommendItemsGetData.Put(v)
}
