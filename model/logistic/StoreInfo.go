package logistic

// StoreInfo 结构体
type StoreInfo struct {
	// 仓库真实名字
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// XXX果园
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 仓库服务代码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
}
