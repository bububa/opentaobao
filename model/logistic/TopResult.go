package logistic

// TopResult 
/* model for simplify = false
type TopResult struct {

    // 发货完成后的物流单号
    
    LgOrderCode   string `json:"lg_order_code,omitempty"`
    

    // 错误信息
    
    SubErrorMsg   string `json:"sub_error_msg,omitempty"`
    

    // 错误code
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误code
    
    SubErrorCode   string `json:"sub_error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 资源列表
    
    Resources  struct {
        ThreePlConsignResourceDto  []ThreePlConsignResourceDto `json:"three_pl_consign_resource_dto,omitempty"`
    } `json:"resources,omitempty"`
    

}
*/

// TopResult 
type TopResult struct {

    // 发货完成后的物流单号
    LgOrderCode   string `json:"lg_order_code,omitempty"`

    // 错误信息
    SubErrorMsg   string `json:"sub_error_msg,omitempty"`

    // 错误code
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误code
    SubErrorCode   string `json:"sub_error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 资源列表
    Resources   []ThreePlConsignResourceDto `json:"resources,omitempty"`

}
