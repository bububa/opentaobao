package ieagency

// ContactsParam 
type ContactsParam struct {
    // 联系人邮箱地址
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // 联系人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 联系人手机号
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 联系人手机号国家编码
    PhoneCountryCode   string `json:"phone_country_code,omitempty" xml:"phone_country_code,omitempty"`
}
