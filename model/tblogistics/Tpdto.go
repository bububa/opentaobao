package tblogistics

// Tpdto 结构体
type Tpdto struct {
	// 公司编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否商家绑定的默认安装公司
	IsDefault bool `json:"is_default,omitempty" xml:"is_default,omitempty"`
}
