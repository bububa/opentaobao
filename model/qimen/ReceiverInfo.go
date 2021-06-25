package qimen

// ReceiverInfo 
type ReceiverInfo struct {

    // 公司名称
    Company   string `json:"company,omitempty"`

    // 姓名
    Name   string `json:"name,omitempty"`

    // 邮编
    ZipCode   string `json:"zipCode,omitempty"`

    // 固定电话
    Tel   string `json:"tel,omitempty"`

    // 移动电话
    Mobile   string `json:"mobile,omitempty"`

    // 收件人证件类型(1-身份证、2-军官证、3-护照、4-其他)
    IdType   string `json:"idType,omitempty"`

    // 收件人证件号码
    IdNumber   string `json:"idNumber,omitempty"`

    // 电子邮箱
    Email   string `json:"email,omitempty"`

    // 国家二字码
    CountryCode   string `json:"countryCode,omitempty"`

    // 省份
    Province   string `json:"province,omitempty"`

    // 城市
    City   string `json:"city,omitempty"`

    // 区域
    Area   string `json:"area,omitempty"`

    // 村镇
    Town   string `json:"town,omitempty"`

    // 详细地址
    DetailAddress   string `json:"detailAddress,omitempty"`

    // oaid
    Oaid   string `json:"oaid,omitempty"`

    // 证件号
    Id   string `json:"id,omitempty"`

}
