package alicom

// DsfSupplierSpuVo 结构体
type DsfSupplierSpuVo struct {
	// 分组
	GroupList []Group `json:"group_list,omitempty" xml:"group_list>group,omitempty"`
	// 业务类型
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
}
