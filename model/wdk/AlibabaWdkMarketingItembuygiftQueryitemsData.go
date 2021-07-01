package wdk

// AlibabaWdkMarketingItembuygiftQueryitemsData 结构体
type AlibabaWdkMarketingItembuygiftQueryitemsData struct {
	// 通用限购信息，-1为不限制，默认为不限制
	LimitInfo *LimitInfo `json:"limit_info,omitempty" xml:"limit_info,omitempty"`
	// 赠品的名称
	GiftSkuName string `json:"gift_sku_name,omitempty" xml:"gift_sku_name,omitempty"`
	// 赠品的skuCode，如果和主商品的skuCode相同，则为买A赠A；如果不同，则为买A赠B
	GiftSkuCode string `json:"gift_sku_code,omitempty" xml:"gift_sku_code,omitempty"`
	// 买N赠1的N
	BuyNum int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
	// 淘宝item和shop的对应关系， k-itemId, v-shopId
	ItemShopRelation string `json:"item_shop_relation,omitempty" xml:"item_shop_relation,omitempty"`
	// 主商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 主商品的skuCode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}
