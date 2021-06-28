package promotion

// MobileBenefitSendResultExt 
/* model for simplify = false
type MobileBenefitSendResultExt struct {

    // 错误码列表
    
    ErrorCodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_code_list,omitempty"`
    

    // 失败份数
    
    FailureCount   int64 `json:"failure_count,omitempty"`
    

    // 成功份数
    
    SuccessCount   int64 `json:"success_count,omitempty"`
    

    // 活动详情id
    
    IndexId   int64 `json:"index_id,omitempty"`
    

}
*/

// MobileBenefitSendResultExt 
type MobileBenefitSendResultExt struct {

    // 错误码列表
    ErrorCodeList   []string `json:"error_code_list,omitempty"`

    // 失败份数
    FailureCount   int64 `json:"failure_count,omitempty"`

    // 成功份数
    SuccessCount   int64 `json:"success_count,omitempty"`

    // 活动详情id
    IndexId   int64 `json:"index_id,omitempty"`

}
