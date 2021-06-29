package btrip

// OpenProjectRq 
type OpenProjectRq struct {

    // 项目代码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 第三方企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 项目名称
    
    ProjectName   string `json:"project_name,omitempty" xml:"project_name,omitempty"`
    

    // 第三方成本中心id
    
    ThirdPartCostCenterId   string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
    

    // 第三方项目id
    
    ThirdPartId   string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
    

    // 第三方发票id
    
    ThirdPartInvoiceId   string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
    

    // 商旅开放平台传2
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

}
