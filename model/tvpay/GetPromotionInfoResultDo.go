package tvpay

// GetPromotionInfoResultDo 结构体
type GetPromotionInfoResultDo struct {
	// 描述
	Hint string `json:"hint,omitempty" xml:"hint,omitempty"`
	// 是否有抽奖活动
	HasPromotionEvent bool `json:"has_promotion_event,omitempty" xml:"has_promotion_event,omitempty"`
}
