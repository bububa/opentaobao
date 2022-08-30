package axintrade

// ContactInfoApiDto 结构体
type ContactInfoApiDto struct {
	// 联系人邮箱
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 联系人手机号
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人固定电话
	ContactTel string `json:"contact_tel,omitempty" xml:"contact_tel,omitempty"`
}
