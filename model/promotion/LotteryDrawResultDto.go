package promotion

// LotteryDrawResultDto 结构体
type LotteryDrawResultDto struct {
	// resultType
	ResultType int64 `json:"result_type,omitempty" xml:"result_type,omitempty"`
	// resultMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// shopTitle
	ShopTitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// shopLogo
	ShopLogo string `json:"shop_logo,omitempty" xml:"shop_logo,omitempty"`
	// template
	Template string `json:"template,omitempty" xml:"template,omitempty"`
	// award
	Award *LotteryAwardDto `json:"award,omitempty" xml:"award,omitempty"`
}
