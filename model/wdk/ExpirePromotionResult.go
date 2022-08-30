package wdk

// ExpirePromotionResult 结构体
type ExpirePromotionResult struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// skuCode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// merchantCode
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// shopId
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
