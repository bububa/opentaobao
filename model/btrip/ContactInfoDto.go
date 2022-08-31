package btrip

// ContactInfoDto 结构体
type ContactInfoDto struct {
	// 联系优先
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 联系人名字
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系手机号
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
}
