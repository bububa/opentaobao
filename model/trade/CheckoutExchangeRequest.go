package trade

// CheckoutExchangeRequest 
/* model for simplify = false
type CheckoutExchangeRequest struct {

    // 基准币种(卖家设置的）
    
    BaseCurrency   string `json:"base_currency,omitempty"`
    

    // 报价币种(买家看到的）
    
    QuoteCurrency   string `json:"quote_currency,omitempty"`
    

}
*/

// CheckoutExchangeRequest 
type CheckoutExchangeRequest struct {

    // 基准币种(卖家设置的）
    BaseCurrency   string `json:"base_currency,omitempty"`

    // 报价币种(买家看到的）
    QuoteCurrency   string `json:"quote_currency,omitempty"`

}
