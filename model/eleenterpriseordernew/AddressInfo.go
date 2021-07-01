package eleenterpriseordernew

// AddressInfo 结构体
type AddressInfo struct {
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收货人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
