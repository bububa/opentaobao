package drug

// FundBillDto 
/* model for simplify = false
type FundBillDto struct {

    // 支付渠道代码
    
    FundChannel   string `json:"fund_channel,omitempty"`
    

    // 金额
    
    Amount   string `json:"amount,omitempty"`
    

}
*/

// FundBillDto 
type FundBillDto struct {

    // 支付渠道代码
    FundChannel   string `json:"fund_channel,omitempty"`

    // 金额
    Amount   string `json:"amount,omitempty"`

}
