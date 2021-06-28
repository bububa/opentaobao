package wlb

// JzPartnerNew 
type JzPartnerNew struct {

    // 服务商编码
    
    TpCode   string `json:"tp_code,omitempty" xml:"tp_code,omitempty"`
    

    // 服务商名称
    
    TpName   string `json:"tp_name,omitempty" xml:"tp_name,omitempty"`
    

    // 服务类型
    
    ServiceType   int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
    

    // 是否是虚拟服务商（家装-商家自有物流）
    
    IsVirtualTp   bool `json:"is_virtual_tp,omitempty" xml:"is_virtual_tp,omitempty"`
    

}
