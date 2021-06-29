package hotelhstdf

// HotelMatchParam 
type HotelMatchParam struct {
    // 卖家酒店id
    Hid   int64 `json:"hid,omitempty" xml:"hid,omitempty"`
    // 标准酒店id
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
