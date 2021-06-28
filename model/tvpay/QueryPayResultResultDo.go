package tvpay

// QueryPayResultResultDo 
/* model for simplify = false
type QueryPayResultResultDo struct {

    // 支付资金组成情况
    
    FundMoney   string `json:"fund_money,omitempty"`
    

    // 支付资金组成情况
    
    FundMoneyCode   string `json:"fund_money_code,omitempty"`
    

    // 订单状态
    
    Status   string `json:"status,omitempty"`
    

}
*/

// QueryPayResultResultDo 
type QueryPayResultResultDo struct {

    // 支付资金组成情况
    FundMoney   string `json:"fund_money,omitempty"`

    // 支付资金组成情况
    FundMoneyCode   string `json:"fund_money_code,omitempty"`

    // 订单状态
    Status   string `json:"status,omitempty"`

}
