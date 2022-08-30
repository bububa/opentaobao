package wdk

// BuyGiftActivitySkuDto 结构体
type BuyGiftActivitySkuDto struct {
	// 操作人ID
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 操作人姓名
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 赠品编码
	GiftSkuCode string `json:"gift_sku_code,omitempty" xml:"gift_sku_code,omitempty"`
	// 赠品条码
	GiftBarCode string `json:"gift_bar_code,omitempty" xml:"gift_bar_code,omitempty"`
	// 营销活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 限购信息
	Limit *LimitDto `json:"limit,omitempty" xml:"limit,omitempty"`
	// 买N赠M的N参数
	BuyNum int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
	// 买N赠M的M参数
	GiftNum int64 `json:"gift_num,omitempty" xml:"gift_num,omitempty"`
}
