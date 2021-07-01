package idleisv

// AppraiseIsvAddressDto 
type AppraiseIsvAddressDto struct {
    // 详细地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 行政区
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 收件人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 省份
    Prov   string `json:"prov,omitempty" xml:"prov,omitempty"`
    // 城镇/街道
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    // 电话号码
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
}
