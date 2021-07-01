package icbulogistics

// ContactAddress 结构体
type ContactAddress struct {
	// 国家、省、市、详细地址信息
	Address *Address `json:"address,omitempty" xml:"address,omitempty"`
	// 联系人
	ContactPerson string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
	// 联系方式(邮箱、电话号码、手机号码等)
	Contact *Contact `json:"contact,omitempty" xml:"contact,omitempty"`
	// 公司名称
	CompanyNameCn string `json:"company_name_cn,omitempty" xml:"company_name_cn,omitempty"`
}
