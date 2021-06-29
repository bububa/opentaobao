package fundplatform

// AccountChargeResponse 
type AccountChargeResponse struct {

    // 充值的账户ID
    
    AccountId   int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
    

    // 充值URL
    
    PayUrl   string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
    

}
