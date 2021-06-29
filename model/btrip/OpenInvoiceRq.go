package btrip

// OpenInvoiceRq 
type OpenInvoiceRq struct {

    // 第三方企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 发票抬头关键词
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 第三方用户id
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 2：商旅开放平台
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

}
