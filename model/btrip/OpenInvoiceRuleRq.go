package btrip

// OpenInvoiceRuleRq 
type OpenInvoiceRuleRq struct {

    // 是否适用所有员工
    
    AllEmploye   bool `json:"all_employe,omitempty" xml:"all_employe,omitempty"`
    

    // 第三方企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // all_employe为true时非必传，否则必传
    
    Entities   []Entity `json:"entities,omitempty" xml:"entities,omitempty"`
    

    // 第三方发票id
    
    ThirdPartId   string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
    

    // 商旅开放平台传2
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

}
