package promotion

// CouponTemplateInvestmentConfig 结构体
type CouponTemplateInvestmentConfig struct {
	// 出资人配置
	InvestmentInfoList []InvestmentInfo `json:"investment_info_list,omitempty" xml:"investment_info_list>investment_info,omitempty"`
}
