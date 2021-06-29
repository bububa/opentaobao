package fivee

// ImportProduct 
type ImportProduct struct {

    // 条形码
    
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    

    // 证照信息
    
    Licences   []Licence `json:"licences,omitempty" xml:"licences,omitempty"`
    

    // 名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 代理商信息
    
    AgentCompanies   []Company `json:"agent_companies,omitempty" xml:"agent_companies,omitempty"`
    

    // 批次或备案编号
    
    AuthCode   string `json:"auth_code,omitempty" xml:"auth_code,omitempty"`
    

    // 检验检疫证明
    
    SanitationCertificates   []SanitationCertificate `json:"sanitation_certificates,omitempty" xml:"sanitation_certificates,omitempty"`
    

    // 业务方内部编号（比如大润发的货号）
    
    InnerCode   string `json:"inner_code,omitempty" xml:"inner_code,omitempty"`
    

    // 业务备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

}
