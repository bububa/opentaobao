package tuanhotel

// TopAppointPhoneVO 
type TopAppointPhoneVO struct {
    // 手机号
    Mobil   string `json:"mobil,omitempty" xml:"mobil,omitempty"`
    // 地区
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 国家
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    // 固话号码
    Fix   string `json:"fix,omitempty" xml:"fix,omitempty"`
    // p400
    P400   string `json:"p400,omitempty" xml:"p400,omitempty"`
    // 类型固话或手机
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
