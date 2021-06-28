package util

// AlibabaInteractSupplierAwardResourceGetCuntaoResultModel 
/* model for simplify = false
type AlibabaInteractSupplierAwardResourceGetCuntaoResultModel struct {

    // 错误的描述信息（如果有）
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 返回值,里面只返回资源ID
    
    Model  *struct {
        AwardActivityDetailDto  *AwardActivityDetailDto `json:"award_activity_detail_dto,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误码（如果有）
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 是否成功，true/false
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaInteractSupplierAwardResourceGetCuntaoResultModel 
type AlibabaInteractSupplierAwardResourceGetCuntaoResultModel struct {

    // 错误的描述信息（如果有）
    ErrorMessage   string `json:"error_message,omitempty"`

    // 返回值,里面只返回资源ID
    Model   *AwardActivityDetailDto `json:"model,omitempty"`

    // 错误码（如果有）
    ErrorCode   string `json:"error_code,omitempty"`

    // 是否成功，true/false
    Success   bool `json:"success,omitempty"`

}
