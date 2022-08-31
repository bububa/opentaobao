package bus

// InsurancePropertyVo 结构体
type InsurancePropertyVo struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 数据
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
