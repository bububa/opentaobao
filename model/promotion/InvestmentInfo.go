package promotion

// InvestmentInfo 
type InvestmentInfo struct {

    // 出资人
    
    Investor   string `json:"investor,omitempty" xml:"investor,omitempty"`
    

    // 出资比例 2000 = 20%
    
    InvestorRatio   int64 `json:"investor_ratio,omitempty" xml:"investor_ratio,omitempty"`
    

}
