package ascp

// CancelDistributeDetail 结构体
type CancelDistributeDetail struct {
	// 传入的商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 传入的商品skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 分销商昵称
	DistributorCompanyName string `json:"distributor_company_name,omitempty" xml:"distributor_company_name,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 分销商 shopUserId
	DistributorShopUserId int64 `json:"distributor_shop_user_id,omitempty" xml:"distributor_shop_user_id,omitempty"`
	// 处理结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
