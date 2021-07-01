package ieagency

// IeChangeContactVo 结构体
type IeChangeContactVo struct {
	// 电话国际区号
	MobileCountryCode string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// 电话号
	MobilePhoneNumber string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
