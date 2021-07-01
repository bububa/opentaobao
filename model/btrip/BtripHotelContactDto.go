package btrip

// BtripHotelContactDto 
type BtripHotelContactDto struct {
    // 邮箱
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // 入住人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 入住人电话
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
}
