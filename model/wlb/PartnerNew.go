package wlb

// PartnerNew 
type PartnerNew struct {

    // 物流商编码
    TpCode   string `json:"tp_code,omitempty"`

    // 物流商名称
    TpName   string `json:"tp_name,omitempty"`

    // 是否虚拟物流商
    IsVirtualTp   bool `json:"is_virtual_tp,omitempty"`

    // 服务类型
    ServiceType   int64 `json:"service_type,omitempty"`

}
