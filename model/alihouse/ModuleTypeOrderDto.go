package alihouse

// ModuleTypeOrderDto 结构体
type ModuleTypeOrderDto struct {
	// 模块类型
	ModuleType int64 `json:"module_type,omitempty" xml:"module_type,omitempty"`
	// 模块排序
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
}
