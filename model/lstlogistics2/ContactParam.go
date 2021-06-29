package lstlogistics2

// ContactParam 
type ContactParam struct {
    // 联系人姓名
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
    // 手机号为11位
    ContactMobile   string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
    // 电话
    ContactPhone   string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
}
