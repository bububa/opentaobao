package wdk

// WdkOpenSkuMerchantCatServiceQueryResult 结构体
type WdkOpenSkuMerchantCatServiceQueryResult struct {
	// 结果描述
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 结果码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// {"skuCode":"categoryCode"}，skuCode和商家叶子类目编码映射关系
	SkuCodeCategoryCodeMap string `json:"sku_code_category_code_map,omitempty" xml:"sku_code_category_code_map,omitempty"`
	// 成功或失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
