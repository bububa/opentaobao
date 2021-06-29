package tmallcar

// LoanReceiveDto 
type LoanReceiveDto struct {

    // 放贷类型
    
    LoanType   string `json:"loan_type,omitempty" xml:"loan_type,omitempty"`
    

    // 金融服务商名称
    
    CapitalName   string `json:"capital_name,omitempty" xml:"capital_name,omitempty"`
    

    // 金融服务商编码
    
    CapitalCode   string `json:"capital_code,omitempty" xml:"capital_code,omitempty"`
    

    // 还款方式
    
    PaymentMethod   string `json:"payment_method,omitempty" xml:"payment_method,omitempty"`
    

    // 放贷利率
    
    LoanRate   string `json:"loan_rate,omitempty" xml:"loan_rate,omitempty"`
    

    // 放贷期数
    
    LoanTerm   int64 `json:"loan_term,omitempty" xml:"loan_term,omitempty"`
    

    // 放贷期数单位:DAY-按天计算;MONTH-按月计算;YEAR-按年计算
    
    LoanTermUnit   string `json:"loan_term_unit,omitempty" xml:"loan_term_unit,omitempty"`
    

    // 最终放贷额度(单位分)
    
    LoanAmount   int64 `json:"loan_amount,omitempty" xml:"loan_amount,omitempty"`
    

    // 最终放款到门店的时间
    
    LoanTime   string `json:"loan_time,omitempty" xml:"loan_time,omitempty"`
    

    // 金融机构放贷内部单号
    
    LoanId   string `json:"loan_id,omitempty" xml:"loan_id,omitempty"`
    

    // 放贷结果
    
    LoanResult   string `json:"loan_result,omitempty" xml:"loan_result,omitempty"`
    

    // 金融机构的授信申请内部单号
    
    CreditId   string `json:"credit_id,omitempty" xml:"credit_id,omitempty"`
    

    // 天猫汽车侧业务单号
    
    TmallBizNo   string `json:"tmall_biz_no,omitempty" xml:"tmall_biz_no,omitempty"`
    

    // 业务类型
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    

    // 请求时间
    
    RequestTime   string `json:"request_time,omitempty" xml:"request_time,omitempty"`
    

    // 扩展字段1
    
    ExtensionField02   string `json:"extension_field02,omitempty" xml:"extension_field02,omitempty"`
    

    // 扩展字段2
    
    ExtensionField01   string `json:"extension_field01,omitempty" xml:"extension_field01,omitempty"`
    

}
