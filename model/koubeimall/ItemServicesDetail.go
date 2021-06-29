package koubeimall

// ItemServicesDetail 
type ItemServicesDetail struct {
    // 服务明细列表
    ServiceDetails   []string `json:"service_details,omitempty" xml:"service_details>string,omitempty"`
    // 服务名称
    ServiceName   string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}
