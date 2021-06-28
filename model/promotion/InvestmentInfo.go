package promotion

// InvestmentInfo 
/* model for simplify = false
type InvestmentInfo struct {

    // 出资人
    
    Investor   string `json:"investor,omitempty"`
    

    // 出资比例 2000 = 20%
    
    InvestorRatio   int64 `json:"investor_ratio,omitempty"`
    

}
*/

// InvestmentInfo 
type InvestmentInfo struct {

    // 出资人
    Investor   string `json:"investor,omitempty"`

    // 出资比例 2000 = 20%
    InvestorRatio   int64 `json:"investor_ratio,omitempty"`

}
