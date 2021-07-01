package aecreatives

// PromoCodeDto 结构体
type PromoCodeDto struct {
	// 专属绑定PID的code码
	PromoCode string `json:"promo_code,omitempty" xml:"promo_code,omitempty"`
	// 优惠方式 1 满减，2 满折
	CodeCampaigntype string `json:"code_campaigntype,omitempty" xml:"code_campaigntype,omitempty"`
	// 面额
	CodeValue string `json:"code_value,omitempty" xml:"code_value,omitempty"`
	// code使用有效期的开始时间
	CodeAvailabletimeStart string `json:"code_availabletime_start,omitempty" xml:"code_availabletime_start,omitempty"`
	// code使用有效期的结束时间
	CodeAvailabletimeEnd string `json:"code_availabletime_end,omitempty" xml:"code_availabletime_end,omitempty"`
	// 最低使用门槛
	CodeMiniSpend string `json:"code_mini_spend,omitempty" xml:"code_mini_spend,omitempty"`
	// code剩余可使用的数量
	CodeQuantity string `json:"code_quantity,omitempty" xml:"code_quantity,omitempty"`
	// 品code合一url
	CodePromotionurl string `json:"code_promotionurl,omitempty" xml:"code_promotionurl,omitempty"`
}
