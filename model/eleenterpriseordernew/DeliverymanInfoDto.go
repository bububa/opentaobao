package eleenterpriseordernew

// DeliverymanInfoDto 结构体
type DeliverymanInfoDto struct {
	// 配送员姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 配送员电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}
