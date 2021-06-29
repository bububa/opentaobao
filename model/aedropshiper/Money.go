package aedropshiper

// Money 
type Money struct {

    // amount
    
    Amount   *BigDecimal `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // cent
    
    Cent   int64 `json:"cent,omitempty" xml:"cent,omitempty"`
    

    // currencyCode
    
    CurrencyCode   string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
    

}
