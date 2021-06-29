package lsttrade

// ReceiverInfo 
type ReceiverInfo struct {
    // 收货人地址区域编码
    ToDivisionCode   string `json:"to_division_code,omitempty" xml:"to_division_code,omitempty"`
    // 收货人街道或镇区域编码，可能为空
    ToTownCode   string `json:"to_town_code,omitempty" xml:"to_town_code,omitempty"`
    // 详情收货地址（到门牌号）
    ToDetailArea   string `json:"to_detail_area,omitempty" xml:"to_detail_area,omitempty"`
    // 收件人
    ToFullName   string `json:"to_full_name,omitempty" xml:"to_full_name,omitempty"`
    // 移动电话
    ToMobile   string `json:"to_mobile,omitempty" xml:"to_mobile,omitempty"`
    // 固定电话
    ToPhone   string `json:"to_phone,omitempty" xml:"to_phone,omitempty"`
    // 邮编
    ToPost   string `json:"to_post,omitempty" xml:"to_post,omitempty"`
    // 收货地址（只到街道）
    ToArea   string `json:"to_area,omitempty" xml:"to_area,omitempty"`
    // 详细地址
    ToDetailAddress   string `json:"to_detail_address,omitempty" xml:"to_detail_address,omitempty"`
}
