package wlb

// ReceiverInfo 
type ReceiverInfo struct {

    // 收货人移动电话
    Mobile   string `json:"mobile,omitempty"`

    // 收货人姓名
    Name   string `json:"name,omitempty"`

    // 收货人详细地址
    DetailAddress   string `json:"detail_address,omitempty"`

    // 收货人镇
    Town   string `json:"town,omitempty"`

    // 收货人区
    Area   string `json:"area,omitempty"`

    // 收货人市
    City   string `json:"city,omitempty"`

    // 收货人省
    Province   string `json:"province,omitempty"`

}
