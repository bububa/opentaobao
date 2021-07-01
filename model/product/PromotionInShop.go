package product

// PromotionInShop 结构体
type PromotionInShop struct {
	// 优惠活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// idValue值
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 优惠详情描述。
	PromotionDetailDesc string `json:"promotion_detail_desc,omitempty" xml:"promotion_detail_desc,omitempty"`
}
