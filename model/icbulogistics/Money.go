package icbulogistics

// Money 
type Money struct {
    // 金额
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 币种
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
}
