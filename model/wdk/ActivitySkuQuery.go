package wdk

// ActivitySkuQuery 结构体
type ActivitySkuQuery struct {
	// 五道口活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 分页参数
	Page *BasePageQuery `json:"page,omitempty" xml:"page,omitempty"`
	// 商家活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 需要查询的商品skuCodes
	SkuCodes []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
}
