package alitripmerchant

// EnumVo 结构体
type EnumVo struct {
	// 枚举值名称
	EnumName string `json:"enum_name,omitempty" xml:"enum_name,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 枚举代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 其他备注
	Other string `json:"other,omitempty" xml:"other,omitempty"`
	// 排序
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
}
