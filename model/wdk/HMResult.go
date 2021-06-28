package wdk

// HMResult 
/* model for simplify = false
type HMResult struct {

    // 配置信息列表
    
    Model  *struct {
        ConveyorBasicConfigDTO  *ConveyorBasicConfigDTO `json:"conveyor_basic_config_dto,omitempty"`
    } `json:"model,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// HMResult 
type HMResult struct {

    // 配置信息列表
    Model   *ConveyorBasicConfigDTO `json:"model,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
