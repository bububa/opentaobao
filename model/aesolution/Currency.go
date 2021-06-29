package aesolution

// Currency 
type Currency struct {

    // The default minimum accuracy of the currency
    
    DefaultFractionDigits   int64 `json:"default_fraction_digits,omitempty" xml:"default_fraction_digits,omitempty"`
    

    // Numeric code of the currency
    
    NumericCode   int64 `json:"numeric_code,omitempty" xml:"numeric_code,omitempty"`
    

    // Currency code
    
    CurrencyCode   string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
    

    // Symbol of the currency
    
    Symbol   string `json:"symbol,omitempty" xml:"symbol,omitempty"`
    

    // Display name of the currency
    
    DisplayName   string `json:"display_name,omitempty" xml:"display_name,omitempty"`
    

}
