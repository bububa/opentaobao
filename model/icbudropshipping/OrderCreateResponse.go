package icbudropshipping

// OrderCreateResponse 
type OrderCreateResponse struct {
    // pay url
    PayUrl   string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
    // order number
    TradeId   string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
}
