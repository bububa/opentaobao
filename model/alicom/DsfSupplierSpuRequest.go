package alicom

// DsfSupplierSpuRequest 结构体
type DsfSupplierSpuRequest struct {
	// 分组
	GroupList []GroupRequest `json:"group_list,omitempty" xml:"group_list>group_request,omitempty"`
	// 业务类型
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
}
