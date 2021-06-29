package tmallservice

// CustomerInfo 
type CustomerInfo struct {
    // 寄件人名称
    AccountNick   string `json:"account_nick,omitempty" xml:"account_nick,omitempty"`
    // 寄件人手机号码
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 寄件地址
    FullAddress   string `json:"full_address,omitempty" xml:"full_address,omitempty"`
}
