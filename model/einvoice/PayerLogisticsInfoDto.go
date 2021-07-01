package einvoice

// PayerLogisticsInfoDto 
type PayerLogisticsInfoDto struct {
    // 收件人地址
    ContactAddr   string `json:"contact_addr,omitempty" xml:"contact_addr,omitempty"`
    // 收件人电话
    ContactMobile   string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
    // 收件人姓名
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
}
