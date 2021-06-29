package tmallcar

// CarLeasePostSchemeDto 
type CarLeasePostSchemeDto struct {

    // 0:不能使用,1:可以使用
    
    CanSelect   int64 `json:"can_select,omitempty" xml:"can_select,omitempty"`
    

    // 分期月数
    
    Month   int64 `json:"month,omitempty" xml:"month,omitempty"`
    

    // 月供，单位：分
    
    MonthlyPay   int64 `json:"monthly_pay,omitempty" xml:"monthly_pay,omitempty"`
    

    // 不能使用原因code
    
    ReasonCode   string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
    

    // 不能使用原因描述
    
    ReasonDesc   string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
    

    // 续租到期后可选择方案列表
    
    RenewSchemeList   []CarLeasePostSchemeDto `json:"renew_scheme_list,omitempty" xml:"renew_scheme_list,omitempty"`
    

    // 尾款，单位：分
    
    RestAmount   int64 `json:"rest_amount,omitempty" xml:"rest_amount,omitempty"`
    

    // 0.退车,1.买断,2.分期，3.续租
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

}
