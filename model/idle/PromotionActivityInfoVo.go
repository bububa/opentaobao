package idle

// PromotionActivityInfoVo 结构体
type PromotionActivityInfoVo struct {
	// 当前用户是否参与了营销活动
	Joined bool `json:"joined,omitempty" xml:"joined,omitempty"`
}
