package xhotelonlineorder

// ValueAddedInfo 
type ValueAddedInfo struct {
    // 电话
    Tel   string `json:"tel,omitempty" xml:"tel,omitempty"`
    // 注册地址
    RegisterAddr   string `json:"register_addr,omitempty" xml:"register_addr,omitempty"`
    // 银行账号
    AccountNo   string `json:"account_no,omitempty" xml:"account_no,omitempty"`
    // 开户行
    OpeningBank   string `json:"opening_bank,omitempty" xml:"opening_bank,omitempty"`
    // 税号
    TaxNo   string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
}
