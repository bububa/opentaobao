package baichuan

// AlibabaBaichuanCtgUserRelationResult 
/* model for simplify = false
type AlibabaBaichuanCtgUserRelationResult struct {

    // 返回结果，数值型，1：代表已绑定达人身份，2代表未绑定达人身份
    
    Module   int64 `json:"module,omitempty"`
    

    // 错误信息
    
    ErrorDetail   string `json:"error_detail,omitempty"`
    

    // 错误码结构体
    
    ErrorCode  *struct {
        ErrorCode  *ErrorCode `json:"error_code,omitempty"`
    } `json:"error_code,omitempty"`
    

    // 返回成功与否
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaBaichuanCtgUserRelationResult 
type AlibabaBaichuanCtgUserRelationResult struct {

    // 返回结果，数值型，1：代表已绑定达人身份，2代表未绑定达人身份
    Module   int64 `json:"module,omitempty"`

    // 错误信息
    ErrorDetail   string `json:"error_detail,omitempty"`

    // 错误码结构体
    ErrorCode   *ErrorCode `json:"error_code,omitempty"`

    // 返回成功与否
    Success   bool `json:"success,omitempty"`

}
