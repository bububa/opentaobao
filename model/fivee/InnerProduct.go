package fivee

// InnerProduct 
type InnerProduct struct {

    // 批次或备案证书编号
    
    AuthCode   string `json:"auth_code,omitempty" xml:"auth_code,omitempty"`
    

    // 条形码
    
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    

    // 证照信息
    
    Licences   []Licence `json:"licences,omitempty" xml:"licences,omitempty"`
    

    // 商品名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 生产商
    
    ProduceCompanies   []Company `json:"produce_companies,omitempty" xml:"produce_companies,omitempty"`
    

    // 供应商信息
    
    ProviderCompanies   []Company `json:"provider_companies,omitempty" xml:"provider_companies,omitempty"`
    

    // 各业务方内部编号（例如RT的内部货号）
    
    InnerCode   string `json:"inner_code,omitempty" xml:"inner_code,omitempty"`
    

    // 业务方备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

}
