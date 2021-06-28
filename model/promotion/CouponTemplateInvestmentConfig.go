package promotion

// CouponTemplateInvestmentConfig 
/* model for simplify = false
type CouponTemplateInvestmentConfig struct {

    // 出资人配置
    
    InvestmentInfoList  struct {
        InvestmentInfo  []InvestmentInfo `json:"investment_info,omitempty"`
    } `json:"investment_info_list,omitempty"`
    

}
*/

// CouponTemplateInvestmentConfig 
type CouponTemplateInvestmentConfig struct {

    // 出资人配置
    InvestmentInfoList   []InvestmentInfo `json:"investment_info_list,omitempty"`

}
