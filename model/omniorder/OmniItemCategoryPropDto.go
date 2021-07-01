package omniorder

// OmniItemCategoryPropDto 结构体
type OmniItemCategoryPropDto struct {
	// 销售属性名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// pid
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
}
