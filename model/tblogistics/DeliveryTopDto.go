package tblogistics

// DeliveryTopDto 结构体
type DeliveryTopDto struct {
	// 配送员电话，支持手机、座机、400电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 配送员姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
