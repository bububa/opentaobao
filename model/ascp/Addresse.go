package ascp

// Addresse 结构体
type Addresse struct {
	// 收货地：省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 收货地：城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货地：区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 收货地：街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}
