package wlb

// ReceiverInfo 
type ReceiverInfo struct {
    // 收货人移动电话
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 收货人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 收货人详细地址
    DetailAddress   string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
    // 收货人镇
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    // 收货人区
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 收货人市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 收货人省
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
}
