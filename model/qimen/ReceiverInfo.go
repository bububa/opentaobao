package qimen

// ReceiverInfo 结构体
type ReceiverInfo struct {
	// 公司名称
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 邮编
	ZipCode string `json:"zipCode,omitempty" xml:"zipCode,omitempty"`
	// 固定电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 收件人证件类型(1-身份证、2-军官证、3-护照、4-其他)
	IdType string `json:"idType,omitempty" xml:"idType,omitempty"`
	// 收件人证件号码
	IdNumber string `json:"idNumber,omitempty" xml:"idNumber,omitempty"`
	// 电子邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 国家二字码
	CountryCode string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区域
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 村镇
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
	// oaid
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 1
	Privacy string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// 收件人所在区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 证件号
	Id string `json:"id,omitempty" xml:"id,omitempty"`
}
