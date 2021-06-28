package logistic

// PartnerDetail 
type PartnerDetail struct {

    // 物流公司支付宝账号
    
    AccountNo   string `json:"account_no,omitempty" xml:"account_no,omitempty"`
    

    // 物流公司代码
    
    CompanyCode   string `json:"company_code,omitempty" xml:"company_code,omitempty"`
    

    // 物流公司id
    
    CompanyId   int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
    

    // 物流公司全名
    
    FullName   string `json:"full_name,omitempty" xml:"full_name,omitempty"`
    

    // 物流公司简称
    
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    

    // 旺旺id
    
    WangwangId   string `json:"wangwang_id,omitempty" xml:"wangwang_id,omitempty"`
    

    // 运单号验证正则表达式
    
    RegMailNo   string `json:"reg_mail_no,omitempty" xml:"reg_mail_no,omitempty"`
    

}
