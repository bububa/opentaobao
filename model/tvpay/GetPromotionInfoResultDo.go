package tvpay

// GetPromotionInfoResultDo 
type GetPromotionInfoResultDo struct {
    // 是否有抽奖活动
    HasPromotionEvent   bool `json:"has_promotion_event,omitempty" xml:"has_promotion_event,omitempty"`
    // 描述
    Hint   string `json:"hint,omitempty" xml:"hint,omitempty"`
}
