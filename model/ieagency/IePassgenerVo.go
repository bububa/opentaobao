package ieagency

// IePassgenerVo 结构体
type IePassgenerVo struct {
	// 美国居住地-城市
	AddressCity string `json:"address_city,omitempty" xml:"address_city,omitempty"`
	// 美国居住地-邮编
	AddressPostcode string `json:"address_postcode,omitempty" xml:"address_postcode,omitempty"`
	// 美国加拿大居住地-州
	AddressState string `json:"address_state,omitempty" xml:"address_state,omitempty"`
	// 美国居住地-街道
	AddressStreet string `json:"address_street,omitempty" xml:"address_street,omitempty"`
	// 出生日期
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 证件签发国
	CertIssueCountry string `json:"cert_issue_country,omitempty" xml:"cert_issue_country,omitempty"`
	// 乘机人证件号
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 证件有效期
	CertPeriod string `json:"cert_period,omitempty" xml:"cert_period,omitempty"`
	// 乘机人证件类型(Passport:护照,Hongkong_Macao:港澳通行证,Taiwan_MTP:台胞证,Home_Return_Permit:回乡证,Taiwan_Pass:台湾通行证,Entry_Taiwan_Permit:入台证);
	CertType string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 性别(MALE:男,FEMALE:女)
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 乘机人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 国籍
	Nationality string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// 乘机人类型(Adult:普通成人,Child:儿童,StudentAbroad:留学生,Infant:婴儿)
	PassengerType string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 常旅客积分分属航司二字码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 常旅客卡号
	AirlineCardNo string `json:"airline_card_no,omitempty" xml:"airline_card_no,omitempty"`
}
