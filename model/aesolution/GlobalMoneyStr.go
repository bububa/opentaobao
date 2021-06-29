package aesolution

// GlobalMoneyStr 
type GlobalMoneyStr struct {
    // Cent
    Cent   int64 `json:"cent,omitempty" xml:"cent,omitempty"`
    // Currency code
    CurrencyCode   string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
    // Amount
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // The factor to the smallest currency unit
    CentFactor   int64 `json:"cent_factor,omitempty" xml:"cent_factor,omitempty"`
    // Currency description
    Currency   *Currency `json:"currency,omitempty" xml:"currency,omitempty"`
    // Amount
    AmountStr   string `json:"amount_str,omitempty" xml:"amount_str,omitempty"`
}
