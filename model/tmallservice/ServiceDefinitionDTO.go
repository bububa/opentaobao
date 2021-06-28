package tmallservice

// ServiceDefinitionDTO 
/* model for simplify = false
type ServiceDefinitionDTO struct {

    // 服务编码
    
    ServiceCode   string `json:"service_code,omitempty"`
    

    // 服务名称
    
    DisplayName   string `json:"display_name,omitempty"`
    

    // 业务编码
    
    BizCode   string `json:"biz_code,omitempty"`
    

}
*/

// ServiceDefinitionDTO 
type ServiceDefinitionDTO struct {

    // 服务编码
    ServiceCode   string `json:"service_code,omitempty"`

    // 服务名称
    DisplayName   string `json:"display_name,omitempty"`

    // 业务编码
    BizCode   string `json:"biz_code,omitempty"`

}
