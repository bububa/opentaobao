package happytrip

// Managers 结构体
type Managers struct {
	// 联系方式
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 联系人名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
