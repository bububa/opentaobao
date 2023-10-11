package alicom

// Interface 结构体
type Interface struct {
	// 接口类型
	ApiType string `json:"api_type,omitempty" xml:"api_type,omitempty"`
	// 接口地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
