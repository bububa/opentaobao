package tmallservice

// Buyer 结构体
type Buyer struct {
	// 座机
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 二级地址
	AddressCity string `json:"address_city,omitempty" xml:"address_city,omitempty"`
	// 三级地址
	AddressDistrict string `json:"address_district,omitempty" xml:"address_district,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 拼装好的地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 四级地址
	AddressTown string `json:"address_town,omitempty" xml:"address_town,omitempty"`
	// 一级地址
	AddressProvince string `json:"address_province,omitempty" xml:"address_province,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// open_uid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 地址编码
	Location int64 `json:"location,omitempty" xml:"location,omitempty"`
}
