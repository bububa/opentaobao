package aeusergrowth

// AliexpressUsergrowthSearchItemsGetData 结构体
type AliexpressUsergrowthSearchItemsGetData struct {
	// product id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// sale price , return  local price
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// product photo url
	Photo string `json:"photo,omitempty" xml:"photo,omitempty"`
	// shop name
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// product name
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// store url
	ShopDomain string `json:"shop_domain,omitempty" xml:"shop_domain,omitempty"`
	// product detail Url
	DetailUrl string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
	// shopId
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// sub category
	SubCategory string `json:"sub_category,omitempty" xml:"sub_category,omitempty"`
	// rating count
	RatingCount int64 `json:"rating_count,omitempty" xml:"rating_count,omitempty"`
	// photo height
	PhotoHeight int64 `json:"photo_height,omitempty" xml:"photo_height,omitempty"`
	// shop rating
	ShopRating string `json:"shop_rating,omitempty" xml:"shop_rating,omitempty"`
	// rating
	RatingValue string `json:"rating_value,omitempty" xml:"rating_value,omitempty"`
	// product category
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// photo width
	PhotoWidth int64 `json:"photo_width,omitempty" xml:"photo_width,omitempty"`
	// delivery freeShipping
	FreeShipping bool `json:"free_shipping,omitempty" xml:"free_shipping,omitempty"`
}
