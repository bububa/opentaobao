package flight

// Taxes 
type Taxes struct {

    // 税值
    Amount   int64 `json:"amount,omitempty"`

    // 税项二字码
    TaxCode   string `json:"tax_code,omitempty"`

}
