package promotion

// BenefitSendResultExt 
/* model for simplify = false
type BenefitSendResultExt struct {

    // 活动详情id
    
    IndexId   int64 `json:"index_id,omitempty"`
    

    // 异常码列表
    
    ErrorCodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_code_list,omitempty"`
    

    // 失败数
    
    FailureCount   int64 `json:"failure_count,omitempty"`
    

    // 成功数
    
    SuccessCount   int64 `json:"success_count,omitempty"`
    

}
*/

// BenefitSendResultExt 
type BenefitSendResultExt struct {

    // 活动详情id
    IndexId   int64 `json:"index_id,omitempty"`

    // 异常码列表
    ErrorCodeList   []string `json:"error_code_list,omitempty"`

    // 失败数
    FailureCount   int64 `json:"failure_count,omitempty"`

    // 成功数
    SuccessCount   int64 `json:"success_count,omitempty"`

}
