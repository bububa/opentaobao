package wlb

// PartnerNew 结构体
type PartnerNew struct {
	// 物流商编码
	TpCode string `json:"tp_code,omitempty" xml:"tp_code,omitempty"`
	// 物流商名称
	TpName string `json:"tp_name,omitempty" xml:"tp_name,omitempty"`
	// 是否虚拟物流商
	IsVirtualTp bool `json:"is_virtual_tp,omitempty" xml:"is_virtual_tp,omitempty"`
	// 服务类型
	ServiceType int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
}
