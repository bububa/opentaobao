package omniorder

// TpDto 结构体
type TpDto struct {
	// 公司编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
