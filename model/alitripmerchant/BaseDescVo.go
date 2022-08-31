package alitripmerchant

// BaseDescVo 结构体
type BaseDescVo struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}
