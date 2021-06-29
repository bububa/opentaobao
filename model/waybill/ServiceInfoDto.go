package waybill

// ServiceInfoDTO 
type ServiceInfoDTO struct {
    // 服务名称
    ServiceName   string `json:"service_name,omitempty" xml:"service_name,omitempty"`
    // 服务编码
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    // 服务属性定义
    ServiceAttributes   []ServiceAttributeDTO `json:"service_attributes,omitempty" xml:"service_attributes>service_attribute_dto,omitempty"`
    // 服务的官方描述，可以用作前端展示
    ServiceDesc   string `json:"service_desc,omitempty" xml:"service_desc,omitempty"`
    // 该服务是否为必选服务
    Required   bool `json:"required,omitempty" xml:"required,omitempty"`
}
