package wlb

// Partner 结构体
type Partner struct {
	// 物流商名称
	TpName string `json:"tp_name,omitempty" xml:"tp_name,omitempty"`
	// 物流商编码
	TpCode string `json:"tp_code,omitempty" xml:"tp_code,omitempty"`
	// 服务类型
	ServiceType int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 是否虚拟物流商
	IsVirtualTp bool `json:"is_virtual_tp,omitempty" xml:"is_virtual_tp,omitempty"`
}
