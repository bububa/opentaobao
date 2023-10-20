package wdk

// AlibabahmmarketingexpirepromotioncreateT 结构体
type AlibabahmmarketingexpirepromotioncreateT struct {
	// 商家code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 商品itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
