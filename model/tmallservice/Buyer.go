package tmallservice

// Buyer 
type Buyer struct {
    // 省
    AddressProvince   string `json:"address_province,omitempty" xml:"address_province,omitempty"`
    // 买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    // 买家手机号
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 买家姓名
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    // 街道
    AddressTown   string `json:"address_town,omitempty" xml:"address_town,omitempty"`
    // 买家详细地址
    AddressDetail   string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
    // 买家固定电话
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 买家完整地址
    FullAddress   string `json:"full_address,omitempty" xml:"full_address,omitempty"`
    // 地区编码
    Location   int64 `json:"location,omitempty" xml:"location,omitempty"`
    // 区
    AddressDistrict   string `json:"address_district,omitempty" xml:"address_district,omitempty"`
    // 市
    AddressCity   string `json:"address_city,omitempty" xml:"address_city,omitempty"`
    // 邮箱
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // 拼装好的地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 邮编
    ZipCode   string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
}
